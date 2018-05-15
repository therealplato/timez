package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var errParse = errors.New("could not parse inputs")

var errNoArgs = errors.New("semaphore for default handling")

var nullTime = time.Time{}

func main() {
	var (
		help bool
		a    map[string]string
		c    = &clock{}
		cfg  config
	)
	flag.BoolVar(&help, "help", false, "--help")
	flag.Parse()
	if help {
		fmt.Println(usage)
		os.Exit(0)
	}

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
