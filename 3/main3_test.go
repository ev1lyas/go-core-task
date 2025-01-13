package main

import (
	"reflect"
	"testing"
)

// Тест для метода Add
func TestAdd(t *testing.T) {
	m := NewStringIntMap()
	m.Add("one", 1)
	m.Add("two", 2)

	expected := map[string]int{"one": 1, "two": 2}
	if !reflect.DeepEqual(m.data, expected) {
		t.Errorf("Expected %v, got %v", expected, m.data)
	}
}

// Тест для метода Remove
func TestRemove(t *testing.T) {
	m := NewStringIntMap()
	m.Add("one", 1)
	m.Add("two", 2)

	m.Remove("one")
	expected := map[string]int{"two": 2}
	if !reflect.DeepEqual(m.data, expected) {
		t.Errorf("Expected %v, got %v", expected, m.data)
	}
}

// Тест для метода Copy
func TestCopy(t *testing.T) {
	m := NewStringIntMap()
	m.Add("one", 1)
	m.Add("two", 2)

	copyMap := m.Copy()
	if !reflect.DeepEqual(copyMap, m.data) {
		t.Errorf("Expected %v, got %v", m.data, copyMap)
	}

	// Проверяем, что копия не ссылается на оригинал
	m.Add("three", 3)
	if reflect.DeepEqual(copyMap, m.data) {
		t.Errorf("Copy is not independent of the original map")
	}
}

// Тест для метода Exists
func TestExists(t *testing.T) {
	m := NewStringIntMap()
	m.Add("one", 1)

	if !m.Exists("one") {
		t.Errorf("Expected 'one' to exist")
	}

	if m.Exists("two") {
		t.Errorf("Expected 'two' to not exist")
	}
}

// Тест для метода Get
func TestGet(t *testing.T) {
	m := NewStringIntMap()
	m.Add("one", 1)

	value, exists := m.Get("one")
	if !exists || value != 1 {
		t.Errorf("Expected 1, got %d", value)
	}

	_, exists = m.Get("two")
	if exists {
		t.Errorf("Expected 'two' to not exist")
	}
}
