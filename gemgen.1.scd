gemgen(1) [""["gemgen-VERSION"]]

# NAME

gemgen - Markdown to Gemtext

# SYNOPSIS

*gemgen [-v | -e | -E | -h] [-H linkmode]*

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
	*-e*
		Print markdown emphasis symbols for bold, italics, inline code, and
		strikethrough.
	*-E*
		Print unicode symbols for 𝗯𝗼𝗹𝗱, 𝘪𝘵𝘢𝘭𝘪𝘤, and s̶t̶r̶i̶k̶e̶t̶h̶r̶o̶u̶g̶h̶.
	*-h*
		Disable blank lines after headings.
	*-H*
		Specify a heading link mode.
			*off*   Ignore links in headings; writing the label of the link in it's
				place.
			*auto*  If the heading contains on links, use the first link instead of
				printing a heading. Otherwise print a heading, ignoreing links.
			*below* Print all links below headings.
	*-P*
		Specify a paragraph link mode.
			*off*   Ignore links in paragraphs; writing the label of the link in it's
				place.
			*below* Print all links below paragraph.

# AUTHORS

Maintained by Dakota Walsh <kota at nilsu.org>.
Up-to-date sources can be found at https://git.sr.ht/~kota/gemgen/
