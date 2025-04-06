package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	tests := []struct {
		nameEnv            string
		expectedValue      interface{}
		expectedNeedRemove bool
	}{
		{nameEnv: "EMPTY", expectedValue: "", expectedNeedRemove: false},
		{nameEnv: "UNSET", expectedValue: "", expectedNeedRemove: true},
		{nameEnv: "HELLO", expectedValue: "\"hello\"", expectedNeedRemove: false},
		{nameEnv: "BAR", expectedValue: "bar", expectedNeedRemove: false},
	}

	mapEnv, err := ReadDir("testdata/env")
	if err != nil {
		fmt.Println(fmt.Errorf("read env dir error: %w", err))
		return
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.nameEnv, func(t *testing.T) {
			require.Equal(t, tc.expectedValue, mapEnv[tc.nameEnv].Value)
			require.Equal(t, tc.expectedNeedRemove, mapEnv[tc.nameEnv].NeedRemove)
		})
	}
}

func TestInvalidDir(t *testing.T) {
	t.Run("invalid dir", func(t *testing.T) {
		_, err := ReadDir("testdata/env22")
		require.Error(t, err)
		require.ErrorContains(t, err, "error reading directory")
	})
}
