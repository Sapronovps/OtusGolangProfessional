package hw02unpackstring

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		{input: "🙃0", expected: ""},
		{input: "aaф0b", expected: "aab"},
		// uncomment if task with asterisk completed
		{input: `qwe\4\5`, expected: `qwe45`},
		{input: `qwe\45`, expected: `qwe44444`},
		{input: `qwe\\5`, expected: `qwe\\\\\`},
		{input: `qwe\\\3`, expected: `qwe\3`},
		{input: "а1б2в3", expected: "аббввв"},
		{input: "a0", expected: ""},
		{input: "a1", expected: "a"},
		{input: "я9", expected: "яяяяяяяяя"},
		{input: "สวัสดี", expected: "สวัสดี"},
		{input: "สวัส4ดี", expected: "สวัสสสสดี"},
		{input: "🙂9", expected: "🙂🙂🙂🙂🙂🙂🙂🙂🙂"},

		{input: "১২৩", expected: "১২৩"},
		{input: "১2২৩0", expected: "১১২"},
		{input: "੩4", expected: "੩੩੩੩"},
		{input: `искус2тво`, expected: "искусство"},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{"3abc", "45", "aaa10b"}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}

func TestHasBackslash(t *testing.T) {
	validStrings := []string{`qwe\4\5`, `qwe\45`, `qwe\\\3`}
	for _, tc := range validStrings {
		t.Run(tc, func(t *testing.T) {
			hasBackslash := hasBackslash(tc)
			require.True(t, hasBackslash, fmt.Sprintf("func hasBackslash for string %s must be return true", tc))
		})
	}
}
