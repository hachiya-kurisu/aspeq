package main

import (
	"blekksprut.net/aspeq"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Fprintln(os.Stderr, "aspeq [images]")
		return
	}

	for _, arg := range os.Args[1:] {
		ratio, err := aspeq.FromImage(arg)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Printf("%s: %s\n", arg, ratio.name)
	}
}
