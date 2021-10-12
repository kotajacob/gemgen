package main

import (
	"fmt"
	"log"
	"os"

	"git.sr.ht/~kota/gemgen/matchtemplate"
	"git.sr.ht/~kota/gemgen/options"
	"github.com/spf13/afero"
	flag "github.com/spf13/pflag"
)

// Version is a semantic version for gemgen. It is set externally at build time
// from the Makefile.
var Version string

func main() {
	log.SetPrefix("")
	log.SetFlags(0)

	// Parse arguments.
	opts, usage, err := options.ParseArgs(os.Args[0], os.Args[1:])
	if err == flag.ErrHelp {
		log.Println(usage)
		os.Exit(0)
	} else if err == options.ErrVersion {
		log.Println("gemgen v" + Version)
		os.Exit(0)
	} else if err != nil {
		log.Println("error:", err)
		if usage != "" {
			log.Println("output:", usage)
		}
		os.Exit(1)
	}

	// Use stdin if no filenames were given.
	if opts.Names == nil {
		err = Convert(os.Stdin, os.Stdout, opts.GemOptions)
		if err != nil {
			log.Fatalf("failed converting STDIN: %v\n", err)
		}
		os.Exit(0)
	}

	// Load templates.
	mt := new(matchtemplate.MatchedTemplates)
	if err := mt.Parse(opts); err != nil {
		log.Fatalf("failed parsing templates: %v\n", err)
	}
	fmt.Println(mt)

	// Convert named files.
	fs := afero.NewOsFs()
	ConvertFiles(fs, opts, mt)
}
