package main

import (
	"gotest.tools/assert"
	"testing"
)

func TestGetConfiguration(t *testing.T) {

	conf := getConfiguration()

	w := conf.Weather

	assert.Check(t, w.URL != "")
	assert.Check(t, w.APIKey != "")
}
