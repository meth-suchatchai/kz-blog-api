package kzstring

import (
	"fmt"
	"testing"
)

func TestCombineAll(t *testing.T) {
	t.Log("start combine")
	countryCode := 66
	mobile := "917436969"

	value := CombineAll(fmt.Sprintf("%d", countryCode), mobile)
	if value == "" {
		t.Fatalf("can't retrive to string")
	}

	t.Log(value)
}

func TestReplaceMobileCountryCode(t *testing.T) {
	countryCode := 66
	mobile := 917436969
	x, err := ReplaceMobileCountryCode(countryCode, mobile)
	if err != nil {
		t.Fatalf("error replace mobile number: %v", err)
	}
	t.Log(x)
}
