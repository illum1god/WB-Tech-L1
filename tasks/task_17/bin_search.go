package main

import (
	"fmt"
)

// BinarySearch ищет target в отсортированном массиве arr.
// Возвращает индекс target или -1, если target не найден.
func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1 // Инициализируем левый и правый указатели.

	for left <= right {
		mid := left + (right-left)/2 // Находим средний индекс, избегая переполнения.

		if arr[mid] == target {
			return mid // Элемент найден, возвращаем его индекс.
		} else if arr[mid] < target {
			left = mid + 1 // Ищем в правой половине массива.
		} else {
			right = mid - 1 // Ищем в левой половине массива.
		}
	}

	// Элемент не найден.
	return -1
}

func main() {
	data := []int{1, 3, 5, 7, 9, 11, 13, 15} // Отсортированный массив чисел.
	target := 7                              // Элемент для поиска.

	fmt.Println("Отсортированный массив:", data)

	index := BinarySearch(data, target) // Выполняем бинарный поиск.

	if index != -1 {
		fmt.Printf("Элемент %d найден на индексе %d.\n", target, index)
	} else {
		fmt.Printf("Элемент %d не найден в массиве.\n", target)
	}
}
