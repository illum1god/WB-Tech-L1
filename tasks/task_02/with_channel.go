package main

import (
	"fmt"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	results := make(chan string, len(numbers)) // Буферизированный канал для хранения результатов

	// Запускаем горутины для вычисления квадратов
	for _, num := range numbers {
		go func(n int) {
			square := n * n
			results <- fmt.Sprintf("Квадрат числа %d равен %d", n, square)
		}(num)
	}

	// Собираем и выводим результаты
	for i := 0; i < len(numbers); i++ {
		fmt.Println(<-results)
	}
}
