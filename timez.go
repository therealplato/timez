package main

import (
	"fmt"
	"strings"
	"time"
)

func timez(cfg config, c clocker, args []string) string {
	z0 := cfg.localTZ
	outputZones, t0, err := parse(cfg, args)
	if err != nil {
		if err == errNoArgs {
			outputZones = append(outputZones, outputZone{
				alias: z0.String(),
				loc:   z0,
			}, outputZone{
				alias: "UTC",
				loc:   time.UTC,
			})
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
	for _, outputZone := range outputZones {
		s := t0.In(outputZone.loc).Format("2006-01-02 15:04:05")
		output += fmt.Sprintf("%s: %s\n", outputZone.alias, s)
	}
	output = strings.TrimRight(output, "\n")
	return output
}
