package test

import (
	"testing"

	maxioadvancedbilling "github.com/maxio-com/ab-golang-sdk"
)

func TestExample(t *testing.T) {
	// test suite debilu
	config := maxioadvancedbilling.CreateConfiguration(
		maxioadvancedbilling.WithBasicAuthUserName("BasicAuthUserName"),
		maxioadvancedbilling.WithBasicAuthPassword("BasicAuthPassword"),
	)

	t.Log(config)
}
