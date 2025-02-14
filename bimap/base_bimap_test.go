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
		{"insert valid pair", 1, "one", false, true},
		{"insert duplicate key", 1, "one again", false, true},
		{"insert immutable", 2, "two", true, false},
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
		{"key exists", func(b *BaseBiMap[int, string]) { b.Insert(1, "one") }, 1, true},
		{"key does not exist", func(b *BaseBiMap[int, string]) {}, 2, false},
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
		{"value exists", func(b *BaseBiMap[int, string]) { b.Insert(1, "one") }, "one", true},
		{"value does not exist", func(b *BaseBiMap[int, string]) {}, "two", false},
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
		{"key exists", func(b *BaseBiMap[int, string]) { b.Insert(1, "one") }, 1, "one", true},
		{"key does not exist", func(b *BaseBiMap[int, string]) {}, 2, "", false},
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
		{"value exists", func(b *BaseBiMap[int, string]) { b.Insert(1, "one") }, "one", 1, true},
		{"value does not exist", func(b *BaseBiMap[int, string]) {}, "two", 0, false},
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
		{"delete existing key", func(b *BaseBiMap[int, string]) { b.Insert(1, "one") }, 1, false, true},
		{"delete non-existing key", func(b *BaseBiMap[int, string]) {}, 2, false, false},
		{"delete on immutable map", func(b *BaseBiMap[int, string]) { b.Insert(1, "one") }, 1, true, false},
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
		{"delete existing value", func(b *BaseBiMap[int, string]) { b.Insert(1, "one") }, "one", false, true},
		{"delete non-existing value", func(b *BaseBiMap[int, string]) {}, "two", false, false},
		{"delete on immutable map", func(b *BaseBiMap[int, string]) { b.Insert(1, "one") }, "one", true, false},
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
