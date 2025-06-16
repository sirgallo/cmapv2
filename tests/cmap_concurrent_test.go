package cmapv2tests

import (
	"bytes"
	"testing"

	"github.com/sirgallo/cmapv2"
)

func TestCMapSmallConcurrentOps(t *testing.T) {
	cMap := cmap.NewMap(totalShards)

	t.Run("test insert", func(t *testing.T) {
		runWithWorkers(smallInputSize, workerCount, func(idx int) {
			key, val := generateKeyVal64(idx)
			cMap.Put(key, val)
		})
	})

	t.Run("test retrieve", func(t *testing.T) {
		runWithWorkers(smallInputSize, workerCount, func(idx int) {
			key, val := generateKeyVal64(idx)
			value := cMap.Get(key)
			if !bytes.Equal(value, val) {
				t.Errorf("actual value not equal to expected: actual(%s), expected(%s)", value, val)
			}
		})
	})

	t.Run("test delete", func(t *testing.T) {
		runWithWorkers(smallInputSize, workerCount, func(idx int) {
			key, _ := generateKeyVal64(idx)
			cMap.Delete(key)
		})
	})

	t.Log("done")
}

func TestCMapLargeConcurrentOps(t *testing.T) {
	cMap := cmap.NewMap(totalShards)

	t.Run("test insert", func(t *testing.T) {
		runWithWorkers(largeInputSize, workerCount, func(idx int) {
			key, val := generateKeyVal64(idx)
			cMap.Put(key, val)
		})
	})

	t.Run("test retrieve", func(t *testing.T) {
		runWithWorkers(largeInputSize, workerCount, func(idx int) {
			key, val := generateKeyVal64(idx)
			value := cMap.Get(key)
			if !bytes.Equal(value, val) {
				t.Errorf("actual value not equal to expected: actual(%s), expected(%s)", value, val)
			}
		})
	})

	t.Run("test delete", func(t *testing.T) {
		runWithWorkers(largeInputSize, workerCount, func(idx int) {
			key, _ := generateKeyVal64(idx)
			cMap.Delete(key)
		})
	})

	t.Log("done")
}
