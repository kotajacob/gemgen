package main

import (
	"log"
	"os"

	flag "github.com/spf13/pflag"
)

// Version is a semantic version for gemgen. It is set externally at build time
// from the Makefile.
var Version string

func main() {
	log.SetPrefix("")
	log.SetFlags(0)

	// get options
	opts, output, err := parseArgs(os.Args[0], os.Args[1:])
	if err == flag.ErrHelp {
		log.Println(output)
		os.Exit(0)
	} else if err != nil {
		log.Println("got error:", err)
		log.Println("output:\n", output)
		os.Exit(1)
	}

	// use stdin if no files were given
	if opts.Names == nil {
		err = convert(os.Stdin, os.Stdout, opts.GemOptions)
		if err != nil {
			log.Fatalf("failed converting STDIN: %v\n", err)
		}
		os.Exit(0)
	}

	// (otherwise) convert named files
	convertFiles(opts)
}
