package main

import (
	"testing"
)

func TestMergeChannels(t *testing.T) {
	// Создаем несколько каналов
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	// Запускаем горутины для отправки данных в каналы
	go func() {
		defer close(ch1)
		ch1 <- 1
		ch1 <- 2
	}()
	go func() {
		defer close(ch2)
		ch2 <- 3
		ch2 <- 4
	}()
	go func() {
		defer close(ch3)
		ch3 <- 5
		ch3 <- 6
	}()

	// Сливаем каналы
	merged := MergeChannels(ch1, ch2, ch3)

	// Собираем данные из объединенного канала
	var result []int
	for num := range merged {
		result = append(result, num)
	}

	// Проверяем, что все данные были получены
	expected := map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true, 6: true}
	if len(result) != len(expected) {
		t.Errorf("Expected %d numbers, got %d", len(expected), len(result))
	}

	// Проверяем, что все ожидаемые числа присутствуют в результате
	for _, num := range result {
		if !expected[num] {
			t.Errorf("Unexpected number %d in result", num)
		}
	}
}
