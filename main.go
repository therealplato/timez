package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

// ErrParse indicates the cli arguments were bad
var ErrParse = errors.New("could not parse inputs")

// ErrNoArgs is just for switching
var ErrNoArgs = errors.New("semaphore for default handling")

var nullTime = time.Time{}

func main() {
	c := &clock{}
	z := &zone{}
	args := make([]string, len(os.Args)-1)
	args = os.Args[1:]
	fmt.Println(
		timez(c, z, args),
	)
}

func timez(c clocker, z zoner, args []string) string {
	outputTZs, t0, inputTZ, err := parse(z, args)
	_ = inputTZ
	if err != nil {
		if err == ErrNoArgs {
			outputTZs = append(outputTZs, z.Zone(), time.UTC)
			t0 = c.Now()
			inputTZ = time.UTC
		} else {
			fmt.Println(err)
			return Usage
		}
	}
	if t0 == nullTime {
		t0 = c.Now()
	}
	output := ""
	if len(outputTZs) == 0 {
		outputTZs = append(outputTZs, z.Zone())
	}
	for _, tz := range outputTZs {
		s := t0.In(tz).Format("2006-01-02 15:04:05")
		output += fmt.Sprintf("%s: %s\n", tz.String(), s)
	}
	return output

}
