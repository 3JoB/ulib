// Package bimap from: https://github.com/vishalkuo/bimap
package bimap

import (
	"github.com/cornelk/hashmap"
)

type hashable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
}

type HashBiMap[K hashable, V hashable] struct {
	immutable bool
	forward   *hashmap.Map[K, V]
	inverse   *hashmap.Map[V, K]
}

func NewHashMap[K hashable, V hashable]() IBiMap[K, V] {
	return &HashBiMap[K, V]{forward: hashmap.New[K, V](), inverse: hashmap.New[V, K](), immutable: false}
}

func NewHashMapFromMap[K hashable, V hashable](f map[K]V) IBiMap[K, V] {
	biMap := &HashBiMap[K, V]{forward: hashmap.New[K, V](), inverse: hashmap.New[V, K](), immutable: false}
	for k, v := range f {
		biMap.forward.Set(k, v)
		biMap.inverse.Set(v, k)
	}
	return biMap
}

func NewImmutableHashMap[K hashable, V hashable](f map[K]V) IBiMap[K, V] {
	biMap := NewHashMapFromMap(f)
	biMap.MakeImmutable()
	return biMap
}

func (b *HashBiMap[K, V]) Insert(k K, v V) bool {
	if b.immutable {
		// Cannot modify immutable map
		return false
	}

	if r, ok := b.forward.Get(k); ok {
		b.inverse.Del(r)
	}

	b.forward.Set(k, v)
	b.inverse.Set(v, k)
	return true
}

func (b *HashBiMap[K, V]) isMutable() bool {
	return !b.immutable
}

func (b *HashBiMap[K, V]) Exists(k K) bool {
	_, ok := b.forward.Get(k)
	return ok
}

func (b *HashBiMap[K, V]) ExistsInverse(k V) bool {
	_, ok := b.inverse.Get(k)
	return ok
}

func (b *HashBiMap[K, V]) Get(k K) (V, bool) {
	return b.forward.Get(k)
}

func (b *HashBiMap[K, V]) GetInverse(v V) (K, bool) {
	return b.inverse.Get(v)
}

func (b *HashBiMap[K, V]) Delete(k K) bool {
	if b.immutable {
		// Cannot modify immutable map
		return false
	}

	val, ok := b.Get(k)
	if !ok {
		return ok
	}

	b.forward.Del(k)
	b.inverse.Del(val)
	return true
}

func (b *HashBiMap[K, V]) DeleteInverse(v V) bool {
	if b.immutable {
		// Cannot modify immutable map
		return false
	}

	key, ok := b.GetInverse(v)
	if !ok {
		return ok
	}

	b.inverse.Del(v)
	b.forward.Del(key)
	return true
}

func (b *HashBiMap[K, V]) Size() int {
	return b.forward.Len()
}

func (b *HashBiMap[K, V]) MakeImmutable() {
	b.immutable = true
}

func (b *HashBiMap[K, V]) GetInverseMap() map[V]K {
	m := make(map[V]K)
	b.inverse.Range(func(v V, k K) bool {
		m[v] = k
		return true
	})
	return m
}

func (b *HashBiMap[K, V]) GetForwardMap() map[K]V {
	m := make(map[K]V)
	b.forward.Range(func(k K, v V) bool {
		m[k] = v
		return true
	})
	return m
}

func (b *HashBiMap[K, V]) Reset() {
	b.forward.Range(func(k K, v V) bool {
		b.inverse.Del(v)
		b.forward.Del(k)
		return true
	})
}
