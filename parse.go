package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

var tsFormats = []string{
	"2006-01-02 15:04:05",
	"2006-01-02 15:04:05 -0700",
	time.RFC1123,
	time.RFC1123Z,
	time.RFC3339,
	time.RFC3339Nano,
	time.RFC822,
	time.RFC822Z,
	time.RFC850,
	time.ANSIC,
}

var errParsingTimestamp = errors.New("found numbers that did not match a known timestamp format")

func parse(z zoner, args []string) (outputZones []*time.Location, t time.Time, inputZone *time.Location, err error) {
	if len(args) == 0 {
		return nil, nullTime, nil, ErrNoArgs
	}

	var (
		zoneRX = regexp.MustCompile(`^[a-zA-Z+\-]`)
		fragRX = regexp.MustCompile("^[0-9]")
		// outputZoneStrings are timezone strings positioned before any timestamps
		outputZoneStrings []string
		// outputZones       []*time.Location
		// frags are number-containing strings that should be fragments of a timestamp
		frags []string
		//inputZoneString are timezone strings positioned after any timestamps
		inputZoneString string
		// inputZone       *time.Location
		tmpZone    *time.Location
		errTZParse error
		errTSParse error
	)

	for _, arg := range args {
		if arg == "at" || arg == "in" || arg == "from" || arg == "to" {
			continue
		}
		isFrag := fragRX.MatchString(arg)
		if isFrag {
			frags = append(frags, arg)
			continue
		}
		isZone := zoneRX.MatchString(arg)
		if isZone {
			if len(frags) == 0 {
				outputZoneStrings = append(outputZoneStrings, arg)
				continue
			}
			inputZoneString = arg
			alias, ok := tzAlias[inputZoneString]
			if ok {
				inputZoneString = alias
			}
		}
	}
	/*
		fmt.Println("Parse results:")
		fmt.Println(outputZoneStrings)
		fmt.Println(frags)
		fmt.Println(inputZone)
	*/
	if inputZoneString != "" {
		inputZone, errTZParse = time.LoadLocation(inputZoneString)
		if errTZParse != nil {
			err = errTZParse
			return outputZones, t, inputZone, err
		}
	}
	for _, tzString := range outputZoneStrings {
		tmpZone, errTZParse = time.LoadLocation(tzString)
		if errTZParse != nil {
			// retry with alias:
			alias, ok := tzAlias[tzString]
			if !ok {
				return nil, nullTime, nil, fmt.Errorf("could not parse %s as timezone", tzString)
			}
			tmpZone, errTZParse = time.LoadLocation(alias)
			if errTZParse != nil {
				return nil, nullTime, nil, fmt.Errorf("could not parse %s as timezone", tzString)
			}
		}
		// successfully parsed tzString into tz
		outputZones = append(outputZones, tmpZone)
	}

	t = nullTime
	if len(frags) > 0 {
		timeString := strings.Join(frags, " ")
		for i, f := range tsFormats {
			if inputZone != nil {
				t, errTSParse = time.ParseInLocation(f, timeString, inputZone)
			} else {
				t, errTSParse = time.ParseInLocation(f, timeString, z.Zone())
			}
			if errTSParse == nil {
				break
			}
			if i == len(tsFormats)-1 {
				err = errParsingTimestamp
			}
		}
	}

	return outputZones, t, nil, err
}
