package main

import (
	"fmt"
)

// IntersectionMap возвращает пересечение двух неупорядоченных множеств (set1 и set2) с использованием карты.
// Параметры:
// - set1: первый срез целых чисел.
// - set2: второй срез целых чисел.
// Возвращает:
// - срез целых чисел, представляющий пересечение set1 и set2.
func IntersectionMap(set1, set2 []int) []int {
	// Создаём карту для хранения элементов первого множества.
	elementMap := make(map[int]bool)
	for _, elem := range set1 {
		elementMap[elem] = true // Помечаем элемент как присутствующий в set1.
	}

	// Инициализируем срез для хранения пересечения.
	intersection := []int{}
	for _, elem := range set2 {
		if elementMap[elem] {
			intersection = append(intersection, elem) // Добавляем элемент в пересечение.
			delete(elementMap, elem)                  // Удаляем элемент из карты, чтобы избежать дубликатов.
		}
	}

	return intersection
}

func main() {
	// Определяем два набора чисел.
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{4, 5, 6, 7, 8}

	// Вызываем функцию IntersectionMap для получения пересечения set1 и set2.
	result := IntersectionMap(set1, set2)

	// Выводим результат на стандартный вывод.
	fmt.Println("Пересечение (map):", result)
}
