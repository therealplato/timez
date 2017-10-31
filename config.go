package main

import (
	"io"
	"io/ioutil"
	"log"
	"time"

	yaml "gopkg.in/yaml.v2"
)

type config struct {
	localTZ *time.Location
	aliases map[string]string
}

func mustLoadAliases(in io.Reader) map[string]string {
	bb, err := ioutil.ReadAll(in)
	if err != nil {
		log.Fatalf("couldn't load config: %s", err)
	}
	var a = make(map[string]string)
	err = yaml.Unmarshal(bb, &a)
	if err != nil {
		log.Fatalf("couldn't unmarshal config: %s", err)
	}
	for k, v := range tzAlias {
		_, exists := a[k]
		if !exists {
			a[k] = v
		}
	}

	return a
}
