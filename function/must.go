package function

// Must initiates a panic if the function provided in it returns an error.
//
// See [regexp.MustCompile] as an example of how to use it.
//
// WARN(toby3d): this utility should only be used for data provided by the
// developer side. Any data from other sources should be handled in the regular
// way.
func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}

	return v
}
