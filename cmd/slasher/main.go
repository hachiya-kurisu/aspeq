package main

import (
	"blekksprut.net/aspeq"
	"flag"
	"fmt"
	"image/jpeg"
	"log"
	"os"
)

func main() {
	v := flag.Bool("v", false, "version")
	l := flag.Bool("l", false, "list available aspect ratios and exit")
	o := flag.String("o", "", "write cropped image to this path")
	a := flag.String("a", "", "desired ratio. defaults to closest defined ratio")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s [-alov] path\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	if *v {
		fmt.Printf("%s %s 🔪\n", os.Args[0], aspeq.Version)
		os.Exit(0)
	}

	if *l {
		for _, ar := range aspeq.Ratios {
			fmt.Printf("%s: %dx%d\n", ar.Name, ar.X, ar.Y)
		}
		os.Exit(0)
	}

	if flag.NArg() < 1 {
		flag.Usage()
		return
	}

	var ratio *aspeq.AspectRatio
	if *a == "" {
		var err error
		ratio, err = aspeq.FromPath(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
	} else {
		for _, ar := range aspeq.Ratios {
			if ar.Name == *a {
				ratio = ar
			}
		}
	}
	if ratio == nil {
		log.Fatalf("unknown aspect ratio %s\n", *a)
	}

	img, err := aspeq.CropPath(flag.Arg(0), ratio)

	if err != nil {
		log.Fatal(err)
	}

	if *o == "" {
		bounds := img.Bounds()
		fmt.Printf("%dx%d\n", bounds.Dx(), bounds.Dy())
	} else {
		file, err := os.Create(*o)
		if err != nil {
			log.Fatalf("error creating file: %s\n", err)
		}
		defer file.Close()

		err = jpeg.Encode(file, img, nil)
		if err != nil {
			log.Fatal(err)
		}
	}
}
