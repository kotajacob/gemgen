gemgen(1) [""["gemgen-VERSION"]]

# NAME

gemgen - Markdown to Gemtext

# SYNOPSIS

_gemgen_ [OPTION]... [FILE]...

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

*-t templates*
	Input file matching patterns and templates. See *TEMPLATES* for details.

*-e, --emphasis mode*
	Representation of bold, italics, inline code, and strikethrough.

	*none* : Do not print emphasis marks.

	*markdown* : Print markdown style emphasis marks.

*-A, --heading-newline*
	Disable blank lines after headings.

*-a, --heading-links mode*
	Specify a heading link mode.

	*off* : Ignore links in headings; writing the label of the link in it's
	place.

	*auto* : If the heading contains on links, use the first link instead of
	printing a heading. Otherwise print a heading, ignoreing links.

	*below* : Print all links below headings.

*-p, --paragraph-links mode*
	Specify a paragraph link mode.

	*off* : Ignore links in paragraphs; writing the label of the link in it's
	place.

	*below* : Print all links below paragraph.

# TEMPLATES

Gemgen provides a mechanism for creating templates to add footers, headers, and
include metadata in your gemtext output files. You can specify any number of
template files with regular expression patterns to match them against your input
files. Only the "most specific" template will be applied to each input file.

For example, you can have a "catch all" template and another more specific
template that only matches a file named "index.md".

; NOTE: scdoc interprets \ inside ``` so we must double them!
```
gemgen -t '.+\\.md,main.tmpl,index\\.md,index.tmpl' file.md index.md
```

The regular expression syntax follows the RE2 standard. It should be very
familiar if you're used PCRE or python regular expressions. An exact definition
of the syntax accepted is described at https://golang.org/s/re2syntax, except
for \\C.

Gemgen uses Go's text/template library for the basis of templating. A template
is is UTF-8-encoded text containing "Actions"--data evaluations or control
structures--are delimited by "{{" and "}}". Several actions exist to add basic
metadata to your output files. Below is an example template that adds a header
link to the top of your output files (using the {{.Content}} action to print the
converted gemtext.

```
=> /index.gmi My Fancy Gemini Capsule

{{.Content}}
```

Here's a slightly more complicated template that prints a title (based on the
filename) and the original filename with a last edit timestamp.

```
{{.Title}}
{{.Filename}} last edited: {{.LastEdit}}

{{.Content}}
```

## ACTIONS

*{{.Content}}*
	Full converted gemtext data of the source file.

*{{.Filename}}*
	Original source filename (including original file extension).

*{{.Title}}*
	A title based on the source filename. The extension is removed, underscores
	are replaced with spaces, and non-adpositional words are capitalized. A
	filename like "disable_a_button.md" becomes "Disable a Button".

*{{.LastEdit}}*
	Most recent modification time of the source file, per filesystem statistics.
	This action has quite a few methods which are useful for stylizing your date
	format. For example, {{.LastEdit.Year}} prints only the year. A list of
	methods is found here: https://pkg.go.dev/time#Time

# AUTHORS

Maintained by Dakota Walsh <kota at nilsu.org>.
Up-to-date sources can be found at https://git.sr.ht/~kota/gemgen/
