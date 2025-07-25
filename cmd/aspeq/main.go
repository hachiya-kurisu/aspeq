package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"blekksprut.net/aspeq"
)

func main() {
	v := flag.Bool("v", false, "version")
	l := flag.Bool("l", false, "list available aspect ratios and exit")
	s := flag.Bool("s", false, "short - don't show filenames")
	x := flag.Bool("x", false, "aspect ratio as w:h")
	o := flag.Bool("o", false, "show image orientation")
	c := flag.Bool("c", false, "generate css and exit")

	flag.Parse()

	if *v {
		fmt.Printf("%s %s\n", os.Args[0], aspeq.Version)
		os.Exit(0)
	}

	if *l {
		for _, ar := range aspeq.Ratios {
			fmt.Printf("%s: %dx%d\n", ar.Name, ar.X, ar.Y)
		}
		os.Exit(0)
	}

	if *c {
		for _, ar := range aspeq.Ratios {
			fmt.Printf(".%s { aspect-ratio: %d/%d; }\n", ar.Name, ar.X, ar.Y)
		}
		return
	}

	if flag.NArg() < 1 {
		flag.Usage()
		return
	}

	var ar *aspeq.AspectRatio
	var err error
	for _, arg := range flag.Args() {
		if arg == "-" {
			ar, err = aspeq.FromReader(os.Stdin)
		} else {
			ar, err = aspeq.FromPath(arg)
		}
		if err != nil {
			log.Fatal(err)
		}

		ratio := ar.Name
		if *x {
			ratio = ar.Xy()
		}
		if *o {
			switch ar.Orientation {
			case aspeq.Portrait:
				ratio = "portrait"
			case aspeq.Landscape:
				ratio = "landscape"
			default:
				ratio = "square"
			}
		}

		if *s {
			fmt.Printf("%s\n", ratio)
		} else {
			fmt.Printf("%s: %s\n", arg, ratio)
		}
	}
}
