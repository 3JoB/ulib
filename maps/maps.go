package maps

func New[E comparable, T any](map[E]T) map[E]T {
	return make(map[E]T)
}

func Copy[E comparable, T any](m map[E]T) map[E]T {
	ms := make(map[E]T, len(m))
	for k, v := range m {
		ms[k] = v
	}
	return ms
}

func Clean[E comparable, T any](m map[E]T) {
	if m == nil {
		return
	}
	for k := range m {
		delete(m, k)
	}
}
