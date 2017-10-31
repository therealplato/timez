package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	f := bytes.NewBuffer([]byte(`
default: Pacific/Auckland
freya: UTC
thor: asgard
`))

	var x = make(map[string]string)
	for k, v := range defaultAlias {
		x[k] = v
	}
	x["default"] = "Pacific/Auckland"
	x["freya"] = "UTC"
	x["thor"] = "asgard"

	actual := mustLoadAliases(f)
	assert.Equal(t, x, actual)
}
