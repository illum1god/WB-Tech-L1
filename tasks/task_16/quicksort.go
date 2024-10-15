package main

import (
	"fmt"
)

// QuickSort сортирует переданный массив чисел и возвращает отсортированный массив.
// Алгоритм использует метод быстрой сортировки с рекурсивным разделением массива.
func QuickSort(arr []int) []int {
	// Базовый случай: массив с менее чем 2 элементами уже отсортирован.
	if len(arr) < 2 {
		return arr
	}

	// Выбираем опорный элемент (pivot) как последний элемент массива.
	pivot := arr[len(arr)-1]

	// Инициализируем индекс для элементов, меньших опорного.
	left := 0

	// Проходим по массиву, переставляя элементы относительно опорного.
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < pivot {
			// Меняем местами текущий элемент с элементом на позиции 'left'.
			arr[i], arr[left] = arr[left], arr[i]
			left++ // Увеличиваем индекс 'left' для следующего меньшего элемента.
		}
	}

	// Перемещаем опорный элемент на его правильную позицию.
	arr[left], arr[len(arr)-1] = arr[len(arr)-1], arr[left]

	// Рекурсивно сортируем левую часть массива (элементы меньше pivot).
	QuickSort(arr[:left])

	// Рекурсивно сортируем правую часть массива (элементы больше pivot).
	QuickSort(arr[left+1:])

	return arr // Возвращаем отсортированный массив.
}

func main() {
	// Инициализируем неотсортированный массив чисел.
	data := []int{10, 7, 8, 9, 1, 5}
	fmt.Println("Неотсортированный массив:", data)

	// Вызываем функцию QuickSort для сортировки массива.
	sorted := QuickSort(data)

	// Выводим отсортированный массив на стандартный вывод.
	fmt.Println("Отсортированный массив:", sorted)
}
