package main

import (
	"log"
	"os"

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
	opts, usage, err := parseArgs(os.Args[0], os.Args[1:])
	if err == flag.ErrHelp {
		log.Println(usage)
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
		err = convert(os.Stdin, os.Stdout, opts.GemOptions)
		if err != nil {
			log.Fatalf("failed converting STDIN: %v\n", err)
		}
		os.Exit(0)
	}

	// Convert named files.
	fs := afero.NewOsFs()
	convertFiles(fs, opts)
}
