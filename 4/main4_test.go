package main

import (
	"reflect"
	"testing"
)

func TestDifference(t *testing.T) {
	tests := []struct {
		slice1   []string
		slice2   []string
		expected []string
	}{
		{
			slice1:   []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
			slice2:   []string{"banana", "date", "fig"},
			expected: []string{"apple", "cherry", "43", "lead", "gno1"},
		},
		{
			slice1:   []string{"a", "b", "c"},
			slice2:   []string{"b"},
			expected: []string{"a", "c"},
		},
		{
			slice1:   []string{"1", "2", "3"},
			slice2:   []string{"1", "2", "3"},
			expected: []string{},
		},
		{
			slice1:   []string{},
			slice2:   []string{"1"},
			expected: []string{},
		},
		{
			slice1:   []string{"unique"},
			slice2:   []string{},
			expected: []string{"unique"},
		},
	}

	for _, tt := range tests {
		result := Difference(tt.slice1, tt.slice2)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("Difference(%v, %v) = %v; expected %v", tt.slice1, tt.slice2, result, tt.expected)
		}
	}
}
