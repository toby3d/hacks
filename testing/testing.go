// The testing package contains utilities for autotests.
package testing

import (
	"bytes"
	"errors"
	"flag"
	"io"
	"os"
	"path/filepath"
	"testing"
)

//nolint:gochecknoglobals // сompiler global flags cannot be read from within tests
var update = flag.Bool("update", false, "save current tests results as golden files")

// GoldenEqual compares the bytes of the provided output with the contents of
// the golden file for a exact match.
//
// When running go test with the -update flag, the contents of golden-files will
// be overwritten with the provided contents of output, creating the testdata/
// directory if it does not exist.
//
// See: https://youtu.be/8hQG7QlcLBk?t=749
//
//nolint:cyclop // no need for splitting
func GoldenEqual(tb testing.TB, output io.Reader) {
	tb.Helper()

	workDir, err := os.Getwd()
	if err != nil {
		tb.Fatal("cannot get current working directory path:", err)
	}

	actual, err := io.ReadAll(output)
	if err != nil {
		tb.Fatal("cannot read provided data:", err)
	}

	dir := filepath.Join(workDir, "testdata")
	file := filepath.Join(dir, tb.Name()+".golden")

	//nolint:nestif // errchecks for testdata folder first, then for output
	if *update {
		_, err = os.Stat(dir)
		if err != nil && !errors.Is(err, os.ErrExist) && !errors.Is(err, os.ErrNotExist) {
			tb.Fatal("cannot create testdata folder for golden files:", err)
		}

		if errors.Is(err, os.ErrNotExist) {
			if err = os.Mkdir(dir, os.ModePerm); err != nil {
				tb.Fatal("cannot create testdata folder for golden files:", err)
			}
		}

		if err = os.WriteFile(file, actual, os.ModePerm); err != nil {
			tb.Fatal("cannot write data into golden file:", err)
		}

		tb.Skip("skipped due force updating golden file")

		return
	}

	expect, err := os.ReadFile(file)
	if err != nil {
		tb.Fatal("cannot read golden file data:", err)
	}

	if !bytes.Equal(actual, expect) {
		tb.Errorf("the test output does not match the contents of %s.golden file", tb.Name())
	}
}
