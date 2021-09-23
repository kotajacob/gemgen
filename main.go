package main

import (
	"bytes"
	"io"
	"log"
	"os"

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
	opts, output, err := options(os.Args[0], os.Args[1:])
	if err == flag.ErrHelp {
		log.Println(output)
		os.Exit(0)
	} else if err != nil {
		log.Println("got error:", err)
		log.Println("output:\n", output)
		os.Exit(1)
	}

	// load markdown
	src, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	// render
	if err := render(src, opts); err != nil {
		log.Fatal(err)
	}
}

func render(src []byte, opts []gem.Option) error {
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
	if err := md.Convert([]byte(src), &buf); err != nil {
		return err
	}
	io.Copy(os.Stdout, &buf)
	return nil
}
