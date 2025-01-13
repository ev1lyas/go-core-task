package main

import (
	"sync"
)

// Pipeline принимает канал для чтения чисел и канал для записи результатов.
func Pipeline(input <-chan uint8, output chan<- float64, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик WaitGroup при завершении функции

	for num := range input {
		// Преобразуем uint8 в float64 и возводим в куб
		result := float64(num) * float64(num) * float64(num)
		output <- result
	}
	close(output) // Закрываем канал после завершения
}

//func main() {
//	// Создаем каналы
//	input := make(chan uint8)
//	output := make(chan float64)
//
//	// Создаем WaitGroup
//	var wg sync.WaitGroup
//
//	// Увеличиваем счетчик WaitGroup на 1 для горутины конвейера
//	wg.Add(1)
//	go Pipeline(input, output, &wg)
//
//	// Записываем числа в первый канал
//	go func() {
//		for i := uint8(1); i <= 5; i++ {
//			input <- i
//			time.Sleep(100 * time.Millisecond) // Имитируем задержку
//		}
//		close(input) // Закрываем канал после записи всех чисел
//	}()
//
//	// Читаем результаты из второго канала и выводим их
//	for result := range output {
//		fmt.Printf("Результат: %.2f\n", result)
//	}
//
//	// Ожидаем завершения горутины конвейера
//	wg.Wait()
//}
