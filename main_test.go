package main

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testcase struct {
	name     string
	input    string
	expected string
}

func TestTimez(t *testing.T) {
	t0, err := time.Parse("2006-01-02 15:04:05-0700", "2017-10-10 23:01:30+1300")
	require.Nil(t, err)
	z0, err := time.LoadLocation("Pacific/Auckland")
	require.Nil(t, err)

	z := &mockZone{}
	z.On("Zone").Return(z0)

	c := &mockClock{}
	c.On("Now").Return(t0)

	tcs := []testcase{
		testcase{
			name:     "junk input",
			input:    "asdf",
			expected: usage,
		},
		testcase{
			name:  "empty input outputs local and utc",
			input: "",
			expected: `Pacific/Auckland: 2017-10-10 23:01:30
UTC: 2017-10-10 10:01:30`},
		testcase{
			name:     "one tz outputs now in that tz",
			input:    "PT",
			expected: `US/Pacific: 2017-10-10 03:01:30`,
		},
		testcase{
			name:  "three tz outputs now in those tz",
			input: "PT Pacific/Auckland UTC",
			expected: `US/Pacific: 2017-10-10 03:01:30
Pacific/Auckland: 2017-10-10 23:01:30
UTC: 2017-10-10 10:01:30`},
		testcase{
			name:     "one tz and a timestamp without zone, outputs that local time converted to that zone",
			input:    "PT 2017-10-10 23:45:00",
			expected: `US/Pacific: 2017-10-10 03:45:00`,
		},
		testcase{
			name:  "multiple tz and a timestamp without zone, outputs local time converted to those zones",
			input: "PT ET 2017-10-10 23:45:00",
			expected: `US/Pacific: 2017-10-10 03:45:00
US/Eastern: 2017-10-10 06:45:00`},
		testcase{
			name:     "one tz and a timestamp and zone, converts second zone to first",
			input:    "PT 2017-10-10 23:45:00 UTC",
			expected: `US/Pacific: 2017-10-10 16:45:00`,
		},
		testcase{
			name:     "`at in from to` are discarded",
			input:    "in to PT at 2017-10-10 23:45:00 from UTC",
			expected: `US/Pacific: 2017-10-10 16:45:00`,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			args := strings.Split(tc.input, " ")
			if len(args) == 1 && args[0] == "" {
				args = nil
			}
			cfg := config{
				localTZ: z0,
			}
			out := timez(cfg, c, args)
			assert.Equal(t, tc.expected, out)
		})
	}
}
