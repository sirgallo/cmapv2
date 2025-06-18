package cmap_benchmark

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/sirgallo/cmapv2"
)

var mixedKeys = make([][]byte, 1000000)
var sKeys = make([]string, 1000000)

func init() {
	for i := range mixedKeys {
		s := strconv.Itoa(rand.Int())
		mixedKeys[i] = []byte(s)
		sKeys[i] = s
	}
}

func BenchmarkMap(b *testing.B) {
	m := cmap.NewMap(4096)
	b.RunParallel(func(pb *testing.PB) {
		rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
		for pb.Next() {
			op := rnd.Intn(3)
			k := mixedKeys[rnd.Intn(len(mixedKeys))]
			switch op {
			case 0: m.Put(k, k)
			case 1: m.Get(k)
			case 2: m.Delete(k)
			}
		}
	})
}

func BenchmarkSyncMap(b *testing.B) {
	m := sync.Map{}
	b.RunParallel(func(pb *testing.PB) {
		rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
		for pb.Next() {
			op := rnd.Intn(3)
			k := sKeys[rnd.Intn(len(sKeys))]
			switch op {
			case 0: m.Store(k, k)
			case 1: m.Load(k)
			case 2: m.Delete(k)
			}
		}
	})
}
