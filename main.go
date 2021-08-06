package main

import (
	"io"
	"log"
	"os"

	gemtext "git.sr.ht/~kota/goldmark-gemtext"
	"git.sr.ht/~sircmpwn/getopt"
)

var (
	in      = os.Stdin
	out     = os.Stdout
	Version string
)

func usage() {
	log.Fatal(`gemgen [-v] [-e] [-i input.md] [-o output.gmi]
 -v : Print version and exit.
 -e : Keep emphasis symbols for bold, italics, inline code, and strikethrough.
 -i : Read from a file instead of standard input.
 -o : Write to an output file instead of standard output.`)
}

func main() {
	log.SetPrefix("")
	log.SetFlags(0)
	opts, _, err := getopt.Getopts(os.Args, "vei:o:")
	if err != nil {
		log.Print(err)
		usage()
	}
	for _, opt := range opts {
		switch opt.Option {
		case 'v':
			log.Println("gemgen v" + Version)
			os.Exit(0)
		case 'e':
			gemtext.Emphasis = true
			gemtext.CodeSpan = true
			gemtext.Strikethrough = true
		case 'i':
			if opt.Value == "-" {
				continue
			}
			in, err = os.Open(opt.Value)
			if err != nil {
				log.Print(err)
				usage()
			}
		case 'o':
			out, err = os.Create(opt.Value)
			if err != nil {
				log.Print(err)
				usage()
			}
		}
	}

	// load markdown
	source, err := io.ReadAll(in)
	if err != nil {
		log.Fatal(err)
	}

	// parse markdown and write reformatted source
	err = gemtext.Format(source, out)
	if err != nil {
		log.Fatal(err)
	}
}
