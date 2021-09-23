package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	gem "git.sr.ht/~kota/goldmark-gemtext"
	flag "github.com/spf13/pflag"
)

// options parses the command-line arguments provided to the program.
// Typically os.Args[0] is provided as 'progname' and os.Args[1:] as 'args'.
// Returns the gemtext options and requested files in case parsing succeeded,
// or an error. In any case, the output of the flag.Parse is returned in
// output.
// A special case is usage requests with -h or -help: then the error
// flag.ErrHelp is returned and output will contain the usage message.
func options(progname string, args []string) (opts []gem.Option, files []string, output string, err error) {
	// setup flagset
	flag := flag.NewFlagSet(progname, flag.ContinueOnError)
	var buf bytes.Buffer
	flag.SetOutput(&buf)
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTION]... [FILE]...\n", os.Args[0])
		flag.PrintDefaults()
	}

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
		return nil, nil, buf.String(), err
	}
	files = flag.Args()
	if *versionFlag {
		log.Println("gemgen v" + Version)
		os.Exit(0)
	}

	// create gemtext options from flags
	switch *emphasisFlag {
	case "none":
	case "markdown":
		opts = append(
			opts,
			gem.WithEmphasis(gem.EmphasisMarkdown),
			gem.WithCodeSpan(gem.CodeSpanMarkdown),
			gem.WithStrikethrough(gem.StrikethroughMarkdown),
		)
	case "unicode":
		opts = append(
			opts,
			gem.WithEmphasis(gem.EmphasisUnicode),
			gem.WithStrikethrough(gem.StrikethroughUnicode),
		)
	}

	if *headingNewlineFlag {
		opts = append(opts, gem.WithHeadingSpace(gem.HeadingSpaceSingle))
	}

	switch *headingLinkFlag {
	case "auto":
	case "off":
		opts = append(opts, gem.WithHeadingLink(gem.HeadingLinkOff))
	case "below":
		opts = append(opts, gem.WithHeadingLink(gem.HeadingLinkBelow))
	default:
		log.Println("unknown link mode")
	}

	switch *paragraphLinkFlag {
	case "below":
	case "off":
		opts = append(opts, gem.WithParagraphLink(gem.ParagraphLinkOff))
	default:
		log.Println("unknown link mode")
	}
	return opts, files, buf.String(), nil
}
