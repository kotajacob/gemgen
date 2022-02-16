# Contributing to gemgen

## Architecture
The project is split into a few seperate repositories.

[gemgen](https://git.sr.ht/~kota/gemgen) is the command line tool for converting
Commonmark Markdown to Gemtext.

[goldmark-gemtext](https://git.sr.ht/~kota/goldmark-gemtext/) is a goldmark
renderer which renders parsed markdown to gemtext instead of html.

[golemark-wiki](https://git.sr.ht/~kota/goldmark-wiki) is an extension to the
goldmark markdown parser. It adds support for a new link syntax used by [vim
wiki](https://github.com/vimwiki/vimwiki) and some other wiki software.

### gemgen
`main.go` - Short and simple main function that calls the argument parser, loads
templates, and then converts STDIN or named files.

`options/options.go` - Command line argument parser that builds up an Opts
struct containing a list of files to convert, options for goldmark, list of
templates, and an output directory.

`matchtemplate/matchtemplate.go` - Defines a type that associates a
text/template with a list of files which have been matched by a regular
expression.

`convert.go` - Provides a Convert and ConvertFiles functions. The latter
concurrently reads and converts a list of files using Opts and a list of
MatchTemplates. The former does the same conversion, but from an io.Reader to an
io.Writer.

### goldmark-gemtext
`render.go` - Implements the basic functions and methods for a goldmark
Renderer, including registering a method for each AST type the renderer can
handle.

`option.go` - Defines a Config type which can be given to the GemRenderer to
adjust its output. Each value in the config can be set with a helper function
such as `WithHeadingLink` rather than building the whole config individually.

`block.go` - Contains GemRenderer methods for "Block" elements, headings,
paragraphs, lists, etc.

`inline.go` - Contains GemRenderer methods for "Inline" elements such as
emphasis, links, codespan, etc.

`extra.go` - Contains GemRenderer methods for non-standard markdown elements
such as strikethrough, wiki links, and check boxes.

### goldmark-wiki
`wiki.go` - Defines a goldmark parser extension for the wiki type and even a
basic HTML renderer. A lot of this is boiler plate and not technically used in
gemgen, but could be used in other goldmark projects that need wiki link
parsing.

`ast/wiki.go` - Defines the AST node to represent the wiki link element.

## Patches, questions, and feature requests
The project uses a public mailing list for patches and communication:
[https://lists.sr.ht/~kota/gemgen](https://lists.sr.ht/~kota/gemgen) You can
create a post by sending [plain
text](https://man.sr.ht/lists.sr.ht/etiquette.md) emails to
[~kota/gemgen@lists.sr.ht](mailto:~kota/GEMGEN@todo.sr.ht).

## Coding standards
Contributions don't need to follow these exactly to be useful.

## Other ways to help
Read the man pages and see if anything seems wrong or unclear. I don't have the
perspective of someone reading it for the first time.
