package rands

import (
	"lukechampine.com/frand"
	"pgregory.net/rand"
)

func Rands[T any](n []T, num int) []T {
	idxs := rand.Perm(len(n))[:num]
	result := make([]T, 0, num)
	for _, idx := range idxs {
		result = append(result, n[idx])
	}
	return result
}

func CRands[T any](n []T, num int) []T {
	idxs := frand.Perm(len(n))[:num]
	result := make([]T, 0, num)
	for _, idx := range idxs {
		result = append(result, n[idx])
	}
	return result
}
