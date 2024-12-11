# aspeq

determine closest "standard" aspect ratio of a rectangle or image

## currently defined ratios

...

## from the command line

```sh
% aspeq *.jpeg
1.66.jpeg: super16
1.77.jpeg: sixteen-nine
2.35.jpeg: cinemascope
```

## basic usage

```go
package main

import (
  "fmt"
  "log"
  "blekksprut.net/aspeq"
)

func main() {
  ratio := aspeq.FromWidthHeight(320, 240)
  fmt.Println(ratio.Xy()) // prints "4:3"

  ar, err := aspeq.FromImage("1.66.jpeg") // a 40:24 image
  if err != nil {
    log.Fatal(err)
    return
  }
  fmt.Println(ar.Name) // prints "super16"
}
```
