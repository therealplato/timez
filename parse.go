package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

var (
	tsFormats = []string{
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
	zoneRX              = regexp.MustCompile(`^[a-zA-Z+\-]`)
	fragRX              = regexp.MustCompile("^[0-9]")
	errParsingTimestamp = errors.New("found numbers that did not match a known timestamp format")
)

func parse(z zoner, args []string) (outputZones []*time.Location, t time.Time, inputZone *time.Location, err error) {
	t = nullTime
	inputZone = z.Zone()
	var (
		tmpZone *time.Location
	)
	if len(args) == 0 {
		return outputZones, t, inputZone, ErrNoArgs
	}

	outputZoneStrings, timeFrags, inputZoneString := parseToStrings(args)

	if inputZoneString != "" {
		inputZone, err = parseZone(inputZoneString)
		if err != nil {
			return outputZones, t, inputZone, err
		}
	}

	for _, tzString := range outputZoneStrings {
		tmpZone, err = parseZone(tzString)
		if err != nil {
			return outputZones, t, inputZone, err
		}
		outputZones = append(outputZones, tmpZone)
	}

	if len(timeFrags) > 0 {
		timeString := strings.Join(timeFrags, " ")
		for _, f := range tsFormats {
			t, err = time.ParseInLocation(f, timeString, inputZone)
			if err == nil {
				break
			}
		}
	}

	return outputZones, t, inputZone, err
}

func parseToStrings(args []string) (outputZoneStrings []string, timeFrags []string, inputZoneString string) {
	for _, arg := range args {
		if arg == "at" || arg == "in" || arg == "from" || arg == "to" {
			continue
		}
		isFrag := fragRX.MatchString(arg)
		if isFrag {
			timeFrags = append(timeFrags, arg)
			continue
		}
		isZone := zoneRX.MatchString(arg)
		if isZone {
			if len(timeFrags) == 0 {
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
	return outputZoneStrings, timeFrags, inputZoneString
}

func parseZone(s string) (*time.Location, error) {
	tmp, err := time.LoadLocation(s)
	if err != nil {
		// retry with alias:
		alias, ok := tzAlias[s]
		if !ok {
			return nil, fmt.Errorf("could not parse %s as timezone", s)
		}
		tmp, err = time.LoadLocation(alias)
	}
	return tmp, err
}
