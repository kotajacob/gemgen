gemgen(1) [""["gemgen-VERSION"]]

# NAME

gemgen - Markdown to Gemtext

# SYNOPSIS

*gemgen [-v] [-e] [-i input.md] [-o output.gmi]*

# DESCRIPTION

*gemgen* is a tool for converting Commonmark Markdown to Gemtext. Gemgen uses
the goldmark markdown parser and my gemtext rendering module.

The goal is to create proper hand-made gemtext. Links and autolinks are
placed below each paragraph, but a paragraph of only links is left intact.
Normally, paragraphs are merged onto a single line, but hardlinks (double
spaces or \ at the end of a line) may be used for manual line breaks. Lists and
headings are simplified to the gemtext format, emphasis markings are removed
(or kept with the -e flag), horizontal rules are turned into 80 character
lines, and indented code is converted to the gemtext fenced format.

# OPTIONS

	*-v*
		Print version and exit.
	*-e*
		Keep emphasis symbols for bold, italics, inline code, and strikethrough.
	*-i input.md*
		Read from a file instead of standard input.
	*-o output.gmi*
		Write to an output file instead of standard output.

# AUTHORS

Maintained by Dakota Walsh <kota at nilsu.org>.
Up-to-date sources can be found at https://git.sr.ht/~kota/gemgen/
