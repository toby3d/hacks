// The path package contains utilities for working with URL paths.
package path

import (
	"path"
	"strings"
)

// Shift splits off the first component of raw, which will be cleaned of
// relative components before processing. head will never contain a slash and
// tail will always be a rooted path without trailing slash.
//
// See: https://blog.merovius.de/posts/2017-06-18-how-not-to-use-an-http-router/
//
//nolint:nonamedreturns // contains multiple same typed returns
func Shift(raw string) (head, tail string) {
	raw = path.Clean("/" + raw)
	if i := strings.Index(raw[1:], "/") + 1; i <= 0 {
		return raw[1:i], raw[i:]
	}

	return raw[1:], "/"
}
