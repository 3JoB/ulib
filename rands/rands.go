package rands

import (
	"lukechampine.com/frand"
	"pgregory.net/rand"
)

func Rands(n []int, num int) []int {
	idxs := rand.Perm(len(n))[:num]
	result := make([]int, 0, num)
	for _, idx := range idxs {
		result = append(result, n[idx])
	}
	return result
}

func CRands(n []int, num int) []int {
	idxs := frand.Perm(len(n))[:num]
	result := make([]int, 0, num)
	for _, idx := range idxs {
		result = append(result, n[idx])
	}
	return result
}

func Rand64(n []int64, num int) []int64 {
	idxs := rand.Perm(len(n))[:num]
	result := make([]int64, 0, num)
	for _, idx := range idxs {
		result = append(result, n[idx])
	}
	return result
}

func CRand64(n []int64, num int) []int64 {
	idxs := frand.Perm(len(n))[:num]
	result := make([]int64, 0, num)
	for _, idx := range idxs {
		result = append(result, n[idx])
	}
	return result
}

func RandString(n []string, num int) []string {
	idxs := rand.Perm(len(n))[:num]
	result := make([]string, 0, num)
	for _, idx := range idxs {
		result = append(result, n[idx])
	}
	return result
}

func CRandString(n []string, num int) []string {
	idxs := frand.Perm(len(n))[:num]
	result := make([]string, 0, num)
	for _, idx := range idxs {
		result = append(result, n[idx])
	}
	return result
}
