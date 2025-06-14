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

  // ===== OR

  // initialize sharded c map
  sMap := cmap.NewShardedMap(16)
  sMap.Put([]byte("hi"), []byte("world"))
  val := sMap.Get([]byte("hi"))
  sMap.Delete([]byte("hi"))
}
```

## tests

```bash
go test -v ./tests
```

or bench:
```bash
go test -v -bench=. -benchmem -cpuprofile cpu.prof -memprofile mem.prof ./tests
```

and check results:
```bash
go tool pprof -http=:8080 tests.test cpu.prof
go tool pprof -http=:8080 tests.test mem.prof
```

## sources

[CMap](./docs/CMap.md)

[Murmur](./docs/Murmur.md)

[Tests](./docs/Tests.md)