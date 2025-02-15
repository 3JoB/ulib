// Package bimap from: https://github.com/vishalkuo/bimap
package bimap

import (
	"maps"
)

type BaseBiMap[K comparable, V comparable] struct {
	immutable bool
	forward   map[K]V
	inverse   map[V]K
}

func NewBaseMap[K comparable, V comparable]() IBiMap[K, V] {
	return &BaseBiMap[K, V]{
		forward:   make(map[K]V),
		inverse:   make(map[V]K),
		immutable: false,
	}
}

func NewBaseMapFromMap[K comparable, V comparable](f map[K]V) IBiMap[K, V] {
	size := len(f)

	biMap := &BaseBiMap[K, V]{
		forward:   make(map[K]V, size),
		inverse:   make(map[V]K, size),
		immutable: false,
	}

	for k, v := range f {
		biMap.inverse[v] = k
		biMap.forward[k] = v
	}
	return biMap
}

func NewImmutableBaseMap[K comparable, V comparable](f map[K]V) IBiMap[K, V] {
	biMap := NewBaseMapFromMap(f)
	biMap.MakeImmutable()
	return biMap
}

func (b *BaseBiMap[K, V]) isMutable() bool {
	return !b.immutable
}

func (b *BaseBiMap[K, V]) Insert(k K, v V) bool {
	if b.immutable {
		// Cannot modify immutable map
		return false
	}

	if r, ok := b.forward[k]; ok {
		delete(b.inverse, r)
	}

	b.forward[k] = v
	b.inverse[v] = k
	return true
}

func (b *BaseBiMap[K, V]) Exists(k K) bool {
	_, ok := b.forward[k]
	return ok
}

func (b *BaseBiMap[K, V]) ExistsInverse(k V) bool {
	_, ok := b.inverse[k]
	return ok
}

func (b *BaseBiMap[K, V]) Get(k K) (V, bool) {
	v, ok := b.forward[k]
	return v, ok
}

func (b *BaseBiMap[K, V]) GetInverse(v V) (K, bool) {
	r, ok := b.inverse[v]
	return r, ok
}

func (b *BaseBiMap[K, V]) Delete(k K) bool {
	if b.immutable {
		// Cannot modify immutable map
		return false
	}

	val, ok := b.Get(k)
	if !ok {
		return ok
	}

	delete(b.inverse, val)
	delete(b.forward, k)
	return true
}

func (b *BaseBiMap[K, V]) DeleteInverse(v V) bool {
	if b.immutable {
		// Cannot modify immutable map
		return false
	}

	key, ok := b.GetInverse(v)
	if !ok {
		return ok
	}

	delete(b.inverse, v)
	delete(b.forward, key)
	return true
}

func (b *BaseBiMap[K, V]) Size() int {
	return len(b.forward)
}

func (b *BaseBiMap[K, V]) MakeImmutable() {
	b.immutable = true
}

func (b *BaseBiMap[K, V]) GetInverseMap() map[V]K {
	m := make(map[V]K, len(b.inverse))
	maps.Copy(m, b.inverse)
	return m
}

func (b *BaseBiMap[K, V]) GetForwardMap() map[K]V {
	m := make(map[K]V, len(b.forward))
	maps.Copy(m, b.forward)
	return m
}

func (b *BaseBiMap[K, V]) Reset() {
	clear(b.forward)
	clear(b.inverse)
}
