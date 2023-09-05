package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	appName = "tfjson2md"
	version = "dev"
)

const (
	ExitOK      = 0
	ExitErr     = 1
	ExitInvalid = 2
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %[1]s [file]\nIf file is a single dash (‘-’) or absent, %[1]s reads from the standard input.\n", appName)
	flag.PrintDefaults()
	os.Exit(ExitInvalid)
}

func main() {
	log.SetFlags(0)
	log.SetPrefix(fmt.Sprintf("%s: ", appName))

	var versionFlag bool
	flag.BoolVar(&versionFlag, "v", false, "show version")
	flag.Usage = usage
	flag.Parse()

	if versionFlag {
		fmt.Printf("%s v%s\n", appName, version)

		return
	}

	var input io.Reader

	file := flag.Arg(0)

	// operate like `cat` and read from stdin if no file or `-`
	if file == "" || file == "-" {
		input = os.Stdin
	} else {
		f, err := os.Open(file)
		if err != nil {
			log.Fatalf("could not open file for reading: %q", err)
		}
		defer f.Close()

		input = f
	}

	if err := run(input); err != nil {
		log.Fatal(err)
	}
}

func run(input io.Reader) error {
	return nil
}
