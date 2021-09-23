package main

import (
	"bytes"
	"log"
	"os"

	gem "git.sr.ht/~kota/goldmark-gemtext"
	flag "github.com/spf13/pflag"
)

// options parses the command-line arguments provided to the program.
// Typically os.Args[0] is provided as 'progname' and os.Args[1:] as 'args'.
// Returns the gemtext options in case parsing succeeded, or an error. In any
// case, the output of the flag.Parse is returned in output.
// A special case is usage requests with -h or -help: then the error
// flag.ErrHelp is returned and output will contain the usage message.
func options(progname string, args []string) (options []gem.Option, output string, err error) {
	// setup flagset
	flag := flag.NewFlagSet(progname, flag.ContinueOnError)
	var buf bytes.Buffer
	flag.SetOutput(&buf)

	// define flags
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

	err = flag.Parse(args)
	if err != nil {
		return nil, buf.String(), err
	}
	if *versionFlag {
		log.Println("gemgen v" + Version)
		os.Exit(0)
	}

	// create gemtext options from flags
	switch *emphasisFlag {
	case "none":
	case "markdown":
		options = append(
			options,
			gem.WithEmphasis(gem.EmphasisMarkdown),
			gem.WithCodeSpan(gem.CodeSpanMarkdown),
			gem.WithStrikethrough(gem.StrikethroughMarkdown),
		)
	case "unicode":
		options = append(
			options,
			gem.WithEmphasis(gem.EmphasisUnicode),
			gem.WithStrikethrough(gem.StrikethroughUnicode),
		)
	}

	if *headingNewlineFlag {
		options = append(options, gem.WithHeadingSpace(gem.HeadingSpaceSingle))
	}

	switch *headingLinkFlag {
	case "auto":
	case "off":
		options = append(options, gem.WithHeadingLink(gem.HeadingLinkOff))
	case "below":
		options = append(options, gem.WithHeadingLink(gem.HeadingLinkBelow))
	default:
		log.Println("unknown link mode")
	}

	switch *paragraphLinkFlag {
	case "off":
		options = append(options, gem.WithParagraphLink(gem.ParagraphLinkOff))
	case "below":
		options = append(options, gem.WithParagraphLink(gem.ParagraphLinkBelow))
	default:
		log.Println("unknown link mode")
	}
	return options, buf.String(), nil
}
