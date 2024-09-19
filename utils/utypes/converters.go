package utypes

// StringToType converts a string to a custom type T, where T can be either a string or a byte slice.
// May be used as a parameter for uslices.Map.
func StringToType[T ~string | ~[]byte](s string) T { return T(s) }

// TypeToString converts a custom type T, where T can be either a string or a byte slice, back to a string.
// May be used as a parameter for uslices.Map.
func TypeToString[T ~string | ~[]byte](t T) string { return string(t) }
