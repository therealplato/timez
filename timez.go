package main

import (
	"fmt"
	"strings"
	"time"
)

func timez(cfg config, c clocker, args []string) string {
	// z0 := cfg.localTZ
	outputZones, t0, tsFmt, err := parse(cfg, args)
	if err != nil {
		if err == errNoArgs {
			t0 = c.Now()
		} else {
			fmt.Println(err)
			return usage
		}
	}
	if t0 == nullTime {
		t0 = c.Now()
	}
	if tsFmt == "" {
		tsFmt = "2006-01-02 15:04:05"
	}
	output := ""
	if len(outputZones) == 0 {
		// outputZones = append(outputZones, outputZone{
		// 	alias: z0.String(),
		// 	loc:   z0,
		// })
		outputZones = append(outputZones, outputZone{
			alias: "UTC",
			loc:   time.UTC,
		})
	}
	for _, outputZone := range outputZones {
		s := t0.In(outputZone.loc).Format(tsFmt)
		output += fmt.Sprintf("%s: %s\n", outputZone.alias, s)
	}
	output = strings.TrimRight(output, "\n")
	return output
}
