package cmap

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
	"unsafe"
)

func TestCMapRace(t *testing.T) {
	t.Run("test race on same key update/read", func(t *testing.T) {
		m := NewMap()
		const N = 100000
		key := []byte("hotkey")
		m.Put(key, []byte("0")) // start with an initial value so Get() always succeeds

		var wg sync.WaitGroup
		for range workerCount {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for i := range N {
					m.Put(key, []byte(strconv.Itoa(i)))
					runtime.Gosched()
				}
			}()
		}

		for range workerCount {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for range N {
					m.Get(key)
					runtime.Gosched()
				}
			}()
		}
		wg.Wait()
	})

	t.Run("test race delete branch", func(t *testing.T) {
		m := NewMap() // single shard for maximum contention
		// build a little bucket of 32 leaves in one slot
		prefix := []byte{0x00}
		for i := range 32 {
			k := append(prefix, byte(i))
			m.Put(k, k)
		}

		victim := append(prefix, byte(16))

		const M = 100000
		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			for range M {
				m.Delete(victim)
				runtime.Gosched()
			}
		}()

		go func() {
			defer wg.Done()
			for range M {
				_ = m.Get(victim)
				runtime.Gosched()
			}
		}()

		wg.Wait()
	})

	t.Run("test race snapshot vs live", func(t *testing.T) {
		m := &cMap{
			bitChunkSize: 5,
			hashChunks:   6,
			root: unsafe.Pointer(&node{
				isLeaf:   false,
				bitmap:   0,
				children: []*node{},
			}),
		}
		const K = 1000
		for i := range K {
			k := []byte(fmt.Sprintf("k-%d", i))
			m.Put(k, k)
		}

		oldRoot := atomic.LoadPointer(&m.root)
		oldMap := &cMap{
			bitChunkSize: m.bitChunkSize,
			hashChunks:   m.hashChunks,
			root:         oldRoot,
		}

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			for i := range K {
				m.Delete([]byte(fmt.Sprintf("k-%d", i)))
				runtime.Gosched()
			}
		}()

		go func() {
			defer wg.Done()
			for i := range K {
				_ = oldMap.Get([]byte(fmt.Sprintf("k-%d", i)))
				runtime.Gosched()
			}
		}()
		wg.Wait()
	})

	t.Log("Done")
}
