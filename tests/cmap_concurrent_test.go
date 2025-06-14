package cmapv2tests

import (
	"bytes"
	"fmt"
	"sync"
	"testing"

	cmap "github.com/sirgallo/cmapv2"
)

var sKeyValPairs, lKeyValPairs []KeyVal
var lKeyValChan chan KeyVal
var fillArrWG sync.WaitGroup

func init() {
	sKeyValPairs = initSmall()
	lKeyValPairs = initLarge()
}

func initSmall() []KeyVal {
	sInputSize := 1000000
	sKeyValPairs = make([]KeyVal, sInputSize)
	for idx := range sKeyValPairs {
		randomBytes, _ := generateRandomBytes(128)
		sKeyValPairs[idx] = KeyVal{Key: randomBytes, Value: randomBytes}
	}

	return sKeyValPairs
}

func initLarge() []KeyVal {
	lInputSize := 10000000
	lKeyValPairs = make([]KeyVal, lInputSize)
	lKeyValChan = make(chan KeyVal, lInputSize)

	for range lKeyValPairs {
		fillArrWG.Add(1)
		go func() {
			defer fillArrWG.Done()
			randomBytes, _ := generateRandomBytes(32)
			lKeyValChan <- KeyVal{Key: randomBytes, Value: randomBytes}
		}()
	}

	fillArrWG.Wait()
	fmt.Println("filled random key val pairs chan with size1:", lInputSize)

	for idx := range lKeyValPairs {
		keyVal := <-lKeyValChan
		lKeyValPairs[idx] = keyVal
	}

	return lKeyValPairs
}

func TestCMapSmallConcurrentOps(t *testing.T) {
	cMap := cmap.NewShardedMap()

	t.Run("test insert", func(t *testing.T) {
		runWithWorkers(sKeyValPairs, workerCount, func(val KeyVal) {
			cMap.Put(val.Key, val.Value)
		})
	})

	t.Run("test retrieve", func(t *testing.T) {
		runWithWorkers(sKeyValPairs, workerCount, func(val KeyVal) {
			value := cMap.Get(val.Key)
			// t.Logf("actual: %s, expected: %s", value, val.Value)
			if !bytes.Equal(value, val.Value) {
				t.Errorf("actual value not equal to expected: actual(%s), expected(%s)", value, val.Value)
			}
		})
	})

	t.Run("test delete", func(t *testing.T) {
		runWithWorkers(sKeyValPairs, workerCount, func(val KeyVal) {
			cMap.Delete(val.Key)
		})
	})

	t.Log("done")
}

func TestCMapLargeConcurrentOps(t *testing.T) {
	cMap := cmap.NewShardedMap()

	t.Run("test insert", func(t *testing.T) {
		runWithWorkers(lKeyValPairs, workerCount, func(val KeyVal) {
			cMap.Put(val.Key, val.Value)
		})
	})

	t.Run("test retrieve", func(t *testing.T) {
		runWithWorkers(lKeyValPairs, workerCount, func(val KeyVal) {
			value := cMap.Get(val.Key)
			// t.Logf("actual: %s, expected: %s", value, val.Value)
			if !bytes.Equal(value, val.Value) {
				t.Errorf("actual value not equal to expected: actual(%s), expected(%s)", value, val.Value)
			}
		})
	})

	t.Run("test delete", func(t *testing.T) {
		runWithWorkers(lKeyValPairs, workerCount, func(val KeyVal) {
			cMap.Delete(val.Key)
		})
	})

	t.Log("done")
}
