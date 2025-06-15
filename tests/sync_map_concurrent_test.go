package cmapv2tests

import (
	"bytes"
	"sync"
	"testing"
)

func TestSyncMapSmallConcurrentOps(t *testing.T) {
	syncMap := sync.Map{}

	t.Run("test insert", func(t *testing.T) {
		runWithWorkers(smallInputSize, workerCount, func(idx int) {
			key, val := generateKeyVal64(idx)
			syncMap.Store(string(key), val)
		})
	})

	t.Run("test retrieve", func(t *testing.T) {
		runWithWorkers(smallInputSize, workerCount, func(idx int) {
			key, val := generateKeyVal64(idx)
			value, _ := syncMap.Load(string(key))
			if !bytes.Equal(value.([]byte), val) {
				t.Errorf("actual value not equal to expected: actual(%s), expected(%s)", value, val)
			}
		})
	})

	t.Run("test delete", func(t *testing.T) {
		runWithWorkers(smallInputSize, workerCount, func(idx int) {
			key, _ := generateKeyVal64(idx)
			syncMap.Delete(string(key))
		})
	})

	t.Log("done")
}

func TestSyncMapLargeConcurrentOps(t *testing.T) {
	syncMap := sync.Map{}

	t.Run("test insert", func(t *testing.T) {
		runWithWorkers(largeInputSize, workerCount, func(idx int) {
			key, val := generateKeyVal64(idx)
			syncMap.Store(string(key), val)
		})
	})

	t.Run("test retrieve", func(t *testing.T) {
		runWithWorkers(largeInputSize, workerCount, func(idx int) {
			key, val := generateKeyVal64(idx)
			value, _ := syncMap.Load(string(key))
			if !bytes.Equal(value.([]byte), val) {
				t.Errorf("actual value not equal to expected: actual(%s), expected(%s)", value, val)
			}
		})
	})

	t.Run("test delete", func(t *testing.T) {
		runWithWorkers(largeInputSize, workerCount, func(idx int) {
			key, _ := generateKeyVal64(idx)
			syncMap.Delete(string(key))
		})
	})

	t.Log("done")
}
