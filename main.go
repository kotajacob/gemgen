package main

import (
	"bytes"
	"io"
	"log"
	"os"

	gem "git.sr.ht/~kota/goldmark-gemtext"
	"git.sr.ht/~sircmpwn/getopt"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

var (
	inputBytes = os.Stdin
	out        = os.Stdout
	Version    string
)

func usage() {
	log.Fatal(`gemgen [-e | -E | -h] [-H linkmode] input.md
-v : Print version and exit.
-e : Print markdown emphasis symbols for bold, italics, inline code, and strikethrough.
-E : Print unicode symbols for 𝗯𝗼𝗹𝗱, 𝘪𝘵𝘢𝘭𝘪𝘤, and s̶t̶r̶i̶k̶e̶t̶h̶r̶o̶u̶g̶h̶.
-h : Disable blank lines after headings.
-H : Specify a heading link mode.
	off : Ignore links in headings; writing the label of the link in it's place
	auto: If the heading contains on links, use the first link instead of printing a heading. Otherwise print a heading, ignoreing links.
	below: Print all links below headings.`)
}

func main() {
	log.SetPrefix("")
	log.SetFlags(0)
	opts, _, err := getopt.Getopts(os.Args, "veEhH:")
	if err != nil {
		log.Print(err)
		usage()
	}

	// create markdown parser
	var buf bytes.Buffer
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.Linkify,
			extension.Strikethrough,
		),
	)

	// get opts
	var gemOptions []gem.Option
	for _, opt := range opts {
		switch opt.Option {
		case 'v':
			log.Println("gemgen v" + Version)
			os.Exit(0)
		case 'e':
			gemOptions = append(
				gemOptions,
				gem.WithEmphasis(gem.EmphasisMarkdown),
				gem.WithCodeSpan(gem.CodeSpanMarkdown),
				gem.WithStrikethrough(gem.StrikethroughMarkdown),
			)
		case 'E':
			gemOptions = append(
				gemOptions,
				gem.WithEmphasis(gem.EmphasisUnicode),
				gem.WithStrikethrough(gem.StrikethroughUnicode),
			)
		case 'h':
			gemOptions = append(gemOptions, gem.WithHeadingSpace(gem.HeadingSpaceSingle))
		case 'H':
			switch opt.Value {
			case "auto":
			case "off":
				gemOptions = append(gemOptions, gem.WithHeadingLink(gem.HeadingLinkOff))
			case "below":
				gemOptions = append(gemOptions, gem.WithHeadingLink(gem.HeadingLinkBelow))
			default:
				log.Println("unknown link mode")
				usage()
			}
		}
	}

	// load markdown
	src, err := io.ReadAll(inputBytes)
	if err != nil {
		log.Fatal(err)
	}

	// attach gemtext renderer
	md.SetRenderer(gem.New(gemOptions...))

	// render
	if err := md.Convert([]byte(src), &buf); err != nil {
		log.Fatal(err)
	}
	io.Copy(os.Stdout, &buf)
}
