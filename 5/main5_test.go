package main

import (
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T) {
	tests := []struct {
		a        []int
		b        []int
		exists   bool
		expected []int
	}{
		{
			a:        []int{65, 3, 58, 678, 64},
			b:        []int{64, 2, 3, 43},
			exists:   true,
			expected: []int{64, 3},
		},
		{
			a:        []int{1, 2, 3},
			b:        []int{4, 5, 6},
			exists:   false,
			expected: []int{},
		},
		{
			a:        []int{},
			b:        []int{1, 2, 3},
			exists:   false,
			expected: []int{},
		},
		{
			a:        []int{10, 20, 30},
			b:        []int{30, 40, 50},
			exists:   true,
			expected: []int{30},
		},
		{
			a:        []int{7, 8, 9},
			b:        []int{},
			exists:   false,
			expected: []int{},
		},
	}

	for _, tt := range tests {
		found, result := Intersection(tt.a, tt.b)
		if found != tt.exists || !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("Intersection(%v, %v) = (%v, %v); expected (%v, %v)", tt.a, tt.b, found, result, tt.exists, tt.expected)
		}
	}
}
