package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

var ErrParse = errors.New("could not parse inputs")

func main() {
	fmt.Println(
		timez(os.Args),
	)
}

func timez(args []string) string {
	from, time, to, err := parse(args)
	if err != nil {
		return Usage
	}
	_, _, _ = from, time, to
	return "unimplemented"
}

func parse(args []string) (from *time.Location, t time.Time, to *time.Location, err error) {
	return nil, time.Time{}, nil, ErrParse
}
