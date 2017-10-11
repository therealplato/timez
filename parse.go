package main

import (
	"fmt"
	"regexp"
	"time"
)

func parse(args []string) (outputTZ []*time.Location, t time.Time, inputTZ *time.Location, err error) {
	if len(args) == 0 {
		return nil, nullTime, nil, ErrNoArgs
	}

	var (
		// outputZoneStrings are timezone strings positioned before any timestamps
		outputZoneStrings []string
		outputZones       []*time.Location
		// frags are number-containing strings that should be fragments of a timestamp
		frags []string
		//inputZones are timezone strings positioned after any timestamps
		inputZones []string
		zoneRX     = regexp.MustCompile(`^[a-zA-Z+\-]`)
		fragRX     = regexp.MustCompile("^[0-9]")
		tz         *time.Location
		errTZParse error
	)

	for _, arg := range args {
		isFrag := fragRX.MatchString(arg)
		if isFrag {
			frags = append(frags, arg)
		}
		isZone := zoneRX.MatchString(arg)
		if isZone {
			if len(frags) == 0 {
				outputZoneStrings = append(outputZoneStrings, arg)
				continue
			}
			inputZones = append(inputZones, arg)
		}
	}
	/*
		fmt.Println("Parse results:")
		fmt.Println(outputZoneStrings)
		fmt.Println(frags)
		fmt.Println(inputZones)
	*/
	for _, tzString := range outputZoneStrings {
		tz, errTZParse = time.LoadLocation(tzString)
		if errTZParse != nil {
			// retry with alias:
			alias, ok := tzAlias[tzString]
			if !ok {
				return nil, nullTime, nil, fmt.Errorf("could not parse %s as timezone", tzString)
			}
			tz, errTZParse = time.LoadLocation(alias)
			if errTZParse != nil {
				return nil, nullTime, nil, fmt.Errorf("could not parse %s as timezone", tzString)
			}
		}
		// successfully parsed tzString into tz
		outputZones = append(outputZones, tz)
	}

	if len(frags) == 0 {
		t = nullTime
	}

	return outputZones, t, nil, nil
}
