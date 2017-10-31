package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var errParse = errors.New("could not parse inputs")

var errNoArgs = errors.New("semaphore for default handling")

var nullTime = time.Time{}

func main() {
	var (
		a   map[string]string
		c   = &clock{}
		cfg config
	)
	f, err := os.Open(filepath.Join(userHomeDir(), ".timezrc"))
	if err != nil {
		a = mustLoadAliases(bytes.NewBuffer(nil))
	} else {
		a = mustLoadAliases(f)
	}
	cfg.aliases = a
	z := &zone{
		cfg: cfg,
	}
	cfg.localTZ = z.Zone()
	args := make([]string, len(os.Args)-1)
	args = os.Args[1:]
	fmt.Println(timez(cfg, c, args))
}

func timez(cfg config, c clocker, args []string) string {
	z0 := cfg.localTZ
	outputTZs, t0, err := parse(z0, args)
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
