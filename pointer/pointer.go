package pointer

// Of returns pointer to provided v.
func Of[T any](v T) *T {
	return &v
}
