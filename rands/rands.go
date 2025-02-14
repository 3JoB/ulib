package rands

import (
	v1rand "math/rand"
	v2rand "math/rand/v2"

	"lukechampine.com/frand"
	"pgregory.net/rand"
)

const (
	STD = iota
	STDV2
	FRAND
	PG
)

func Rand[T any](algorithm int, n []T, num int) []T {
	switch algorithm {
	case STD:
		return stdV1(n, num)
	case STDV2:
		return stdV2(n, num)
	case FRAND:
		return frandRand(n, num)
	case PG:
		return pgRand(n, num)
	default:
		return nil
	}
}

func stdV2[T any](n []T, num int) []T {
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

func stdV1[T any](n []T, num int) []T {
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

func pgRand[T any](n []T, num int) []T {
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

func frandRand[T any](n []T, num int) []T {
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
