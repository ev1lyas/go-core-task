package main

import (
	"reflect"
	"testing"
)

func TestCreateVariables(t *testing.T) {
	numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum := createVariables()

	if numDecimal != 42 {
		t.Errorf("Expected numDecimal to be 42, got %d", numDecimal)
	}
	if numOctal != 42 {
		t.Errorf("Expected numOctal to be 42, got %d", numOctal)
	}
	if numHexadecimal != 42 {
		t.Errorf("Expected numHexadecimal to be 42, got %d", numHexadecimal)
	}
	if pi != 3.14 {
		t.Errorf("Expected pi to be 3.14, got %f", pi)
	}
	if name != "Golang" {
		t.Errorf("Expected name to be 'Golang', got %s", name)
	}
	if isActive != true {
		t.Errorf("Expected isActive to be true, got %v", isActive)
	}
	if complexNum != 1+2i {
		t.Errorf("Expected complexNum to be 1+2i, got %v", complexNum)
	}
}

func TestPrintVariableTypes(t *testing.T) {
	numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum := createVariables()

	types := []reflect.Type{
		reflect.TypeOf(numDecimal),
		reflect.TypeOf(numOctal),
		reflect.TypeOf(numHexadecimal),
		reflect.TypeOf(pi),
		reflect.TypeOf(name),
		reflect.TypeOf(isActive),
		reflect.TypeOf(complexNum),
	}

	for i, v := range []interface{}{numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum} {
		if reflect.TypeOf(v) != types[i] {
			t.Errorf("Expected type %v, got %v", types[i], reflect.TypeOf(v))
		}
	}
}

func TestConvertToCombinedString(t *testing.T) {
	numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum := createVariables()

	expected := "4242423.14Golangtrue(1+2i)"
	result := convertToCombinedString(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)

	if result != expected {
		t.Errorf("Expected combined string %s, got %s", expected, result)
	}
}

func TestStringToRunes(t *testing.T) {
	input := "test"
	expected := []rune{'t', 'e', 's', 't'}

	result := stringToRunes(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected runes %v, got %v", expected, result)
	}
}

func TestHashWithSalt(t *testing.T) {
	tests := []struct {
		input    string
		salt     string
		expected string
	}{
		{"example", "go-2024", "3241b3e810213a81a5ccae9abcbea51e892476d8dbf086d162f327837c14a775"},
		{"", "go-2024", "66802df107aace17871a5b610ff9eb11706e13477bb24e93966ca80671c0fac6"},
		{"test", "", "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			runes := []rune(tt.input)
			result := hashWithSalt(runes, tt.salt)
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}
