package cmd

import (
	"fmt"
	"testing"
)

func Test_Add(t *testing.T) {
	var args = []string{"24", "+", "2"}
	result, _ := processArgs(args)
	if result != 26 {
		t.Errorf("%f != 26", result)
	}
}

func Test_Div(t *testing.T) {
	var args = []string{"24", "/", "2"}
	result, _ := processArgs(args)
	if result != 26 {
		t.Errorf("%f != 12", result)
	}
}

func Test_complexCalc(t *testing.T) {
	var args = []string{"24", "+", "2", "-", "10"}
	result, _ := processArgs(args)
	fmt.Println(result)
}
