package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"time"
)

// ErrParse indicates the cli arguments were bad
var ErrParse = errors.New("could not parse inputs")

var nullTime = time.Time{}

func main() {
	c := &clock{}
	z := &zone{}
	args := make([]string, len(os.Args)-1)
	args = os.Args[1:]
	fmt.Println(
		timez(c, z, args),
	)
}

func timez(c clocker, z zoner, args []string) string {
	from, t0, to, err := parse(c, z, args)
	if err != nil {
		return Usage
	}
	t1s := t0.In(from).Format("2006-01-02 15:04:05")
	t2s := t0.In(to).Format("2006-01-02 15:04:05")

	return fmt.Sprintf("%s: %s\n%s: %s\n", from.String(), t1s, to.String(), t2s)
}

func parse(c clocker, z zoner, args []string) (from *time.Location, t time.Time, to *time.Location, err error) {
	if len(args) == 0 {
		t0 := c.Now()
		z0 := z.Zone()
		z1 := time.UTC
		return z0, t0, z1, nil
	}

	var (
		// leftZones are timezone strings positioned before any timestamps
		leftZones []string
		// frags are number-containing strings that should be fragments of a timestamp
		frags []string
		//rightZones are timezone strings positioned after any timestamps
		rightZones []string
		zoneRX     = regexp.MustCompile(`^[a-zA-Z+\-]`)
		fragRX     = regexp.MustCompile("^[0-9]")
	)

	for _, arg := range args {
		isFrag := fragRX.MatchString(arg)
		if isFrag {
			frags = append(frags, arg)
		}
		isZone := zoneRX.MatchString(arg)
		if isZone {
			if len(frags) == 0 {
				leftZones = append(leftZones, arg)
				continue
			}
			rightZones = append(rightZones, arg)
		}
	}
	fmt.Println("Parse results:")
	fmt.Println(leftZones)
	fmt.Println(frags)
	fmt.Println(rightZones)

	return nil, time.Time{}, nil, ErrParse
}
