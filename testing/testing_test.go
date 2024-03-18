package testing_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	hacktesting "source.toby3d.me/toby3d/hacks/testing"
)

func TestGoldenEqual(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest(http.MethodGet, "https://example.com/", nil)
	w := httptest.NewRecorder()
	testHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Testing</title>
  </head>
  <body>
    <h1>Hello, World!</h1>
    <p>This is a testing HTML page of %s website.</p>
  </body>
</html>`, r.Host) // NOTE(toby3d): Host must be 'example.com'
	}

	testHandler(w, req)

	// NOTE(toby3d): compare recorded response body against saved golden file
	hacktesting.GoldenEqual(t, w.Result().Body)
}
