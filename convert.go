package main

import (
	"bytes"
	"fmt"
	"io"

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
