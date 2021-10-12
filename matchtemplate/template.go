package matchtemplate

import (
	"fmt"
	"regexp"
	"text/template"

	"git.sr.ht/~kota/gemgen/options"
)

// MatchedTemplates is a list of matchTemplates.
type MatchedTemplates []MatchTemplate

// MatchTemplate contains a list of input filenames and a template.
// The matches are filesnames representing which input files the template could
// be applied onto. Only the most specific template should be applied to each
// input file. Most specific means templates are sorted from least matches to
// most and the "top" template for each file is applied.
type MatchTemplate struct {
	matches []string
	t       *template.Template
}

// Parse the command line template selection into matchedTemplates.
// The args slice should contain an even number of elements in the form:
// "pattern,/path/to/template"
func (m *MatchedTemplates) Parse(opts *options.Opts) error {
	// Ensure even number of elements.
	if len(opts.TemplateArgs)%2 != 0 {
		return fmt.Errorf(`uneven number of template options: template flag requires the form "pattern,/path/to/template"`)
	}
	// Iterate over the patterns.
	for i := 0; i <= len(opts.TemplateArgs)-1; i += 2 {
		var mt MatchTemplate
		// Parse regex and match files.
		pattern := opts.TemplateArgs[i]
		tName := opts.TemplateArgs[i+1]
		re, err := regexp.Compile(pattern)
		if err != nil {
			return fmt.Errorf("failed to parse pattern as a regular expression, see https://github.com/google/re2/wiki/Syntax for documentation of the syntax: %v\n", err)
		}
		mt.matches = matches(re, opts.Names)
		if len(mt.matches) == 0 {
			continue // Skip templates with 0 matches.
		}

		// Load and parse template file.
		mt.t, err = template.New(tName).ParseFiles(tName)
		if err != nil {
			return fmt.Errorf("failed to parse template file: %v\n", err)
		}
		*m = append(*m, mt)
	}
	return nil
}

// Lookup returns the most appropriate template for an input file.
// Check each template for matches. Return the first matching template.
func (m *MatchedTemplates) Lookup(name string) *template.Template {
	for _, t := range *m {
		for _, match := range t.matches {
			if name == match {
				return t.t
			}
		}
	}
	return nil
}

// matches filters a list of names with a regular expression and returns a list
// of strings which contained at least one match.
func matches(re *regexp.Regexp, names []string) []string {
	var matches []string
	for _, name := range names {
		if re.MatchString(name) {
			matches = append(matches, name)
		}
	}
	return matches
}
