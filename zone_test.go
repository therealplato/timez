package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestZoneRespectsDefaultAlias(t *testing.T) {
	expected, err := time.LoadLocation("Asia/Dubai")
	require.Nil(t, err)
	z := zone{
		cfg: config{
			aliases: map[string]string{"default": "Asia/Dubai"},
		},
	}
	actual := z.Zone()
	assert.Equal(t, expected, actual)
}
