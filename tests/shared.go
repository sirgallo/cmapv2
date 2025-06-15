package cmapv2tests

import (
	"encoding/binary"
	"sync"
)

var largeInputSize int = 10000000
var smallInputSize int = 1000000
var totalShards int = 2048
var workerCount int = 3

func generateKeyVal64(index int) ([]byte, []byte) {
	key := make([]byte, 64)
	val := make([]byte, 64)
	for idx := range 8 {
		offset := idx * 8
		binary.LittleEndian.PutUint64(key[offset:], uint64(index*(idx+1)))
		binary.LittleEndian.PutUint64(val[offset:], uint64((index+1)*(idx+2)))
	}

	return key, val
}

func runWithWorkers(total int, workerCount int, fn func(int)) {
	jobs := make(chan int, total)
	var wg sync.WaitGroup
	for range workerCount {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for idx := range jobs {
				fn(idx)
			}
		}()
	}

	for idx := range total {
		jobs <- idx
	}

	close(jobs)
	wg.Wait()
}
