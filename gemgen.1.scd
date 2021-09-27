gemgen(1) [""["gemgen-VERSION"]]

# NAME

gemgen - Markdown to Gemtext

# SYNOPSIS

*gemgen [OPTION]... [FILE]...*

# DESCRIPTION

*gemgen* is a tool for converting Commonmark Markdown to Gemtext. Gemgen uses
the goldmark markdown parser and my gemtext rendering module.

The goal is to create proper _hand-made_ gemtext. Links and "autolinks" get
placed below each paragraph, but a list of links is not printed twice.
Paragraphs get merged onto a single line, but hardlinks (double spaces or \ at
the end of a line) will insert manual line breaks. Lists and headings get
simplified to the gemtext format, emphasis markings get removed (or kept with
the *-e* option), horizontal rules get turned into 80 character lines, and
indented code gets converted to the gemtext "fenced" format.

# OPTIONS

	*-v*
		Print version and exit.
	*-o, --output directory*
		Directory to write gemtext files.
	*-e, --emphasis mode*
		Representation of bold, italics, inline code, and strikethrough.
			*none*      Do not print emphasis marks.
			*markdown*  Print markdown style emphasis marks.
	*-A, --heading-newline*
		Disable blank lines after headings.
	*-a, --heading-links mode*
		Specify a heading link mode.
			*off*   Ignore links in headings; writing the label of the link in it's
				place.
			*auto*  If the heading contains on links, use the first link instead of
				printing a heading. Otherwise print a heading, ignoreing links.
			*below* Print all links below headings.
	*-p, --paragraph-links mode*
		Specify a paragraph link mode.
			*off*   Ignore links in paragraphs; writing the label of the link in it's
				place.
			*below* Print all links below paragraph.

# AUTHORS

Maintained by Dakota Walsh <kota at nilsu.org>.
Up-to-date sources can be found at https://git.sr.ht/~kota/gemgen/
