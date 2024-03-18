package path_test

import (
	"fmt"

	"source.toby3d.me/toby3d/hacks/path"
)

func Example() {
	head, tail := path.Shift("/foo/bar/index.html")
	fmt.Println(head, tail)
	// Output: foo /bar/index.html
}
