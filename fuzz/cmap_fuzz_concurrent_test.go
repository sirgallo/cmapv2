package cmap_fuzz

import (
	"sync"
	"testing"

	"github.com/sirgallo/cmapv2"
)

func FuzzConcurrentOps(f *testing.F) {
	f.Add([]byte("foo"), []byte("bar"))
	f.Add([]byte("baz"), []byte("qux"))

	f.Fuzz(func(t *testing.T, key, val []byte) {
		m := cmap.NewMap()
		var wg sync.WaitGroup
		wg.Add(3)
		go func() { defer wg.Done(); m.Put(key, val) }()
		go func() { defer wg.Done(); m.Get(key) }()
		go func() { defer wg.Done(); m.Delete(key) }()
		wg.Wait()
	})
}
