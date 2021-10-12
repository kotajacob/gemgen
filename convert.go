package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"path/filepath"
	"sync"
	"text/template"

	"git.sr.ht/~kota/gemgen/matchtemplate"
	"git.sr.ht/~kota/gemgen/options"
	gem "git.sr.ht/~kota/goldmark-gemtext"
	"github.com/spf13/afero"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

var DefaultTemplate = template.Must(template.New("default").Parse("{{.Content}}"))

type Gemtext struct {
	Content string
}

// ConvertFiles reads Opts and converts the list of named files concurrently.
// An afero filesystem is used for abstraction. You can create an OS based
// filesystem with afero.NewOsFs() or a memory backed system with
// afero.NewMemMapFs().
// Files are written with the .gmi extension in the source directory.
// Encountering an error stops the program with an appropriate message.
func ConvertFiles(fs afero.Fs, opts *options.Opts, mt *matchtemplate.MatchedTemplates) error {
	// Read and convert the list of files concurrently.
	var wg sync.WaitGroup
	for _, name := range opts.Names {
		wg.Add(1)
		go func(name string) {
			// Create Gemtext to store converted data and metadata.
			var g Gemtext

			// Decrement the counter when the goroutine completes.
			defer wg.Done()

			// Read input file.
			src, err := fs.Open(name)
			if err != nil {
				log.Fatalf("failed reading input file %s: %v\n", name, err)
			}

			// Convert to gemtext and store output.
			var buf bytes.Buffer
			err = Convert(src, &buf, opts.GemOptions)
			if err != nil {
				log.Fatalf("failed converting file %s: %v\n", name, err)
			}
			g.Content = buf.String()

			// Apply template
			var tmplOut bytes.Buffer
			tmpl := mt.Lookup(name)
			if tmpl == nil {
				// No template found. Use default template.
				tmpl = DefaultTemplate
			}
			tmpl.Execute(&tmplOut, g)

			// Write output to file.
			if err := store(fs, name, opts.Output, tmplOut.Bytes()); err != nil {
				log.Fatalf("failed writing file %s: %v\n", name, err)
			}
		}(name)
	}
	wg.Wait()
	return nil
}

// Convert reads markdown data and writes it as gemtext using opts.
func Convert(r io.Reader, w io.Writer, opts []gem.Option) error {
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

// store data in a file.
// inpath is the full path to the input file (with a markdown extension).
// output is an optional output path (just the directory without the file).
// An afero filesystem is used for abstraction. You can create an OS based
// filesystem with afero.NewOsFs() or a memory backed system with
// afero.NewMemMapFs().
func store(fs afero.Fs, input string, output string, data []byte) error {
	base := filepath.Base(input)
	outName := base[0:len(base)-len(filepath.Ext(base))] + ".gmi"
	var path string
	if output != "" {
		// Use configured output directory.
		path = filepath.Join(output, outName)
	} else {
		// Use input directory as output directory.
		path = filepath.Join(filepath.Dir(input), outName)
	}
	err := afero.WriteFile(fs, path, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
