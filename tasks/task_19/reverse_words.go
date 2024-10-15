package main

import (
	"fmt"
)

// reverseString переворачивает заданную строку, закодированную в UTF-8
func reverseString(input string) string {
	// Преобразуем строку в срез рун для обработки символов Unicode
	runes := []rune(input)

	// Переворачиваем срез рун
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Преобразуем перевернутый срез рун обратно в строку
	return string(runes)
}

func main() {
	// Пример входной строки
	input := "главрыба"

	// Переворачиваем строку
	reversed := reverseString(input)

	// Выводим перевернутую строку
	fmt.Printf("Оригинал: %s\nПеревернуто: %s\n", input, reversed)
}
