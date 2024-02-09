// The pointer package contains utilities for working with pointers.
package pointer

// Of returns pointer to provided v.
func Of[T any](v T) *T {
	return &v
}
