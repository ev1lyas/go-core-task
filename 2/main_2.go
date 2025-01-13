package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 1. Создание слайса случайных чисел
func generateRandomSlice(size int) []int {
	rand.Seed(time.Now().UnixNano())
	slice := make([]int, size)
	for i := range slice {
		slice[i] = rand.Intn(100) // Генерация случайного числа от 0 до 99
	}
	return slice
}

// 2. Возвращает новый слайс только с четными числами
func sliceExample(slice []int) []int {
	var result []int
	for _, num := range slice {
		if num%2 == 0 {
			result = append(result, num)
		}
	}
	return result
}

// 3. Добавляет число в конец слайса
func addElements(slice []int, num int) []int {
	return append(slice, num)
}

// 4. Возвращает копию слайса
func copySlice(slice []int) []int {
	copySlice := make([]int, len(slice))
	copy(copySlice, slice)
	return copySlice
}

// 5. Функция removeElement: удаляет элемент по индексу
func removeElement(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		fmt.Println("Индекс выходит за пределы слайса")
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	// 1. Создаем слайс случайных чисел
	originalSlice := generateRandomSlice(10)
	fmt.Println("Original Slice:", originalSlice)

	// 2. Получаем слайс только с четными числами
	evenSlice := sliceExample(originalSlice)
	fmt.Println("Even Numbers:", evenSlice)

	// 3. Добавляем число в конец слайса
	modifiedSlice := addElements(originalSlice, 100)
	fmt.Println("Modified Slice with Added Element:", modifiedSlice)

	// 4. Создаем копию слайса
	copiedSlice := copySlice(originalSlice)
	fmt.Println("Copied Slice:", copiedSlice)

	// 5. Удаляем элемент по индексу
	removedSlice := removeElement(originalSlice, 2)
	fmt.Println("Slice with Removed Element:", removedSlice)
}
