// from: https://github.com/vishalkuo/bimap
package bimap

import (
	"github.com/cornelk/hashmap"
)

type hashable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
}

type BiMap[K hashable, V hashable] struct {
	immutable bool
	forward   *hashmap.Map[K, V]
	inverse   *hashmap.Map[V, K]
}

func New[K hashable, V hashable]() *BiMap[K, V] {
	return &BiMap[K, V]{forward: hashmap.New[K, V](), inverse: hashmap.New[V, K](), immutable: false}
}

func NewFromMap[K hashable, V hashable](f map[K]V) *BiMap[K, V] {
	biMap := New[K, V]()
	for k, v := range f {
		biMap.Insert(k, v)
	}
	return biMap
}

func (b *BiMap[K, V]) Insert(k K, v V) bool {
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

func (b *BiMap[K, V]) Exists(k K) bool {
	_, ok := b.forward.Get(k)
	return ok
}

func (b *BiMap[K, V]) ExistsInverse(k V) bool {
	_, ok := b.inverse.Get(k)
	return ok
}

func (b *BiMap[K, V]) Get(k K) (V, bool) {
	return b.forward.Get(k)
}

func (b *BiMap[K, V]) GetInverse(v V) (K, bool) {
	return b.inverse.Get(v)
}

func (b *BiMap[K, V]) Delete(k K) bool {
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

func (b *BiMap[K, V]) DeleteInverse(v V) bool {
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

func (b *BiMap[K, V]) Size() int {
	return b.forward.Len()
}

func (b *BiMap[K, V]) MakeImmutable() {
	b.immutable = true
}

func (b *BiMap[K, V]) GetInverseMap() map[V]K {
	m := make(map[V]K)
	b.inverse.Range(func(v V, k K) bool {
		m[v] = k
		return true
	})
	return m
}

func (b *BiMap[K, V]) GetForwardMap() map[K]V {
	m := make(map[K]V)
	b.forward.Range(func(k K, v V) bool {
		m[k] = v
		return true
	})
	return m
}

func (b *BiMap[K, V]) Reset() {
	b.forward.Range(func(k K, v V) bool {
		b.inverse.Del(v)
		b.forward.Del(k)
		return true
	})
}
