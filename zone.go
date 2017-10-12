package main

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/stretchr/testify/mock"
)

type zoner interface {
	Zone() *time.Location
}

type zone struct{}

func (z *zone) Zone() *time.Location {
	loc, err := tzFromRC()
	if err == nil {
		return loc
	}
	loc, err = tzFromShell()
	if err == nil {
		return loc
	}
	return time.UTC
}

type mockZone struct {
	mock.Mock
}

func (z *mockZone) Zone() *time.Location {
	args := z.Called()
	z0, ok := args.Get(0).(*time.Location)
	if !ok {
		return time.UTC
	}
	return z0
}

// https://stackoverflow.com/a/41786440
func userHomeDir() string {
	env := "HOME"
	if runtime.GOOS == "windows" {
		env = "USERPROFILE"
	} else if runtime.GOOS == "plan9" {
		env = "home"
	}
	return os.Getenv(env)
}

func tzFromRC() (*time.Location, error) {
	dir := userHomeDir()
	rc := filepath.Join(dir, ".timezrc")
	bb, err := ioutil.ReadFile(rc)
	if err != nil {
		return nil, err
	}
	s := string(bb)
	lines := strings.Split(s, "\n")
	for _, l := range lines {
		loc, err := parseZone(l)
		if err == nil {
			return loc, nil
		}
	}
	return nil, errors.New("~/.timezrc had no parsable timezones")
}

func tzFromShell() (*time.Location, error) {
	bb, err := exec.Command("date", "+%z").Output()
	if err != nil {
		return nil, err
	}
	name := strings.TrimSpace(string(bb))

	dateRX := regexp.MustCompile(`^([+\-])(\d\d)(\d\d)\s?$`)
	match := dateRX.FindSubmatch(bb)
	if len(match) == 0 {
		return nil, errors.New("date +%z output didn't match expected regex")
	}

	signString := string(match[1][0])
	sign := 1
	if signString == "-" {
		sign = -1
	}

	if len(match[2]) != 2 {
		return nil, errors.New(`date +%z output did not have two hour bytes`)
	}
	if len(match[3]) != 2 {
		return nil, errors.New(`date +%z output did not have two minute bytes`)
	}

	hh := string(match[2])
	mm := string(match[3])

	h, err := strconv.Atoi(hh)
	if err != nil {
		return nil, errors.New("second and third bytes of date +%z were not integers")
	}
	m, err := strconv.Atoi(mm)
	if err != nil {
		return nil, errors.New("fourth and fifth bytes of date +%z were not integers")
	}

	s := (60*m + 3600*h) * sign
	return time.FixedZone(name, s), nil
}
