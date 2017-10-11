package main

import (
	"time"

	"github.com/stretchr/testify/mock"
)

type clocker interface {
	Now() time.Time
}

type clock struct{}

func (*clock) Now() time.Time {
	return time.Now()
}

type mockClock struct {
	mock.Mock
}

func (c *mockClock) Now() time.Time {
	args := c.Called()
	t, ok := args.Get(0).(time.Time)
	if !ok {
		return time.Time{}
	}
	return t
}
