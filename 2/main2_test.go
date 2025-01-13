package main

import (
	"reflect"
	"testing"
)

func TestGenerateRandomSlice(t *testing.T) {
	slice := generateRandomSlice(10)

	// Проверяем, что длина слайса равна 10
	if len(slice) != 10 {
		t.Errorf("Expected slice length 10, got %d", len(slice))
	}

	// Проверяем, что все элементы в слайсе находятся в диапазоне от 0 до 99
	for _, num := range slice {
		if num < 0 || num >= 100 {
			t.Errorf("Expected numbers in range 0-99, got %d", num)
		}
	}
}

func TestSliceExample(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{2, 4}},
		{[]int{10, 11, 12, 13, 14}, []int{10, 12, 14}},
		{[]int{1, 3, 5, 7, 9}, []int{}},
		{[]int{}, []int{}},
	}

	for _, tt := range tests {
		result := sliceExample(tt.input)

		// Проверяем длину слайсов
		if len(result) != len(tt.expected) {
			t.Errorf("For input %v, expected length %d, got length %d", tt.input, len(tt.expected), len(result))
		}

		// Проверяем содержимое слайсов
		for i, val := range result {
			if val != tt.expected[i] {
				t.Errorf("For input %v, expected %v, got %v", tt.input, tt.expected, result)
				break
			}
		}
	}
}

func TestAddElements(t *testing.T) {
	tests := []struct {
		input    []int
		num      int
		expected []int
	}{
		{[]int{1, 2, 3}, 4, []int{1, 2, 3, 4}},
		{[]int{}, 10, []int{10}},
		{[]int{5, 6}, 7, []int{5, 6, 7}},
	}

	for _, tt := range tests {
		result := addElements(tt.input, tt.num)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("For input %v and num %d, expected %v, got %v", tt.input, tt.num, tt.expected, result)
		}
	}
}

func TestCopySlice(t *testing.T) {
	tests := []struct {
		input []int
	}{
		{[]int{1, 2, 3}},
		{[]int{10, 20, 30}},
		{[]int{}},
	}

	for _, tt := range tests {
		copied := copySlice(tt.input)
		if !reflect.DeepEqual(copied, tt.input) {
			t.Errorf("For input %v, expected %v, got %v", tt.input, tt.input, copied)
		}
		// Проверяем, что копия не ссылается на оригинал
		if len(tt.input) > 0 { // Проверяем, что слайс не пустой
			tt.input[0] = 99
			if copied[0] == 99 {
				t.Errorf("Copy is not independent of the original slice")
			}
		}
	}
}

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		input    []int
		index    int
		expected []int
	}{
		{[]int{1, 2, 3, 4}, 2, []int{1, 2, 4}},
		{[]int{10, 20, 30}, 0, []int{20, 30}},
		{[]int{5, 6, 7}, 2, []int{5, 6}},
		{[]int{1}, 0, []int{}},
		{[]int{1, 2, 3}, 5, []int{1, 2, 3}}, // Индекс за пределами слайса
	}

	for _, tt := range tests {
		result := removeElement(tt.input, tt.index)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("For input %v and index %d, expected %v, got %v", tt.input, tt.index, tt.expected, result)
		}
	}
}
