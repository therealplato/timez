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
		testcase{"junk input", "asdf", Usage},
		testcase{"empty input outputs local and utc", "", `Pacific/Auckland: 2017-10-10 23:01:30
UTC: 2017-10-10 10:01:30
`,
		},
		testcase{"one tz outputs now in that tz", "PT", `US/Pacific: 2017-10-10 03:01:30
`,
		},
		testcase{"three tz outputs now in those tz", "PT Pacific/Auckland UTC", `US/Pacific: 2017-10-10 03:01:30
Pacific/Auckland: 2017-10-10 23:01:30
UTC: 2017-10-10 10:01:30
`,
		},
		testcase{"one tz and a timestamp without zone, outputs  that local time converted to that zone", "PT 2017-10-10 23:45:00", `US/Pacific: 2017-10-10 03:45:00
`,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			args := strings.Split(tc.input, " ")
			if len(args) == 1 && args[0] == "" {
				args = nil
			}
			out := timez(c, z, args)
			assert.Equal(t, tc.expected, out)
		})
	}
}
