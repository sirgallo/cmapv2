package cmap

import (
	"testing"
)

func TestMurmur(t *testing.T) {
	t.Run("test hash", func(t *testing.T) {
		key := []byte("hello")
		seed := uint32(1)
		hash := Murmur32(key, seed)
		t.Log("hash:", hash)
	})

	t.Run("test reseed hash", func(t *testing.T) {
		m := &cMap{
			bitChunkSize: 5,
			hashChunks:   6,
		}

		m.root.Store(&node{
			isLeaf:   false,
			bitmap:   0,
			children: []*node{},
		})

		key := []byte("hello")
		levels := make([]int, 17)
		totalLevels := 6
		chunkSize := 5
		hash := m.calculateHashForCurrentLevel(key, 0, 0)
		for idx := range levels {
			index := GetIndexForLevel(hash, chunkSize, idx, totalLevels)
			t.Logf("hash: %d, index: %d", hash, index)
			hash = m.calculateHashForCurrentLevel(key, hash, idx+1)
		}
	})

	t.Log("Done")
}
