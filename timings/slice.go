package timings

import (
	"runtime"
	"sync"
)

func AttrSlice[T any](v []T) int {
	length := len(v)
	if length == 0 {
		return 1
	}

	numCPU := runtime.NumCPU()
	if length < numCPU {
		return 1
	}

	return numCPU
}

func AttrRange[T any](slice []T, numThreads int, fn func(v int, b T)) {
	length := len(slice)
	if length == 0 || numThreads <= 0 {
		return
	}

	chunkSize := (length + numThreads - 1) / numThreads

	var wg sync.WaitGroup

	for i := 0; i < numThreads && i*chunkSize < length; i++ {
		wg.Add(1)
		go func(start int) {
			defer wg.Done()
			end := start + chunkSize
			if end > length {
				end = length
			}
			for j := start; j < end; j++ {
				fn(j, slice[j])
			}
		}(i * chunkSize)
	}

	wg.Wait()
}
