package main

import (
	"fmt"
)

// StringIntMap - структура для хранения пар "строка - целое число"
type StringIntMap struct {
	data map[string]int
}

// NewStringIntMap создает новую карту
func NewStringIntMap() *StringIntMap {
	return &StringIntMap{
		data: make(map[string]int),
	}
}

// Add добавляет пару "ключ-значение" в карту
func (m *StringIntMap) Add(key string, value int) {
	m.data[key] = value
}

// Remove удаляет элемент по ключу из карты
func (m *StringIntMap) Remove(key string) {
	delete(m.data, key)
}

// Copy возвращает копию карты
func (m *StringIntMap) Copy() map[string]int {
	copyMap := make(map[string]int)
	for key, value := range m.data {
		copyMap[key] = value
	}
	return copyMap
}

// Exists проверяет, существует ли ключ в карте
func (m *StringIntMap) Exists(key string) bool {
	_, exists := m.data[key]
	return exists
}

// Get возвращает значение по ключу и флаг успешности операции
func (m *StringIntMap) Get(key string) (int, bool) {
	if value, exists := m.data[key]; exists {
		return value, exists
	}
	return 0, false
}

// String реализует интерфейс fmt.Stringer для удобного вывода
func (m *StringIntMap) String() string {
	return fmt.Sprintf("%v", m.data)
}

func main() {
	m := NewStringIntMap()
	m.Add("one", 1)
	m.Add("two", 2)
	fmt.Println("Map after Add:", m)
	fmt.Println("Exists 'one':", m.Exists("one"))

	copied := m.Copy()
	fmt.Println("Copied map:", copied)

	m.Remove("one")
	fmt.Println("'m' after remove 'one':", m)
	fmt.Println("'copied' after remove 'one':", copied)
}
