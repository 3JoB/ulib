package bimap

// IBiMap represents a bidirectional map interface allowing key-value and value-key lookups with mutual exclusivity.
type IBiMap[K comparable, V comparable] interface {
	// Insert adds a key-value pair to the map. Returns true if successful, false if the map is immutable.
	Insert(k K, v V) bool
	// Exists checks if the specified key exists in the forward map. Returns true if the key exists.
	Exists(k K) bool
	// ExistsInverse checks if the specified value exists in the inverse map. Returns true if the value exists.
	ExistsInverse(k V) bool
	// Get retrieves the value corresponding to the given key. Returns the value and a bool indicating success.
	Get(k K) (V, bool)
	// GetInverse retrieves the key corresponding to the given value. Returns the key and a bool indicating success.
	GetInverse(v V) (K, bool)
	// Delete removes the entry associated with the given key. Returns true if successful, false if the map is immutable.
	Delete(k K) bool
	// DeleteInverse removes the entry associated with the given value. Returns true if successful, false if immutable.
	DeleteInverse(v V) bool
	// GetInverseMap returns a copy of the inverse map (value-to-key mapping) for external use.
	GetInverseMap() map[V]K
	// GetForwardMap returns a copy of the forward map (key-to-value mapping) for external use.
	GetForwardMap() map[K]V
	// MakeImmutable converts the map to an immutable state, preventing further insertions or modifications.
	MakeImmutable()
	// isMutable checks if the map is still in a mutable state. Returns true if mutable, false if immutable.
	isMutable() bool
	// Size returns the number of entries currently stored in the map.
	Size() int
	// Reset clears all entries in the map, resetting it to an empty state. No effect if the map is immutable.
	Reset()
}
