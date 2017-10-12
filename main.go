package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

var errParse = errors.New("could not parse inputs")

var errNoArgs = errors.New("semaphore for default handling")

var nullTime = time.Time{}

func main() {
	c := &clock{}
	z := &zone{}
	args := make([]string, len(os.Args)-1)
	args = os.Args[1:]
	fmt.Println(timez(c, z, args))
}

func timez(c clocker, z zoner, args []string) string {
	outputTZs, t0, err := parse(z, args)
	if err != nil {
		if err == errNoArgs {
			outputTZs = append(outputTZs, z.Zone(), time.UTC)
			t0 = c.Now()
		} else {
			fmt.Println(err)
			return usage
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
	output = strings.TrimRight(output, "\n")
	return output
}
