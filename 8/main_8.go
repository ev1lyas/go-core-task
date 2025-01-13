package main

import (
	"sync"
)

// WaitGroup представляет кастомную реализацию WaitGroup на основе семафора.
type WaitGroup struct {
	counter int32         // Счетчик активных горутин
	done    chan struct{} // Канал для сигнализации завершения
	mu      sync.Mutex    // Мьютекс для защиты от гонок
	closed  bool          // Флаг, указывающий, был ли канал закрыт
}

// NewWaitGroup создает новый экземпляр WaitGroup.
func NewWaitGroup() *WaitGroup {
	return &WaitGroup{
		done: make(chan struct{}),
	}
}

// Add добавляет delta к счетчику WaitGroup.
func (wg *WaitGroup) Add(delta int) {
	wg.mu.Lock()
	defer wg.mu.Unlock()

	newCounter := wg.counter + int32(delta)
	if newCounter < 0 {
		panic("negative WaitGroup counter")
	}
	wg.counter = newCounter

	if wg.counter == 0 && !wg.closed {
		close(wg.done) // Закрываем канал, если он еще не закрыт
		wg.closed = true
	}
}

// Done уменьшает счетчик WaitGroup на 1.
func (wg *WaitGroup) Done() {
	wg.Add(-1)
}

// Wait блокирует выполнение, пока счетчик WaitGroup не станет равным 0.
func (wg *WaitGroup) Wait() {
	<-wg.done
}
