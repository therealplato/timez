package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	name     string
	input    string
	expected string
}

func TestTimez(t *testing.T) {
	tcs := []testcase{
		testcase{"empty input", "", Usage},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			args := strings.Split(tc.input, " ")
			out := timez(args)
			assert.Equal(t, tc.expected, out)
		})
	}
}
