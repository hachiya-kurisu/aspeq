package main

import (
	"blekksprut.net/aspeq"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	version := flag.Bool("v", false, "version")
	short := flag.Bool("s", false, "short - don't show filenames")
	x := flag.Bool("x", false, "aspect ratio as w:h")
	css := flag.Bool("c", false, "generate css and exit")
	flag.Parse()

	if *version {
		fmt.Printf("%s %s\n", os.Args[0], aspeq.Version)
		os.Exit(0)
	}

	if *css {
		for _, ar := range aspeq.Ratios {
			fmt.Printf(".%s { aspect-ratio: %d/%d; }\n", ar.Name, ar.X, ar.Y)
		}
		return
	}

	if flag.NArg() < 1 {
		fmt.Fprintf(os.Stderr, "%s [-hxv] images...\n", os.Args[0])
		return
	}

	for _, arg := range flag.Args() {
		ar, err := aspeq.FromImage(arg)
		if err != nil {
			log.Fatal(err)
			return
		}

		var ratio string
		if *x {
			ratio = ar.Xy()
		} else {
			ratio = ar.Name
		}

		if *short {
			fmt.Printf("%s\n", ratio)
		} else {
			fmt.Printf("%s: %s\n", arg, ratio)
		}
	}
}
