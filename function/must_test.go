package function_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"source.toby3d.me/toby3d/hacks/function"
)

func TestMust(t *testing.T) {
	t.Parallel()

	testFunc := func(input string) (result [2]float64, err error) {
		coords := strings.Split(input, ",")
		if len(coords) < 2 {
			return result, fmt.Errorf("got '%s', want input in format '9.999,9.999'", input)
		}

		for i := 0; i < 2; i++ {
			if result[i], err = strconv.ParseFloat(coords[i], 64); err != nil {
				return result, fmt.Errorf("cannot parse coordinate '%s': %w", coords[i], err)
			}
		}

		return result, nil
	}

	if actual := function.Must(testFunc("1.23,9.87")); len(actual) != 2 {
		t.Errorf("want 2 lenght float array, got %v", actual)
	}
}
