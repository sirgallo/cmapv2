package cmapv2tests

import (
	"crypto/rand"
	"sync"
)

var workerCount int = 1

type KeyVal struct {
	Key   []byte
	Value []byte
}

func generateRandomBytes(length int) ([]byte, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}
	return randomBytes, nil
}

func runWithWorkers(pairs []KeyVal, workerCount int, fn func(KeyVal)) {
	jobs := make(chan KeyVal, len(pairs))
	var wg sync.WaitGroup
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for val := range jobs {
				fn(val)
			}
		}()
	}

	for _, pair := range pairs {
		jobs <- pair
	}

	close(jobs)
	wg.Wait()
}
