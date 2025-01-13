package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
)

// 1. Создаем переменные различных типов данных
func createVariables() (int, int, int, float64, string, bool, complex64) {
	var numDecimal int = 42
	var numOctal int = 052
	var numHexadecimal int = 0x2A
	var pi float64 = 3.14
	var name string = "Golang"
	var isActive bool = true
	var complexNum complex64 = 1 + 2i

	return numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum
}

// 2. Определяем тип каждой переменной и выводим его на экран
func printVariableTypes(vars ...interface{}) {
	for _, v := range vars {
		fmt.Printf("Type of %v: %T\n", v, v)
	}
}

// 3. Преобразуем все переменные в строковый тип и объединяем их в одну строку
func convertToCombinedString(vars ...interface{}) string {
	var combinedString string
	for _, v := range vars {
		switch val := v.(type) {
		case int:
			combinedString += strconv.Itoa(val)
		case float64:
			combinedString += strconv.FormatFloat(val, 'f', -1, 64)
		case string:
			combinedString += val
		case bool:
			combinedString += strconv.FormatBool(val)
		case complex64:
			combinedString += fmt.Sprintf("%v", val)
		}
	}
	return combinedString
}

// 4. Преобразуем строку в срез рун
func stringToRunes(s string) []rune {
	return []rune(s)
}

// 5. Захэшируем срез рун SHA256, добавив в середину соль "go-2024"
func hashWithSalt(runes []rune, salt string) string {
	mid := len(runes) / 2
	hashedRunes := append(runes[:mid], append([]rune(salt), runes[mid:]...)...)
	hashedString := string(hashedRunes)

	hash := sha256.Sum256([]byte(hashedString))
	return hex.EncodeToString(hash[:])
}

func main() {
	// Создаем переменные
	numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum := createVariables()

	// Выводим типы переменных
	printVariableTypes(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)

	// Преобразуем переменные в строку
	combinedString := convertToCombinedString(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)

	// Преобразуем строку в срез рун
	runes := stringToRunes(combinedString)

	// Захэшируем срез рун с солью
	salt := "go-2024"
	hashHex := hashWithSalt(runes, salt)

	// Выводим результат
	fmt.Println("SHA256 hash:", hashHex)
}
