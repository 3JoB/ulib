package timings

import (
	"runtime"
	"sync"
)

// AttrSlice returns the optimal number of slices based on the input slice length and the number of CPU cores available.
func AttrSlice[T any](v []T) int {
	length := len(v)
	numCPU := runtime.NumCPU()
	if length == 0 || length < numCPU {
		return 1
	}
	return numCPU
}

// ParallelForEach executes a provided function concurrently on each element of a slice using multiple goroutines.
// The number of goroutines is determined by the numThreads parameter.
// If the slice is empty or numThreads is less than or equal to zero, the function returns immediately.
// The fn parameter is a function that takes the index and value of an element in the slice as arguments.
func ParallelForEach[T any](slice []T, numThreads int, fn func(index int, value T)) {
	length := len(slice)
	if length == 0 || numThreads <= 0 {
		return
	}
	chunkSize := (length + numThreads - 1) / numThreads
	var wg sync.WaitGroup
	for i := 0; i < numThreads && i*chunkSize < length; i++ {
		wg.Add(1)
		go processChunk(slice, i*chunkSize, chunkSize, length, fn, &wg)
	}
	wg.Wait()
}

// processChunk processes a chunk of a slice concurrently, applying the provided function to each element in the chunk.
// It starts at the given index and processes up to the specified chunk size or until the slice length is reached.
// The WaitGroup ensures proper synchronization of the concurrent goroutines that use this function.
func processChunk[T any](slice []T, start int, chunkSize int, length int, fn func(index int, value T), wg *sync.WaitGroup) {
	defer wg.Done()
	end := start + chunkSize
	if end > length {
		end = length
	}
	for j := start; j < end; j++ {
		fn(j, slice[j])
	}
}
