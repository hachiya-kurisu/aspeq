# aspeq

aspeq finds the closest "standard" aspect ratio from an image file, or from relative or absolute dimensions (width and height)

## defined ratios

* insta (9:16)
* classic (2:3)
* instax (3:4)
* square (1:1)
* movietone (1.19:1)
* four-thirds (4:3)
* academy (1.375:1)
* leica (3:2)
* super16 (5:3)
* sixteen-nine (16:9)
* flat (1.85:1)
* univisium (2:1)
* cinemascope (2.35:1)
* cinerama (2.59:1)
* widelux (3:1)
* polyvision (4:1)
* circle-vision (12:1)

## command line

```
$ aspeq *.jpeg
1.66.jpeg: super16
1.77.jpeg: sixteen-nine
2.35.jpeg: cinemascope

$ aspeq -x 1.66.jpeg
1.66.jpeg: 5:3

$ aspeq -o 1.66.jpeg
1.66.jpeg: landscape
```

slasher lets you crop images to a specific aspect ratio

```
$ slasher -o leica.jpg -a leica super16.jpg
$ slasher -o closest.jpg weird.jpg # crop to closest aspect ratio
$ slasher -a cinerama cat.jpg # prints cinerama dimensions
1600x617
$ slasher -l # lists available aspect ratios
```


## go

```
package main

import (
  "fmt"
  "blekksprut.net/aspeq"
)

func main() {
  ratio := aspeq.Match(320, 240)
  fmt.Println(ratio.Xy()) // prints "4:3"

  ar, err := aspeq.FromPath("1.66.jpeg") // a 40:24 image
  if err != nil {
    panic(err)
  }
  fmt.Println(ar.Name) // prints "super16"
}
```
