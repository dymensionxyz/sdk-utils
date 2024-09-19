package uslices

// ToKeySet converts a slice of elements into a set (represented as a map with empty structs).
// This function ensures that all elements in the resulting map are unique.
func ToKeySet[E comparable](slice []E) map[E]struct{} {
	res := make(map[E]struct{}, len(slice))
	for _, e := range slice {
		res[e] = struct{}{}
	}
	return res
}

// Merge takes multiple slices and merges them into a single slice. The resulting slice contains all
// elements from the input slices in the order they are passed.
func Merge[S ~[]V, V any](s ...S) S {
	var l int
	for _, slice := range s {
		l += len(slice)
	}
	r := make(S, 0, l)
	for _, slice := range s {
		r = append(r, slice...)
	}
	return r
}

// Map applies a provided function f to each element in the input slice and returns a slice of modified values.
func Map[E, V any](slice []E, f func(E) V) []V {
	res := make([]V, 0, len(slice))
	for _, e := range slice {
		res = append(res, f(e))
	}
	return res
}
