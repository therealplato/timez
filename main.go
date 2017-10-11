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

// ErrNoArgs is just for switching
var ErrNoArgs = errors.New("semaphore for default handling")

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
	outputTZs, t0, inputTZ, err := parse(args)
	_ = inputTZ
	if err != nil {
		if err == ErrNoArgs {
			outputTZs = append(outputTZs, z.Zone(), time.UTC)
			t0 = c.Now()
			inputTZ = time.UTC
		} else {
			return Usage
		}
	}
	output := ""
	for _, tz := range outputTZs {
		s := t0.In(tz).Format("2006-01-02 15:04:05")
		output += fmt.Sprintf("%s: %s\n", tz.String(), s)
	}
	return output

}

func parse(args []string) (outputTZ []*time.Location, t time.Time, inputTZ *time.Location, err error) {
	if len(args) == 0 {
		return nil, time.Time{}, nil, ErrNoArgs
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
