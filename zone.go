package main

import (
	"time"

	"github.com/stretchr/testify/mock"
)

type zoner interface {
	Zone() *time.Location
}

type zone struct {
	cfg config
}

func (z *zone) Zone() *time.Location {
	var (
		loc *time.Location
		err error
	)
	d, ok := z.cfg.aliases["default"]
	if ok {
		loc, err := parseZone(z.cfg, d)
		if err == nil {
			return loc
		}
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
