package pointer_test

import (
	"source.toby3d.me/toby3d/hacks/pointer"
)

func ExampleOf() {
	// NOTE(toby3d): ResponsePayload can hold data in three states:
	//
	//   * nil value
	//   * zero value: "", 0, false
	//   * any valid value: "abc", 42, true
	//
	// So, for support all these states struct must have pointers to basic
	// types.
	type ResponsePayload struct {
		Text   *string
		Ok     *bool
		Number *int
	}

	// NOTE(toby3d): the old way
	ok := true
	text := "hello, world"
	num := 0
	_ = ResponsePayload{
		Text:   &text,
		Ok:     &ok,
		Number: &num,
	}

	// NOTE(toby3d): hack-way
	_ = ResponsePayload{
		Text:   pointer.Of("hello, world"),
		Ok:     pointer.Of(true),
		Number: pointer.Of(0),
	}
}
