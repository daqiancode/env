package env_test

import (
	"testing"

	"github.com/daqiancode/env"
)

func TestGet(t *testing.T) {
	if env.Get("app") != "env" {
		t.Fail()
	}
	if env.GetIntMust("int") != 1 {
		t.Error("GetenvIntMust failed")
	}
	if env.GetFloatMust("float") != 1.2 {
		t.Error("GetenvFloat64 failed")
	}
}
