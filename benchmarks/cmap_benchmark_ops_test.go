package cmap_benchmark

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/sirgallo/cmapv2"
)

var opKeys = make([][]byte, 1000000)

func init() {
	for i := range opKeys {
		s := strconv.Itoa(rand.Int())
		opKeys[i] = []byte(s)
	}
}

func BenchmarkPut(b *testing.B) {
  m := cmap.NewMap(4096)
  b.RunParallel(func(pb *testing.PB) {
    rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
    for pb.Next() {
			k := opKeys[rnd.Intn(len(mixedKeys))]
      m.Put(k, k)
    }
  })
}

func BenchmarkGet(b *testing.B) {
  m := cmap.NewMap(4096)
  for _, k := range opKeys {
    m.Put(k, k)
  }

  b.ResetTimer()
  b.RunParallel(func(pb *testing.PB) {
    rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
    for pb.Next() {
      m.Get(opKeys[rnd.Intn(len(opKeys))])
    }
  })
}

func BenchmarkDelete(b *testing.B) {
	m := cmap.NewMap(4096)
  for _, k := range opKeys {
    m.Put(k, k)
  }

	b.ResetTimer()
  b.RunParallel(func(pb *testing.PB) {
    rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
    for pb.Next() {
      m.Delete(opKeys[rnd.Intn(len(opKeys))])
    }
  })
}
