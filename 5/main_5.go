package main

import "fmt"

// Intersection проверяет пересечения между двумя слайсами и возвращает bool и пересечения
func Intersection(a, b []int) (bool, []int) {
	set := make(map[int]struct{})
	for _, num := range a {
		set[num] = struct{}{}
	}

	result := []int{}
	found := false
	for _, num := range b {
		if _, exists := set[num]; exists {
			result = append(result, num)
			found = true
		}
	}

	return found, result
}

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}
	found, intersect := Intersection(a, b)
	fmt.Println(found, intersect) // true, []int{64, 3}
}
