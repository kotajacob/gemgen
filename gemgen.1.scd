gemgen(1) [""["gemgen-VERSION"]]

# NAME

gemgen - Markdown to Gemtext

# SYNOPSIS

_gemgen_ [OPTION]... [FILE]...

# DESCRIPTION

*gemgen* is a tool for converting Commonmark Markdown to Gemtext. Gemgen uses
the goldmark markdown parser and my gemtext rendering module. You may pass
gemgen a list of files to convert and rename concurrently or pass markdown into
STDIN and get gemtext from STDOUT.

The goal is to create proper _hand-made_ gemtext. Links and "autolinks" get
placed below each paragraph, but a list of links is not printed twice.
Paragraphs get merged onto a single line, but hardlinks (double spaces or \ at
the end of a line) will insert manual line breaks.

Lists and headings get simplified to the gemtext format, emphasis markings get
removed (or kept with the *-e* option), horizontal rules get turned into ~~~ or
any other string you'd like with *-r* option, and indented code gets converted
to the gemtext "fenced" format. Commonmark markdown is fully supported and a few
common extensions have been added: autolinks, strikethrough, and wiki style
links.

Gemgen includes a templating system allowing you to add headers, footers, last
edit timestamps and more automatically to your output files. See *TEMPLATES*
below for details.

# OPTIONS

*-v*
	Print version and exit.

*-o, --output directory*
	Directory to write gemtext files.

*-r, --horizontal-rule string*
	Representation of horizontal rules (default "~~~").

*-t templates*
	Input file matching patterns and templates. See *TEMPLATES* for details.

*-l linkregex*
	Specify a regular expression subsitution to apply to different kinds of links.
	See *LINKREGEX* for details.

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

# LINKREGEX

Gemgen provides a mechanism for modifying the destination of links while they're
being parsed. You can create subsitutions for different kinds of links
(markdown, auto, wiki, and image), then match part of the destination with a
regular expression, and then replace it with a subsitution string.

The syntax is similar to *TEMPLATES*. The form is "type,regex,replacement". For
example if I want to match wiki style links and replace whitespace with
underscores:

```
gemgen -l "wiki,\s,_"
```

You can include multiple link subsitutions in a row and they will be applied in
order for each matching link. In my personal vim.wiki I use the following:

```
gemgen -l "wiki,\s,_,wiki,^,/memex/,image,^,/memex/,wiki,$,.gmi"
```

*NOTE:* You can also use ^ or $ to match the start / end of a link if you need
to add a string there without replacing anything.

## LINK TYPES

There are currently 4 different kinds of links recognized by gemgen.

*markdown*
	These are your standard markdown style links. They use the form
	*[text](destination)*.

*wiki*
	This type of link is a markdown extension used by some popular wiki software
	such as vim.wiki. They use the form *[[destination]]* or the form
	*[[destination|text]]*.

*auto*
	Automatic link detection is technically another markdown extension, but are
	used in most markdown processors such as github. They work by simply pasting a
	link in your document. It cannot contain spaces for specifics visit:
	https://github.github.com/gfm/#autolinks-extension-

*image*
	These are standard markdown image links. In HTML they would be displayed
	inline, but in gemtext they're represented as any other link. Syntax for
	images is like the syntax for links but prepended with an exclamation mark
	*![text](destination)*.

# AUTHORS

Maintained by Dakota Walsh <kota at nilsu.org>.
Up-to-date sources can be found at https://git.sr.ht/~kota/gemgen/
