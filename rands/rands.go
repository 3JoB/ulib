package rands

import (
	"lukechampine.com/frand"
	"pgregory.net/rand"
)

func Rands[T any](n []T, num int) []T {
	if len(n) < 1 && len(n) <= num {
		return nil
	}
	result := make([]T, 0, num)
	for _, idx := range rand.Perm(len(n))[:num] {
		result = append(result, n[idx])
	}
	return result
}

func CRands[T any](n []T, num int) []T {
	if len(n) < 1 && len(n) <= num {
		return nil
	}
	result := make([]T, 0, num)
	for _, idx := range frand.Perm(len(n))[:num] {
		result = append(result, n[idx])
	}
	return result
}
