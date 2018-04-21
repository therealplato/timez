package main

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testcase struct {
	name     string
	args     string
	cfg      *config
	expected string
}

func TestTimez(t *testing.T) {

	// 1522741510

	t0, err := time.Parse("2006-01-02 15:04:05-0700", "2017-10-10 23:01:30+1300")
	require.Nil(t, err)
	z0, err := time.LoadLocation("Pacific/Auckland")
	require.Nil(t, err)
	z1, err := time.LoadLocation("Asia/Dubai")
	require.Nil(t, err)

	z := &mockZone{}
	z.On("Zone").Return(z0)

	c := &mockClock{}
	c.On("Now").Return(t0)

	tcs := []testcase{
		testcase{
			name:     "junk input",
			args:     "asdf",
			expected: usage,
		},
		testcase{
			name:     "empty input outputs local and utc",
			args:     "",
			expected: `UTC: 2017-10-10 10:01:30`,
		},
		// Old behavior: Specifying an input timezone and no output timezone, outputs local
		// testcase{
		// 	name:     "no output zone assumes local",
		// 	args:     "2017-10-10 10:59:00 UTC",
		// 	expected: `Pacific/Auckland: 2017-10-10 23:59:00`,
		// },
		testcase{
			name:     "no output zone assumes local",
			args:     "2017-10-10 10:59:00 UTC",
			expected: `UTC: 2017-10-10 10:59:00`,
		},
		testcase{
			name:     "one tz outputs now in that tz",
			args:     "US/Pacific",
			expected: `US/Pacific: 2017-10-10 03:01:30`,
		},
		testcase{
			name: "three tz outputs now in those tz",
			args: "US/Pacific Pacific/Auckland UTC",
			expected: `US/Pacific: 2017-10-10 03:01:30
Pacific/Auckland: 2017-10-10 23:01:30
UTC: 2017-10-10 10:01:30`},
		testcase{
			name:     "one tz and a timestamp without zone, outputs that local time converted to that zone",
			args:     "US/Pacific 2017-10-10 23:45:00",
			expected: `US/Pacific: 2017-10-10 03:45:00`,
		},
		testcase{
			name: "multiple tz and a timestamp without zone, outputs local time converted to those zones",
			args: "US/Pacific US/Eastern 2017-10-10 23:45:00",
			expected: `US/Pacific: 2017-10-10 03:45:00
US/Eastern: 2017-10-10 06:45:00`},
		testcase{
			name:     "one tz and a timestamp and zone, converts second zone to first",
			args:     "US/Pacific 2017-10-10 23:45:00 UTC",
			expected: `US/Pacific: 2017-10-10 16:45:00`,
		},
		testcase{
			name:     "`at in from to` are discarded",
			args:     "in to US/Pacific at 2017-10-10 23:45:00 from UTC",
			expected: `US/Pacific: 2017-10-10 16:45:00`,
		},
		// 		testcase{
		// 			name: "respects configured `default` alias",
		// 			args: "",
		// 			cfg:  &config{localTZ: z1},
		// 			expected: `Asia/Dubai: 2017-10-10 14:01:30
		// UTC: 2017-10-10 10:01:30`,
		// 		},
		testcase{
			name:     "respects configured `custom` alias",
			args:     "custom",
			cfg:      &config{localTZ: z1, aliases: map[string]string{"custom": "Europe/London"}},
			expected: `custom: 2017-10-10 11:01:30`,
		},
		testcase{
			cfg: &config{
				localTZ: z1,
			},
			name:     "transforms Unix to default",
			args:     "1522741510",
			expected: `UTC: 2018-04-03 07:45:10`,
		},
		testcase{
			name:     "transforms Unix to UTC",
			args:     "UTC 1522741510",
			expected: `UTC: 2018-04-03 07:45:10`,
		},
		testcase{
			name:     "one tz outputs now in that tz",
			args:     "US/Pacific",
			expected: `US/Pacific: 2017-10-10 03:01:30`,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			args := strings.Split(tc.args, " ")
			if len(args) == 1 && args[0] == "" {
				args = nil
			}
			if tc.cfg == nil {
				tc.cfg = &config{
					localTZ: z0,
					aliases: defaultAlias,
				}
			}
			out := timez(*tc.cfg, c, args)
			assert.Equal(t, tc.expected, out, fmt.Sprintf("input: %s", tc.args))
		})
	}
}
