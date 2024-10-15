package main

import (
	"fmt"
	"strings"
)

// reverseWords принимает строку и возвращает строку с перевернутыми словами
func reverseSentence(s string) string {
	// Разбиваем строку на слова по пробелам
	words := strings.Fields(s)

	// Инициализируем два указателя для переворота списка слов
	left, right := 0, len(words)-1

	// Переворачиваем слова местами
	for left < right {
		words[left], words[right] = words[right], words[left]
		left++
		right--
	}

	// Объединяем перевернутые слова обратно в строку
	return strings.Join(words, " ")
}

func main() {
	// Пример использования функции reverseWords
	input := "snow dog sun"
	output := reverseSentence(input)
	fmt.Println(output) // Вывод: "sun dog snow"
}
