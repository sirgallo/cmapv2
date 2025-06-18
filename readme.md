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
  cMap := cmap.NewMap()

  // insert key/val pair
  cMap.Put([]byte("hi"), []byte("world"))

  // retrieve value for key
  val := cMap.Get([]byte("hi"))

  // delete key/val pair
  cMap.Delete([]byte("hi"))

  // ===== OR

  // initialize sharded c map with number of shards
  sMap := cmap.NewMap(16)
  sMap.Put([]byte("hi"), []byte("world"))
  val := sMap.Get([]byte("hi"))
  sMap.Delete([]byte("hi"))
}
```

## tests

```bash
go test -v ./
```

or with race:
```bash
go test -v -race ./
```

or bench:
```bash
go test -v -bench=. -benchmem -cpuprofile cpu.prof -memprofile mem.prof -coverprofile=coverage.out ./
go tool cover -html=coverage.out -o cov.html
open cov.html
```

and check results:
```bash
go tool pprof -http=:8080 cmapv2.test cpu.prof
go tool pprof -http=:8080 cmapv2.test mem.prof
```

## benchmark

```bash
go test -v -bench=. -benchmem -cpuprofile cpu.prof -memprofile mem.prof ./benchmarks
```

or with race:
```bash
go test -race -v -bench=. -benchmem -cpuprofile cpu.prof -memprofile mem.prof ./benchmarks
```
## fuzz

```bash
go test -fuzz=FuzzConcurrentOps -fuzztime=30s ./fuzz 
```

## sources

[CMap](./docs/CMap.md)

[Murmur](./docs/Murmur.md)

[Tests](./docs/Tests.md)