package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	zoneRX              = regexp.MustCompile(`^[a-zA-Z+\-]`)
	fragRX              = regexp.MustCompile("^[0-9]")
	UnixLiteralRegex    = regexp.MustCompile("(?i)unix")
	UnixIntegerRegex    = regexp.MustCompile("^[0-9]{9,20}$")
	UnixSRegex          = regexp.MustCompile("^[0-9]{9,11}$")
	UnixMSRegex         = regexp.MustCompile("^[0-9]{12,14}$")
	UnixNSRegex         = regexp.MustCompile("^[0-9]{18,20}$")
	errParsingTimestamp = errors.New("found numbers that did not match a known timestamp format")
	errNoLocalTimezone  = errors.New("parse called without configuring a local timezone")
)

type outputZone struct {
	alias  string
	loc    *time.Location
	isUnix bool
}

func parse(cfg config, args []string) (outputZones []outputZone, t time.Time, matchedTimestampFormat string, err error) {
	var (
		inputZone     *time.Location
		tmpZone       *time.Location
		i             int
		isUnixInteger bool
	)
	t = nullTime
	inputZone = cfg.localTZ
	if inputZone == nil {
		return outputZones, t, "", errNoLocalTimezone
	}
	if len(args) == 0 {
		return outputZones, t, "", errNoArgs
	}

	outputZoneStrings, timeFrags, inputZoneString := parseToStrings(args)

	// fmt.Printf("out: %s\nin: %s\ntime: %s\n", outputZoneStrings, inputZoneString, timeFrags)

	isUnixInteger = len(timeFrags) > 0 && UnixIntegerRegex.MatchString(timeFrags[0])
	if isUnixInteger {
		i, err = strconv.Atoi(timeFrags[0])
		if err != nil {
			return outputZones, t, "", err
		}

		if len(inputZoneString) > 0 {
			fmt.Fprintln(os.Stderr, "ignoring input zone; Unix time is UTC")
		}
		outputZoneStrings = ensureUTC(outputZoneStrings)
		if UnixSRegex.MatchString(timeFrags[0]) {
			t = time.Unix(int64(i), 0)
		}
		if UnixMSRegex.MatchString(timeFrags[0]) {
			var (
				s, ms, ns int64
			)
			// i is milliseconds
			// s is seconds
			// ms is truncated ms of input
			// ns is nanoseconds
			s = int64(i / 1000) // 1111/10 truncates to 111
			// ms = int64(i) - s
			ms = int64(i) / s
			ns = 1E6 * (ms)

			// fmt.Printf(" i: %v\n s: %v\nms: %v\nns: %v\n", i, s, ms, ns)
			t = time.Unix(s, ns)
		}
		if UnixNSRegex.MatchString(timeFrags[0]) {
			t = time.Unix(0, int64(i))
		}
		timeFrags = nil
	}

	if inputZoneString != "" {
		inputZone, err = parseZone(cfg, inputZoneString)
		if err != nil {
			return outputZones, t, "", err
		}
	}

	for _, tzString := range outputZoneStrings {
		tmpZone, err = parseZone(cfg, tzString)
		if err != nil {
			if err == errIsUnixZone {
				outputZones = append(outputZones, outputZone{
					alias:  tzString,
					loc:    nil,
					isUnix: true,
				})
				t = time.Now()
				err = nil
				continue
			}
			return outputZones, t, "", err
		}
		outputZones = append(outputZones, outputZone{
			alias: tzString,
			loc:   tmpZone,
		})
	}

	if len(outputZones) == 0 && inputZone == time.UTC {
		var p, q string
		for _, frag := range timeFrags {
			q += frag
			q += " "
		}
		q = strings.TrimSpace(q)
		p = fmt.Sprintf("converting UTC -> UTC, did you mean to convert local to UTC? `timez UTC %s US/Eastern`\n", q)
		fmt.Fprintf(os.Stderr, p)
	}

	if len(timeFrags) > 0 {
		timeString := strings.Join(timeFrags, " ")
		for _, matchedTimestampFormat = range tsFormats {
			t, err = time.ParseInLocation(matchedTimestampFormat, timeString, inputZone)
			if err == nil {
				break
			}
		}
	}

	return outputZones, t, matchedTimestampFormat, err
}

// take CLI inputs and split them up into output zones, space separated bits of time, and one input zone
func parseToStrings(args []string) (outputZoneStrings []string, timeFrags []string, inputZoneString string) {
	for _, arg := range args {
		if arg == "at" || arg == "in" || arg == "from" || arg == "to" {
			continue
		}

		if UnixIntegerRegex.MatchString(arg) {
			timeFrags = append(timeFrags, arg)
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
		}
	}
	return outputZoneStrings, timeFrags, inputZoneString
}

var errIsUnixZone = errors.New("'Unix' time zone is UTC")

func parseZone(cfg config, s string) (*time.Location, error) {
	tmp, err := time.LoadLocation(s)
	if UnixLiteralRegex.MatchString(s) {
		return nil, errIsUnixZone
	}
	if err != nil {
		// retry with alias:
		alias, ok := cfg.aliases[s]
		if !ok {
			return nil, fmt.Errorf("could not parse %s as timezone", s)
		}
		tmp, err = time.LoadLocation(alias)
	}
	return tmp, err
}

func ensureUTC(ss []string) []string {
	for _, s := range ss {
		if strings.ToLower(s) == "utc" {
			return ss
		}
	}
	ss = append(ss, "UTC")
	return ss
}
