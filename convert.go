package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	gem "git.sr.ht/~kota/goldmark-gemtext"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

// convert reads markdown data and writes it as gemtext using opts.
func convert(r io.Reader, w io.Writer, opts []gem.Option) error {
	// create markdown parser
	var buf bytes.Buffer
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.Linkify,
			extension.Strikethrough,
		),
	)
	// read markdown
	src, err := io.ReadAll(r)
	if err != nil {
		return fmt.Errorf("failed to read input file: %v", err)
	}

	// render
	md.SetRenderer(gem.New(opts...))
	if err := md.Convert(src, &buf); err != nil {
		return fmt.Errorf("failed to convert markdown to gemtext: %v", err)
	}
	io.Copy(w, &buf)
	return nil
}

// convertFiles reads Opts and converts the list of named files concurrently.
// Encountering an error stops the program with an appropriate message.
func convertFiles(opts *Opts) {
	// read and convert the list of files concurrently
	var wg sync.WaitGroup
	for _, name := range opts.Names {
		wg.Add(1)
		go func(name string) {
			// decrement the counter when the goroutine completes
			defer wg.Done()
			// read input file
			src, err := os.Open(name)
			if err != nil {
				log.Fatalf("failed reading file %s: %v\n", name, err)
			}
			err = convert(src, os.Stdout, opts.GemOptions)
			if err != nil {
				log.Fatalf("failed converting file %s: %v\n", name, err)
			}
		}(name)
	}
	wg.Wait()
}
