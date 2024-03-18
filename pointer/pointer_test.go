package pointer_test

import (
	"fmt"

	"source.toby3d.me/toby3d/hacks/pointer"
)

func Example() {
	val := struct {
		Text string
	}{Text: "Hello, World!"}
	fmt.Printf("Value: %v\n", val)

	point := pointer.Of(val)
	point.Text = "Hello, Go!"
	fmt.Printf("Pointer: %v\n", point)

	// Output:
	// Value: {Hello, World!}
	// Pointer: &{Hello, Go!}
}
