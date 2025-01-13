package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

// TestWaitGroupConcurrentAdd проверяет конкурентное добавление и завершение горутин.
func TestWaitGroupConcurrentAdd(t *testing.T) {
	wg := NewWaitGroup()
	const numGoroutines = 10

	var counter int32
	var wgStart sync.WaitGroup // WaitGroup для синхронизации старта горутин

	wgStart.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			wgStart.Done() // Сигнализируем, что горутина начала выполнение
			wgStart.Wait() // Ожидаем, пока все горутины начнут выполнение
			atomic.AddInt32(&counter, 1)
		}()
	}

	// Увеличиваем счетчик WaitGroup на количество горутин
	wg.Add(numGoroutines)

	wg.Wait() // Ожидаем завершения всех горутин

	if counter != numGoroutines {
		t.Errorf("Expected counter to be %d, but got %d", numGoroutines, counter)
	}
}
