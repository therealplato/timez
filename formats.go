package main

import "time"

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
