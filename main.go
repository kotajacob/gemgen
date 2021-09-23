package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"sync"

	gem "git.sr.ht/~kota/goldmark-gemtext"
	flag "github.com/spf13/pflag"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

// Version is a semantic version for gemgen. It is set externally at build time
// from the Makefile.
var Version string

func main() {
	log.SetPrefix("")
	log.SetFlags(0)

	// get options
	opts, output, err := parseArgs(os.Args[0], os.Args[1:])
	if err == flag.ErrHelp {
		log.Println(output)
		os.Exit(0)
	} else if err != nil {
		log.Println("got error:", err)
		log.Println("output:\n", output)
		os.Exit(1)
	}

	// use stdin if no files were given
	if opts.Names == nil {
		src, err := io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalf("failed reading STDIN: %v\n", err)
		}
		err = convertFile(&src, opts.GemOptions)
		if err != nil {
			log.Fatalf("failed converting STDIN: %v\n", err)
		}
		os.Exit(0)
	}

	// read and convert the list of files concurrently
	var wg sync.WaitGroup
	for _, name := range opts.Names {
		wg.Add(1)
		go func(name string) {
			// decrement the counter when the goroutine completes
			defer wg.Done()
			src, err := os.ReadFile(name)
			if err != nil {
				log.Fatalf("failed reading file %s: %v\n", name, err)
			}
			err = convertFile(&src, opts.GemOptions)
			if err != nil {
				log.Fatalf("failed converting file %s: %v\n", name, err)
			}
		}(name)
	}
	wg.Wait()
}

// convertFile reads the file and converts it to gemtext using opts.
func convertFile(file *[]byte, opts []gem.Option) error {
	// create markdown parser
	var buf bytes.Buffer
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.Linkify,
			extension.Strikethrough,
		),
	)

	// render
	md.SetRenderer(gem.New(opts...))
	if err := md.Convert([]byte(*file), &buf); err != nil {
		return err
	}
	io.Copy(os.Stdout, &buf)
	return nil
}
