package cmd

import (
	"testing"
)

func Test_Add(t *testing.T) {
	var args = []string{"24", "+", "2"}
	result := processArgs(args)
	if result != 26 {
		t.Errorf("%f != 26", result)
	}
}

func Test_Div(t *testing.T) {
	var args = []string{"24", "/", "2"}
	result := processArgs(args)
	if result != 12 {
		t.Errorf("%f != 12", result)
	}
}

func Test_ComplexCalc(t *testing.T) {
	var args = []string{"24", "+", "2", "-", "10"}
	result := processArgs(args)
	if result != 16 {
		t.Errorf("%f != 16", result)
	}
}
