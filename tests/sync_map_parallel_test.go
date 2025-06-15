package cmapv2tests

import (
	"bytes"
	"sync"
	"testing"
)

var parallelSyncMap sync.Map

func init() {
	parallelSyncMap = sync.Map{}
	runWithWorkers(smallInputSize, workerCount, func(idx int) {
		key, val := generateKeyVal64(idx)
		parallelSyncMap.Store(string(key), val)
	})
}

func TestParallelSyncMapReadWrites(t *testing.T) {
	t.Run("test init key val pairs in map", func(t *testing.T) {
		t.Parallel()
		runWithWorkers(smallInputSize, workerCount, func(idx int) {
			key, val := generateKeyVal64(idx)
			value, _ := parallelSyncMap.Load(string(key))
			if !bytes.Equal(value.([]byte), val) {
				t.Errorf("actual value not equal to expected: actual(%s), expected(%s)", value, val)
			}
		})
	})

	t.Run("test write new key vals in map", func(t *testing.T) {
		t.Parallel()
		runWithWorkers(smallInputSize, workerCount, func(idx int) {
			key, val := generateKeyVal64(idx + smallInputSize)
			parallelSyncMap.Store(string(key), val)
		})
	})

	t.Log("Done")
}
