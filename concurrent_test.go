package cmap

import (
	"bytes"
	"testing"
)

func TestCMapConcurrentOps(t *testing.T) {
	cMap := NewMap(totalShards)

	t.Run("test put", func(t *testing.T) {
		runWithWorkers(largeInputSize, workerCount, func(idx int) {
			key := generateKeyVal64(idx)
			cMap.Put(key, key)
		})
	})

	t.Run("test retrieve", func(t *testing.T) {
		runWithWorkers(largeInputSize, workerCount, func(idx int) {
			key := generateKeyVal64(idx)
			value := cMap.Get(key)
			if !bytes.Equal(value, key) {
				t.Errorf("actual value not equal to expected: actual(%s), expected(%s)", value, key)
			}
		})
	})

	t.Run("test delete", func(t *testing.T) {
		runWithWorkers(largeInputSize, workerCount, func(idx int) {
			key := generateKeyVal64(idx)
			cMap.Delete(key)
		})
	})

	t.Log("done")
}

func TestCMapLargeConcurrentOps(t *testing.T) {
	cMap := NewMap(totalShards)

	t.Run("test concurrent small", func(t *testing.T) {
		t.Parallel()
		runWithWorkers(smallInputSize, workerCount, func(idx int) {
			key := generateKeyVal64(idx)
			val := generateKeyVal64(idx + largeInputSize)
			cMap.Put(key, val)
			value := cMap.Get(key)
			if !bytes.Equal(value, val) {
				t.Errorf("actual value not equal to expected: actual(%s), expected(%s)", value, val)
			}
		})
	})

	t.Run("test concurrent large", func(t *testing.T) {
		t.Parallel()
		runWithWorkers(largeInputSize, workerCount, func(idx int) {
			key := generateKeyVal64(idx)
			cMap.Put(key, key)
			value := cMap.Get(key)
			if !bytes.Equal(value, key) {
				t.Errorf("actual value not equal to expected: actual(%s), expected(%s)", value, key)
			}
			cMap.Delete(key)
		})
	})

	t.Log("done")
}
