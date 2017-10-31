package main

import "time"

var defaultAlias = map[string]string{
	"ET":       "US/Eastern",
	"et":       "US/Eastern",
	"eastern":  "US/Eastern",
	"CT":       "US/Central",
	"ct":       "US/Central",
	"central":  "US/Central",
	"MT":       "US/Mountain",
	"mt":       "US/Mountain",
	"mountain": "US/Mountain",
	"PT":       "US/Pacific",
	"pt":       "US/Pacific",
	"pacific":  "US/Pacific",
	"Auckland": "Pacific/Auckland",
	"auckland": "Pacific/Auckland",
	"London":   "Europe/London",
	"london":   "Europe/London",
}

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
