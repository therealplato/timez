package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTestParseToStringsFindsUnix(t *testing.T) {
	u := "1522741510"
	zOut, times, zIn := parseToStrings([]string{u})
	_, _ = times, zOut
	assert.Equal(t, "", zIn)
	assert.Nil(t, zOut)
	assert.Equal(t, []string{"1522741510"}, times)
}
