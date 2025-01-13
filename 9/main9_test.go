package main

import (
	"sync"
	"testing"
)

// TestPipeline проверяет корректность работы конвейера.
func TestPipeline(t *testing.T) {
	input := make(chan uint8)
	output := make(chan float64)

	// Создаем WaitGroup
	var wg sync.WaitGroup

	// Увеличиваем счетчик WaitGroup на 1 для горутины конвейера
	wg.Add(1)
	go Pipeline(input, output, &wg)

	// Записываем числа в первый канал
	go func() {
		input <- 2
		input <- 3
		input <- 4
		close(input)
	}()

	// Ожидаемые результаты
	expected := []float64{8, 27, 64}
	index := 0

	// Читаем результаты из второго канала и сравниваем с ожидаемыми
	for result := range output {
		if result != expected[index] {
			t.Errorf("Ожидалось %.2f, получено %.2f", expected[index], result)
		}
		index++
	}

	// Ожидаем завершения горутины конвейера
	wg.Wait()
}

// TestPipelineWithZero проверяет обработку нуля.
func TestPipelineWithZero(t *testing.T) {
	input := make(chan uint8)
	output := make(chan float64)

	// Создаем WaitGroup
	var wg sync.WaitGroup

	// Увеличиваем счетчик WaitGroup на 1 для горутины конвейера
	wg.Add(1)
	go Pipeline(input, output, &wg)

	go func() {
		input <- 0
		close(input)
	}()

	expected := float64(0)
	result := <-output

	if result != expected {
		t.Errorf("Ожидалось %.2f, получено %.2f", expected, result)
	}

	// Ожидаем завершения горутины конвейера
	wg.Wait()
}

// TestPipelineWithLargeNumber проверяет обработку большого числа.
func TestPipelineWithLargeNumber(t *testing.T) {
	input := make(chan uint8)
	output := make(chan float64)

	// Создаем WaitGroup
	var wg sync.WaitGroup

	// Увеличиваем счетчик WaitGroup на 1 для горутины конвейера
	wg.Add(1)
	go Pipeline(input, output, &wg)

	go func() {
		input <- 10
		close(input)
	}()

	expected := float64(1000)
	result := <-output

	if result != expected {
		t.Errorf("Ожидалось %.2f, получено %.2f", expected, result)
	}

	// Ожидаем завершения горутины конвейера
	wg.Wait()
}
