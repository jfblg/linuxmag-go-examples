package main

import (
	"flag"
	"fmt"
	"os"
	"path"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s 'search/replace' file ...\n",
		path.Base(os.Args[0]))
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	dryrun := flag.Bool("d", false, "dryrun only")
	verbose := flag.Bool("v", false, "verbose mode")
	flag.Usage = usage
	flag.Parse()

	if *dryrun {
		fmt.Printf("Dry run mode\n")
	}

	if len(flag.Args()) < 2 {
		usage()
	}

	cmd := flag.Args()[0]
	files := flag.Args()[1:]
	modifier, err := mkmodifier(cmd)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid command %s\n", cmd)
		usage()
	}

	for _, file := range files {
		modfile := modifier(file)
		if file == modfile {
			continue
		}

		if *verbose || *dryrun {
			fmt.Printf("%s -> %s\n", file, modfile)
		}
		if *dryrun {
			continue
		}
		err := os.Rename(file, modfile)
		if err != nil {
			fmt.Printf("Renaming %s -> &%s failed: %v\n", file, modfile, err)
			break
		}
	}
}
