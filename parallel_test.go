package cmap

import (
	"bytes"
	"testing"
)

var parallelTestMap CMap

func init() {
	parallelTestMap = NewMap(totalShards)
	runWithWorkers(smallInputSize, workerCount, func(idx int) {
		key := generateKeyVal64(idx)
		parallelTestMap.Put(key, key)
	})
}

func TestParallelReadWrites(t *testing.T) {
	t.Run("test init key val pairs in map", func(t *testing.T) {
		t.Parallel()
		runWithWorkers(smallInputSize, workerCount, func(idx int) {
			key := generateKeyVal64(idx)
			value := parallelTestMap.Get(key)
			if !bytes.Equal(value, key) {
				t.Errorf("actual value not equal to expected: actual(%s), expected(%s)", value, key)
			}
		})
	})

	t.Run("test write new key vals in map", func(t *testing.T) {
		t.Parallel()
		runWithWorkers(smallInputSize, workerCount, func(idx int) {
			key := generateKeyVal64(idx + smallInputSize)
			parallelTestMap.Put(key, key)
		})
	})

	t.Log("Done")
}
