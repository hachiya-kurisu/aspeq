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
	x := flag.Bool("x", false, "aspect ratio as w:h")
	flag.Parse()

	if *version {
		fmt.Printf("%s %s\n", os.Args[0], aspeq.Version)
		os.Exit(0)
	}

	if flag.NArg() < 1 {
		fmt.Fprintf(os.Stderr, "%s [-hxv] images...\n", os.Args[0])
		return
	}

	for _, arg := range flag.Args() {
		ratio, err := aspeq.FromImage(arg)
		if err != nil {
			log.Fatal(err)
			return
		}
		if *x {
			fmt.Printf("%s: %s\n", arg, ratio.Xy())
		} else {
			fmt.Printf("%s: %s\n", arg, ratio.Name)
		}
	}
}
