# cmapv2

## installation

In your `Go` project main directory (where the `go.mod` file is located)
```bash
go get github.com/sirgallo/cmapv2
go mod tidy
```

Make sure to run go mod tidy to install dependencies.


## usage

```go
package main

import (
  "github.com/sirgallo/cmapv2"
)

func main() {
  // initialize c map
  cMap := cmap.NewCMap()

  // insert key/val pair
  cMap.Put([]byte("hi"), []byte("world"))

  // retrieve value for key
  val := cMap.Get([]byte("hi"))

  // delete key/val pair
  cMap.Delete([]byte("hi"))
}
```

## tests

```bash
go test -v ./tests
```

## godoc

For in depth definitions of types and functions, `godoc` can generate documentation from the formatted function comments. If `godoc` is not installed, it can be installed with the following:
```bash
go install golang.org/x/tools/cmd/godoc
```

To run the `godoc` server and view definitions for the package:
```bash
godoc -http=:6060
```

Then, in your browser, navigate to:
```
http://localhost:6060/pkg/github.com/sirgallo/cmap/
```

## sources

[CMap](./docs/CMap.md)

[Murmur](./docs/Murmur.md)