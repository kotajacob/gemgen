package options

import (
	"bytes"
	"errors"
	"fmt"
	"os"

	gem "git.sr.ht/~kota/goldmark-gemtext"
	flag "github.com/spf13/pflag"
)

var ErrVersion = errors.New("version requested")

// Opts represents options selected from command line flags.
type Opts struct {
	// GemOptions contains options for goldmark-gemtext.
	GemOptions []gem.Option
	// Names of files to convert.
	Names []string
	// Output specifies where to write gemtext files.
	// If output is blank gemtext files will be written in the source folder.
	Output string
	// TemplateArgs is an even slice of strings.
	// Every even string (starting at 0) should be a regular expression for
	// matching input filenames. Every odd string should be a filepath to a
	// loadable template file.
	// TemplateArgs should be parsed into a matchedTemplates.
	TemplateArgs []string
}

// ParseArgs parses the command-line arguments provided to the program.
// Typically os.Args[0] is provided as 'progname' and os.Args[1:] as 'args'.
// Returns Opts in case parsing succeeded, or an error. In any case, the usage
// text of the flag.Parse is returned.
// A special case is usage requests with -h or -help: then the error
// flag.ErrHelp is returned and output will contain the usage message.
// Another special case is version in which the error will be ErrVersion.
func ParseArgs(progname string, args []string) (*Opts, string, error) {
	// Create flagset.
	flag := flag.NewFlagSet(progname, flag.ContinueOnError)
	var buf bytes.Buffer
	flag.SetOutput(&buf)
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTION]... [FILE]...\n", os.Args[0])
		flag.PrintDefaults()
	}

	// Define flags.
	versionFlag := flag.BoolP("version", "v", false, "print version and exit")
	outputFlag := flag.StringP("output", "o", "", "directory to write gemtext files")
	templateFlag := flag.StringSliceP("template", "t", nil, "specify templates with a regular expression matching input filenames\n\tuse the form \"pattern,/path/to/template\"")
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
	horizontalRuleFlag := flag.StringP("horizontal-rule", "r", "~~~", "representation of horizontal rules")

	err := flag.Parse(args)
	if err != nil {
		return nil, buf.String(), err
	}
	var opts Opts
	opts.Names = flag.Args()
	if *versionFlag {
		return nil, "", ErrVersion
	}

	// Read boolean and string flags.
	opts.Output = *outputFlag
	opts.TemplateArgs = *templateFlag
	opts.GemOptions = append(
		opts.GemOptions,
		gem.WithHorizontalRule(*horizontalRuleFlag),
	)

	// Create gemtext options from flags.
	switch *emphasisFlag {
	case "none":
	case "markdown":
		opts.GemOptions = append(
			opts.GemOptions,
			gem.WithEmphasis(gem.EmphasisMarkdown),
			gem.WithCodeSpan(gem.CodeSpanMarkdown),
			gem.WithStrikethrough(gem.StrikethroughMarkdown),
		)
	case "unicode":
		opts.GemOptions = append(
			opts.GemOptions,
			gem.WithEmphasis(gem.EmphasisUnicode),
			gem.WithStrikethrough(gem.StrikethroughUnicode),
		)
	}

	if *headingNewlineFlag {
		opts.GemOptions = append(opts.GemOptions, gem.WithHeadingSpace(gem.HeadingSpaceSingle))
	}

	switch *headingLinkFlag {
	case "auto":
	case "off":
		opts.GemOptions = append(opts.GemOptions, gem.WithHeadingLink(gem.HeadingLinkOff))
	case "below":
		opts.GemOptions = append(opts.GemOptions, gem.WithHeadingLink(gem.HeadingLinkBelow))
	default:
		return nil, "", fmt.Errorf("heading link flag type %s is invalid", *headingLinkFlag)
	}

	switch *paragraphLinkFlag {
	case "below":
	case "off":
		opts.GemOptions = append(opts.GemOptions, gem.WithParagraphLink(gem.ParagraphLinkOff))
	default:
		return nil, "", fmt.Errorf("paragraph link flag type %s is invalid", *paragraphLinkFlag)
	}
	return &opts, buf.String(), nil
}
