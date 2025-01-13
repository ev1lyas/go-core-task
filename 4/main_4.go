package main

import "fmt"

// Difference возвращает элементы первого слайса, которых нет во втором
func Difference(slice1, slice2 []string) []string {
	set := make(map[string]struct{})
	for _, item := range slice2 {
		set[item] = struct{}{}
	}

	result := []string{}
	for _, item := range slice1 {
		if _, found := set[item]; !found {
			result = append(result, item)
		}
	}

	return result
}

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	result := Difference(slice1, slice2)
	fmt.Println(result)
}
