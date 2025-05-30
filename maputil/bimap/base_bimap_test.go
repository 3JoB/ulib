package bimap

import (
	"testing"
)

func TestBaseInsert(t *testing.T) {
	tests := []struct {
		name      string
		key       int
		value     string
		immutable bool
		want      bool
	}{
		{name: "insert valid pair", key: 1, value: "one", immutable: false, want: true},
		{name: "insert duplicate key", key: 1, value: "one again", immutable: false, want: true},
		{name: "insert immutable", key: 2, value: "two", immutable: true, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBaseBiMap[int, string]()
			b.immutable = tt.immutable
			got := b.Insert(tt.key, tt.value)
			if got != tt.want {
				t.Errorf("Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaseExists(t *testing.T) {
	tests := []struct {
		name  string
		setup func(b *BaseBiMap[int, string])
		key   int
		want  bool
	}{
		{name: "key exists", setup: func(b *BaseBiMap[int, string]) { b.Insert(1, "one") }, key: 1, want: true},
		{name: "key does not exist", setup: func(b *BaseBiMap[int, string]) {}, key: 2, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBaseBiMap[int, string]()
			tt.setup(b)
			got := b.Exists(tt.key)
			if got != tt.want {
				t.Errorf("Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaseExistsInverse(t *testing.T) {
	tests := []struct {
		name  string
		setup func(b *BaseBiMap[int, string])
		value string
		want  bool
	}{
		{name: "value exists", setup: func(b *BaseBiMap[int, string]) { b.Insert(1, "one") }, value: "one", want: true},
		{name: "value does not exist", setup: func(b *BaseBiMap[int, string]) {}, value: "two", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBaseBiMap[int, string]()
			tt.setup(b)
			got := b.ExistsInverse(tt.value)
			if got != tt.want {
				t.Errorf("ExistsInverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaseGet(t *testing.T) {
	tests := []struct {
		name  string
		setup func(b *BaseBiMap[int, string])
		key   int
		want  string
		found bool
	}{
		{name: "key exists", setup: func(b *BaseBiMap[int, string]) { b.Insert(1, "one") }, key: 1, want: "one", found: true},
		{name: "key does not exist", setup: func(b *BaseBiMap[int, string]) {}, key: 2, want: "", found: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBaseBiMap[int, string]()
			tt.setup(b)
			got, found := b.Get(tt.key)
			if got != tt.want || found != tt.found {
				t.Errorf("Get() = %v, %v, want %v, %v", got, found, tt.want, tt.found)
			}
		})
	}
}

func TestBaseGetInverse(t *testing.T) {
	tests := []struct {
		name  string
		setup func(b *BaseBiMap[int, string])
		value string
		want  int
		found bool
	}{
		{name: "value exists", setup: func(b *BaseBiMap[int, string]) { b.Insert(1, "one") }, value: "one", want: 1, found: true},
		{name: "value does not exist", setup: func(b *BaseBiMap[int, string]) {}, value: "two", want: 0, found: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBaseBiMap[int, string]()
			tt.setup(b)
			got, found := b.GetInverse(tt.value)
			if got != tt.want || found != tt.found {
				t.Errorf("GetInverse() = %v, %v, want %v, %v", got, found, tt.want, tt.found)
			}
		})
	}
}

func TestBaseDelete(t *testing.T) {
	tests := []struct {
		name      string
		setup     func(b *BaseBiMap[int, string])
		key       int
		immutable bool
		want      bool
	}{
		{name: "delete existing key", setup: func(b *BaseBiMap[int, string]) { b.Insert(1, "one") }, key: 1, immutable: false, want: true},
		{name: "delete non-existing key", setup: func(b *BaseBiMap[int, string]) {}, key: 2, immutable: false, want: false},
		{name: "delete on immutable map", setup: func(b *BaseBiMap[int, string]) { b.Insert(1, "one") }, key: 1, immutable: true, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBaseBiMap[int, string]()
			b.immutable = tt.immutable
			tt.setup(b)
			got := b.Delete(tt.key)
			if got != tt.want {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaseDeleteInverse(t *testing.T) {
	tests := []struct {
		name      string
		setup     func(b *BaseBiMap[int, string])
		value     string
		immutable bool
		want      bool
	}{
		{name: "delete existing value", setup: func(b *BaseBiMap[int, string]) { b.Insert(1, "one") }, value: "one", immutable: false, want: true},
		{name: "delete non-existing value", setup: func(b *BaseBiMap[int, string]) {}, value: "two", immutable: false, want: false},
		{name: "delete on immutable map", setup: func(b *BaseBiMap[int, string]) { b.Insert(1, "one") }, value: "one", immutable: true, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBaseBiMap[int, string]()
			b.immutable = tt.immutable
			tt.setup(b)
			got := b.DeleteInverse(tt.value)
			if got != tt.want {
				t.Errorf("DeleteInverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
