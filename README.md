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
  "blekksprut.net/aspeq"
)

func main() {
  ratio := aspeq.FromWidthHeight(320, 240)
  fmt.Println(ratio)

  imgratio, err := aspeq.FromImage("test.jpg")
  if err != nil {
    log.Fatal(err)
    return
  }
  fmt.Println(imgratio)
}
```

