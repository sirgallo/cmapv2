package cmapv2tests

import (
	"testing"

	"github.com/sirgallo/cmapv2"
)

func TestMurmur(t *testing.T) {
	t.Run("test hash", func(t *testing.T) {
		key := []byte("hello")
		seed := uint32(1)
		hash := cmap.Murmur32(key, seed)
		t.Log("hash:", hash)
	})

	t.Run("test reseed hash", func(t *testing.T) {
		key := []byte("hello")
		levels := make([]int, 17)
		totalLevels := 6
		chunkSize := 5
		hash := calculateHashForCurrentLevel(key, 0, 0)
		for idx := range levels {
			index := cmap.GetIndexForLevel(hash, chunkSize, idx, totalLevels)
			t.Logf("hash: %d, index: %d", hash, index)
			hash = calculateHashForCurrentLevel(key, hash, idx+1)
		}
	})

	t.Log("Done")
}

func calculateHashForCurrentLevel(key []byte, hash uint32, level int) uint32 {
	if level%6 == 0 || hash == 0 {
		currChunk := level / 6
		seed := uint32(currChunk + 1)
		return cmap.Murmur32(key, seed)
	}

	return hash
}
