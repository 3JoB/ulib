package maputil

import (
	"maps"

	"github.com/cornelk/hashmap"
)

type hashable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
}

// Copy creates and returns a new map by copying all the key-value pairs from the source map.
func Copy[M ~map[K]V, K comparable, V any](src M) map[K]V {
	m := make(map[K]V, len(src))
	for k, v := range src {
		m[k] = v
	}
	return m
}

// Clone creates and returns a shallow copy of the given map, preserving the original map's key-value pairs.
func Clone[M ~map[K]V, K comparable, V any](src M) map[K]V {
	return maps.Clone(src)
}

// MapToHashMap converts a generic map of type M into a hashmap.Map and returns a pointer to the resulting hashmap.
func MapToHashMap[M ~map[K]V, K hashable, V any](src M) *hashmap.Map[K, V] {
	m := hashmap.New[K, V]()
	for k, v := range src {
		m.Set(k, v)
	}
	return m
}

// HashMapToMap converts a hashmap.Map to a native Go map with the same key-value pairs.
// Accepts a pointer to hashmap.Map and returns a map of the specified type.
// Ensures type compatibility using type parameters and constraints for key and value types.
// Utilizes Range to iterate over the hashmap and populate the native map.
func HashMapToMap[M ~map[K]V, K hashable, V any](src *hashmap.Map[K, V]) map[K]V {
	m := make(M, src.Len())
	src.Range(func(k K, v V) bool {
		m[k] = v
		return true
	})
	return m
}
