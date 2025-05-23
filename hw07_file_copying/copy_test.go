package main

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	inputFile := "testdata/input.txt"
	outputFile := "testdata/output.txt"

	tests := []struct {
		offset       int64
		limit        int64
		expectedFile string
	}{
		{offset: 0, limit: 0, expectedFile: "testdata/out_offset0_limit0.txt"},
		{offset: 0, limit: 10, expectedFile: "testdata/out_offset0_limit10.txt"},
		{offset: 0, limit: 1000, expectedFile: "testdata/out_offset0_limit1000.txt"},
		{offset: 0, limit: 10000, expectedFile: "testdata/out_offset0_limit10000.txt"},
		{offset: 100, limit: 1000, expectedFile: "testdata/out_offset100_limit1000.txt"},
		{offset: 100, limit: 1000, expectedFile: "testdata/out_offset100_limit1000.txt"},
		{offset: 6000, limit: 1000, expectedFile: "testdata/out_offset6000_limit1000.txt"},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.expectedFile, func(t *testing.T) {
			err := Copy(inputFile, outputFile, tc.offset, tc.limit)
			expectedFile, _ := os.ReadFile(tc.expectedFile)
			resultFile, _ := os.ReadFile(outputFile)

			if string(expectedFile) != string(resultFile) {
				t.Errorf("expected: %s, got: %s", expectedFile, resultFile)
			}
			require.NoError(t, err)
		})
		_ = os.Remove(outputFile)
	}
}

func TestInvalidCopy(t *testing.T) {
	inputFile := "testdata/input.txt"
	outputFile := "testdata/output.txt"

	tests := []struct {
		inputFile   string
		offset      int64
		expectedErr error
	}{
		{inputFile: "", offset: 0, expectedErr: ErrUnsupportedFile},
		{inputFile: inputFile, offset: 50000, expectedErr: ErrOffsetExceedsFileSize},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.inputFile, func(t *testing.T) {
			err := Copy(tc.inputFile, outputFile, tc.offset, 0)

			require.Truef(t, errors.Is(err, tc.expectedErr), "actual error %q", err)
		})
	}
}
