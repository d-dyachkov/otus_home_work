package hw02_unpack_string //nolint:golint,stylecheck

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type test struct {
	message  string
	input    string
	expected string
	err      error
}

func TestUnpack(t *testing.T) {
	for _, tst := range [...]test{
		{
			message:  "Basic packed string",
			input:    "a4bc2d5e",
			expected: "aaaabccddddde",
		},
		{
			message:  "Only letters",
			input:    "abccd",
			expected: "abccd",
		},
		{
			message:  "The length of the packed character at the beginning of the line",
			input:    "3abc",
			expected: "",
			err:      ErrInvalidString,
		},
		{
			message:  "Only digits",
			input:    "45",
			expected: "",
			err:      ErrInvalidString,
		},
		{
			message:  "Two digits in a row",
			input:    "aaa10b",
			expected: "",
			err:      ErrInvalidString,
		},
		{
			message:  "Empty string",
			input:    "",
			expected: "",
		},
		{
			message:  "Packed character with length 0",
			input:    "aaa0b",
			expected: "aab",
		},
		{
			message:  "String with packed newline character",
			input:    "d\n5abc",
			expected: "d\n\n\n\n\nabc",
		},
		{
			message:  "Russian packed text",
			input:    "приве0т",
			expected: "привт",
		},
		{
			message:  "The length of the packed character in the end string",
			input:    "привет2",
			expected: "приветт",
		},
		{
			message:  "The negative length of the packed character",
			input:    "привет-1",
			expected: "",
			err:      ErrInvalidString,
		},
		{
			message:  "Just a space",
			input:    " ",
			expected: " ",
		},
		{
			message:  "Different register",
			input:    "A1a0a1",
			expected: "Aa",
		},
	} {
		tst := tst
		t.Run(tst.message, func(t *testing.T) {
			result, err := Unpack(tst.input)
			require.Equal(t, tst.err, err)
			require.Equal(t, tst.expected, result)
		})
	}
}

func TestUnpackWithEscape(t *testing.T) {
	t.Skip() // NeedRemove if task with asterisk completed

	for _, tst := range [...]test{
		{
			input:    `qwe\4\5`,
			expected: `qwe45`,
		},
		{
			input:    `qwe\45`,
			expected: `qwe44444`,
		},
		{
			input:    `qwe\\5`,
			expected: `qwe\\\\\`,
		},
		{
			input:    `qwe\\\3`,
			expected: `qwe\3`,
		},
	} {
		result, err := Unpack(tst.input)
		require.Equal(t, tst.err, err)
		require.Equal(t, tst.expected, result)
	}
}
