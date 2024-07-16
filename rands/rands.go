package rands

import (
	v1rand "math/rand"
	v2rand "math/rand/v2"

	"lukechampine.com/frand"
	"pgregory.net/rand"
)

func RandStd[T any](n []T, num int) []T {
	un := len(n)
	if un < 1 || num <= 0 || un <= num {
		return nil
	}
	result := make([]T, 0, num)
	for _, idx := range v2rand.Perm(un)[:num] {
		result = append(result, n[idx])
	}
	return result
}

func RandStdV1[T any](n []T, num int) []T {
	un := len(n)
	if un < 1 || num <= 0 || un <= num {
		return nil
	}
	result := make([]T, 0, num)
	for _, idx := range v1rand.Perm(un)[:num] {
		result = append(result, n[idx])
	}
	return result
}

func Rands[T any](n []T, num int) []T {
	un := len(n)
	if un < 1 || num <= 0 || un <= num {
		return nil
	}
	result := make([]T, 0, num)
	for _, idx := range rand.Perm(un)[:num] {
		result = append(result, n[idx])
	}
	return result
}

func CRands[T any](n []T, num int) []T {
	un := len(n)
	if un < 1 || num <= 0 || un <= num {
		return nil
	}
	result := make([]T, 0, num)
	for _, idx := range frand.Perm(un)[:num] {
		result = append(result, n[idx])
	}
	return result
}
