# Change Log

## [0.5.1] - 2022-01-31
### Fixed
- Reading markdown from STDIN and writing gemtext to STDOUT.

## [0.5.0] - 2021-12-29
### Added
- Horizontal rule flag allowing any string including newlines.

### Changed
- Default horizontal was changed to ~~~ for better accessibility.

## [0.4.2] - 2021-11-21
### Fixed
- Fix a bug in which the newly added wiki links could be skipped.

## [0.4.1] - 2021-11-20
### Added
- Parsing and rendering of links from the wiki style link markdown extension.

## [0.4.0] - 2021-10-31
### Added
- Templating which allows automatically adding headers, footers, and metadata to
  your output gemtext files.

## [0.3.0] - 2021-09-28
### Added
- Concurrent file reading/writing.
- Options for link handling in paragraphs and headers.
- Option for blank lines after headings.
- Long versions of every cli flag.

### Changed
- The -E option was replaced with -e type, where type can be off, markdown, or
  unicode.

## [0.2.0] - 2021-08-11
### Added
- The -E option which prints bold, italic, and strikethrough with unicode hacks.
  Note: this has since been changed and may be removed. It's quite bad for
  accessibility.

## [0.1.0] - 2021-08-06
### Added
- First release! Basic markdown to gemtext functionality is implemented.
