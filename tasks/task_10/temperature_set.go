package main

import (
	"fmt"
)

func main() {
	// Исходные данные температурных колебаний
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Создаем карту для хранения групп температур
	// Ключом является начало диапазона, а значением — срез температур в этом диапазоне
	groups := make(map[int][]float64)

	// Проходим по каждой температуре и группируем ее
	for _, temp := range temperatures {
		// Вычисляем ключ группы
		// Делим температуру на 10, приводим к целому числу и умножаем обратно на 10
		// Это определяет диапазон, например, для -25.4: int(-25.4/10)*10 = -20
		key := int(temp/10) * 10

		// Добавляем температуру в соответствующую группу
		groups[key] = append(groups[key], temp)
	}

	// Выводим результаты группировки
	for key, temps := range groups {
		fmt.Printf("%d: %v\n", key, temps)
	}
}
