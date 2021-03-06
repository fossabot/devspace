package paramutil

import (
	"testing"
)

//This test struct is from stdinutil
type GetFromStdin_params struct {
	Question               string
	DefaultValue           string
	ValidationRegexPattern string
	InputTerminationString string
}

func TestSetDefaults(t *testing.T) {

	defaultParams := &GetFromStdin_params{
		ValidationRegexPattern: ".*",
		InputTerminationString: "\n",
	}

	params := &GetFromStdin_params{
		ValidationRegexPattern: "",
		InputTerminationString: " ",
	}

	SetDefaults(params, defaultParams)

	if defaultParams.ValidationRegexPattern != ".*" || defaultParams.InputTerminationString != "\n" {
		t.Error("defaultParams changed during method call")
		t.Fail()
	}

	if params.ValidationRegexPattern != ".*" {
		t.Error("empty param isn't set to default")
		t.Fail()
	}

	if params.InputTerminationString != " " {
		t.Error("Non-empty param is set to default")
		t.Fail()
	}

	
}
