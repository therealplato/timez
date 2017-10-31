package main

import (
	"fmt"
	"strings"
	"time"
)

func timez(cfg config, c clocker, args []string) string {
	z0 := cfg.localTZ
	outputTZs, t0, err := parse(cfg, args)
	if err != nil {
		if err == errNoArgs {
			outputTZs = append(outputTZs, z0, time.UTC)
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
		outputTZs = append(outputTZs)
	}
	for _, tz := range outputTZs {
		s := t0.In(tz).Format("2006-01-02 15:04:05")
		output += fmt.Sprintf("%s: %s\n", tz.String(), s)
	}
	output = strings.TrimRight(output, "\n")
	return output
}
