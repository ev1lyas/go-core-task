package main

import (
	"sync"
)

// MergeChannels сливает N каналов в один
func MergeChannels(channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Функция для копирования данных из одного канала в выходной канал
	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			out <- n
		}
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go output(c)
	}

	// Запускаем горутину для закрытия выходного канала после завершения всех горутин
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
