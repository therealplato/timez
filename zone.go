package main

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/stretchr/testify/mock"
)

type zoner interface {
	Zone() *time.Location
}

type zone struct{}

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
	return nil, errors.New("unimplemented")
}

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
