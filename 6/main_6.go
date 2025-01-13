package main

import (
	"fmt"
	"math/rand"
	"time"
)

// RandomNumberGenerator создает генератор случайных чисел через небуферизированный канал
func RandomNumberGenerator() <-chan int {
	ch := make(chan int)
	go func() {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for {
			ch <- r.Int()
		}
	}()
	return ch
}

func main() {
	randomNumbers := RandomNumberGenerator()

	for i := 0; i < 5; i++ {
		fmt.Println(<-randomNumbers)
	}
}
