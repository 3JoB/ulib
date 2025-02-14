package bimap

import (
	"testing"
)

func TestInsert(t *testing.T) {
	tests := []struct {
		name      string
		immutable bool
		key       int
		value     string
		expected  bool
	}{
		{"insert new pair", false, 1, "a", true},
		{"insert replace existing key", false, 1, "b", true},
		{"insert replace existing value inverse", false, 2, "b", true},
		{"insert immutable", true, 1, "a", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := New[int, string]()
			if tt.immutable {
				b.MakeImmutable()
			}
			res := b.Insert(tt.key, tt.value)
			if res != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, res)
			}
		})
	}
}

func TestExists(t *testing.T) {
	b := New[int, string]()
	b.Insert(1, "a")

	if !b.Exists(1) {
		t.Error("key should exist")
	}
	if b.Exists(2) {
		t.Error("key should not exist")
	}
}

func TestExistsInverse(t *testing.T) {
	b := New[int, string]()
	b.Insert(1, "a")

	if !b.ExistsInverse("a") {
		t.Error("value should exist")
	}
	if b.ExistsInverse("b") {
		t.Error("value should not exist")
	}
}

func TestGet(t *testing.T) {
	b := New[int, string]()
	b.Insert(1, "a")

	if val, ok := b.Get(1); !ok || val != "a" {
		t.Error("expected to get value 'a', but got different result")
	}
	if _, ok := b.Get(2); ok {
		t.Error("expected not to get a value for non-existing key")
	}
}

func TestGetInverse(t *testing.T) {
	b := New[int, string]()
	b.Insert(1, "a")

	if key, ok := b.GetInverse("a"); !ok || key != 1 {
		t.Error("expected to get key '1', but got different result")
	}
	if _, ok := b.GetInverse("b"); ok {
		t.Error("expected not to get a key for non-existing value")
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		name      string
		immutable bool
		key       int
		expected  bool
	}{
		{"delete existing key", false, 1, true},
		{"delete non-existing key", false, 2, false},
		{"delete immutable map", true, 1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := New[int, string]()
			b.Insert(1, "a")
			if tt.immutable {
				b.MakeImmutable()
			}
			res := b.Delete(tt.key)
			if res != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, res)
			}
		})
	}
}

func TestDeleteInverse(t *testing.T) {
	tests := []struct {
		name      string
		immutable bool
		value     string
		expected  bool
	}{
		{"delete existing value", false, "a", true},
		{"delete non-existing value", false, "b", false},
		{"delete immutable map", true, "a", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := New[int, string]()
			b.Insert(1, "a")
			if tt.immutable {
				b.MakeImmutable()
			}
			res := b.DeleteInverse(tt.value)
			if res != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, res)
			}
		})
	}
}

func TestSize(t *testing.T) {
	b := New[int, string]()
	if size := b.Size(); size != 0 {
		t.Errorf("expected size to be 0, got %d", size)
	}
	b.Insert(1, "a")
	if size := b.Size(); size != 1 {
		t.Errorf("expected size to be 1, got %d", size)
	}
	b.Insert(2, "b")
	if size := b.Size(); size != 2 {
		t.Errorf("expected size to be 2, got %d", size)
	}
}

func TestMakeImmutable(t *testing.T) {
	b := New[int, string]()
	b.MakeImmutable()
	if !b.Insert(1, "a") {
		t.Error("expected insert to fail in immutable map")
	}
}

func TestReset(t *testing.T) {
	b := New[int, string]()
	b.Insert(1, "a")
	b.Insert(2, "b")
	b.Reset()

	if size := b.Size(); size != 0 {
		t.Errorf("expected size to be 0 after reset, got %d", size)
	}
}

func TestNewFromMap(t *testing.T) {
	m := map[int]string{1: "a", 2: "b"}
	b := NewFromMap(m)

	if size := b.Size(); size != len(m) {
		t.Errorf("expected size %d, got %d", len(m), size)
	}

	for k, v := range m {
		if val, ok := b.Get(k); !ok || val != v {
			t.Errorf("missing key-value pair: %d -> %s", k, v)
		}
	}
}
