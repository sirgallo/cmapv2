package cmapv2tests

import (
	"bytes"
	"testing"

	"github.com/sirgallo/cmapv2"
)

var parallelTestMap cmap.CMap

func init() {
	parallelTestMap = cmap.NewShardedMap(totalShards)
	runWithWorkers(smallInputSize, workerCount, func(idx int) {
		key, val := generateKeyVal64(idx)
		parallelTestMap.Put(key, val)
	})
}

func TestParallelReadWrites(t *testing.T) {
	t.Run("test init key val pairs in map", func(t *testing.T) {
		t.Parallel()
		runWithWorkers(smallInputSize, workerCount, func(idx int) {
			key, val := generateKeyVal64(idx)
			value := parallelTestMap.Get(key)
			if !bytes.Equal(value, val) {
				t.Errorf("actual value not equal to expected: actual(%s), expected(%s)", value, val)
			}
		})
	})

	t.Run("test write new key vals in map", func(t *testing.T) {
		t.Parallel()
		runWithWorkers(smallInputSize, workerCount, func(idx int) {
			key, val := generateKeyVal64(idx + smallInputSize)
			parallelTestMap.Put(key, val)
		})
	})

	t.Log("Done")
}
