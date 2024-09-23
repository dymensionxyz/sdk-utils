package uptr

// To returns a pointer to the passed argument
func To[T any](x T) *T {
	return &x
}

// From returns the value of the passed pointer
func From[T any](x *T) T {
	return *x
}
