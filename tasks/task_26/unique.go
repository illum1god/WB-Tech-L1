package main

import (
	"fmt"
	"strings"
)

// Функция для проверки уникальности символов в строке
func areAllCharactersUnique(s string) bool {
	// Создаем карту для отслеживания встреченных символов
	seen := make(map[rune]bool)

	// Проходим по каждому символу строки
	for _, char := range s {
		// Приводим символ к нижнему регистру
		lowerChar := rune(strings.ToLower(string(char))[0])

		// Проверяем, встречался ли символ ранее
		if seen[lowerChar] {
			// Если символ уже встречался, возвращаем false
			return false
		}

		// Отмечаем символ как встреченный
		seen[lowerChar] = true
	}

	// Если все символы уникальны, возвращаем true
	return true
}

func main() {
	// Примеры использования функции
	fmt.Println(areAllCharactersUnique("abcd"))      // true
	fmt.Println(areAllCharactersUnique("abCdefAaf")) // false
	fmt.Println(areAllCharactersUnique("aabcd"))     // false
}
