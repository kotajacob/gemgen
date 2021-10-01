package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"

	gem "git.sr.ht/~kota/goldmark-gemtext"
	"github.com/spf13/afero"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

// convertFiles reads Opts and converts the list of named files concurrently.
// Files are written with the .gmi extension in the source directory.
// Encountering an error stops the program with an appropriate message.
func convertFiles(fs afero.Fs, opts *Opts) error {
	// Read and convert the list of files concurrently.
	var wg sync.WaitGroup
	for _, name := range opts.Names {
		wg.Add(1)
		go func(name string) error {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()

			// Read input file.
			src, err := fs.Open(name)
			if err != nil {
				log.Fatalf("failed reading input file %s: %v\n", name, err)
			}

			// Open output file.
			base := filepath.Base(name)
			outName := base[0:len(base)-len(filepath.Ext(base))] + ".gmi"
			var outPath string
			if opts.Output != "" {
				// Use input directory as output directory.
				outPath = filepath.Join(opts.Output, outName)
			} else {
				// Use configured output directory.
				outPath = filepath.Join(filepath.Dir(name), outName)
			}
			out, err := fs.OpenFile(outPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
			if err != nil {
				log.Fatalf("failed opening output file %s: %v\n", outPath, err)
			}

			// Convert to gemtext and write output.
			err = convert(src, out, opts.GemOptions)
			if err != nil {
				log.Fatalf("failed converting file %s: %v\n", name, err)
			}
			return nil
		}(name)
	}
	wg.Wait()
	return nil
}

// convert reads markdown data and writes it as gemtext using opts.
func convert(r io.Reader, w io.Writer, opts []gem.Option) error {
	// Create markdown parser.
	var buf bytes.Buffer
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.Linkify,
			extension.Strikethrough,
		),
	)

	// Read markdown.
	src, err := io.ReadAll(r)
	if err != nil {
		return fmt.Errorf("failed to read input file: %v", err)
	}

	// Render to gemtext.
	md.SetRenderer(gem.New(opts...))
	if err := md.Convert(src, &buf); err != nil {
		return fmt.Errorf("failed to convert markdown to gemtext: %v", err)
	}
	io.Copy(w, &buf)
	return nil
}
