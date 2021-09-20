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

var Version string

func main() {
	log.SetPrefix("")
	log.SetFlags(0)

	// get options
	opts := options()

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

func options() []gem.Option {
	versionFlag := flag.BoolP("version", "v", false, "print version and exit")
	emphasisFlag := flag.StringP("emphasis", "e", "none", `representation of bold, italics, inline code, and strikethrough
	none     : do not print emphasis marks
	markdown : print markdown style emphasis marks`)
	headingLinkFlag := flag.StringP("heading-links", "a", "auto", `specify how links in headings are printed
	auto  : print link-only headings as links
	below : print links in headings below the heading
	off   : ignore links in headings`)
	paragraphLinkFlag := flag.StringP("paragraph-links", "p", "below", `specify how links in paragraphs are printed
	below : print links in paragraphs below the paragraph
	off   : ignore links in paragraphs`)
	headingNewlineFlag := flag.BoolP("heading-newline", "A", false, `disable printing a newline below each heading`)
	flag.Parse()

	// use command line flags to create parser options
	if *versionFlag == true {
		log.Println("gemgen v" + Version)
		os.Exit(0)
	}
	var gemOptions []gem.Option
	switch *emphasisFlag {
	case "none":
	case "markdown":
		gemOptions = append(
			gemOptions,
			gem.WithEmphasis(gem.EmphasisMarkdown),
			gem.WithCodeSpan(gem.CodeSpanMarkdown),
			gem.WithStrikethrough(gem.StrikethroughMarkdown),
		)
	case "unicode":
		gemOptions = append(
			gemOptions,
			gem.WithEmphasis(gem.EmphasisUnicode),
			gem.WithStrikethrough(gem.StrikethroughUnicode),
		)
	}

	if *headingNewlineFlag == true {
		gemOptions = append(gemOptions, gem.WithHeadingSpace(gem.HeadingSpaceSingle))
	}

	switch *headingLinkFlag {
	case "auto":
	case "off":
		gemOptions = append(gemOptions, gem.WithHeadingLink(gem.HeadingLinkOff))
	case "below":
		gemOptions = append(gemOptions, gem.WithHeadingLink(gem.HeadingLinkBelow))
	default:
		log.Println("unknown link mode")
	}

	switch *paragraphLinkFlag {
	case "off":
		gemOptions = append(gemOptions, gem.WithParagraphLink(gem.ParagraphLinkOff))
	case "below":
		gemOptions = append(gemOptions, gem.WithParagraphLink(gem.ParagraphLinkBelow))
	default:
		log.Println("unknown link mode")
	}
	return gemOptions
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
