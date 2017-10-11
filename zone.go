package main

import (
	"time"

	"github.com/stretchr/testify/mock"
)

type zoner interface {
	Zone() *time.Location
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
