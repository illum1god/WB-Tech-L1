package main

import (
	"fmt"
	"time"
)

// worker функция, которая выполняет работу до возникновения паники
func worker(id int) {
	fmt.Printf("Воркер %d запущен.\n", id)
	for i := 1; i <= 3; i++ {
		fmt.Printf("Воркер %d выполняет задачу %d.\n", id, i)
		time.Sleep(1 * time.Second)
		if i == 2 && id == 2 {
			panic(fmt.Sprintf("Воркер %d столкнулся с критической ошибкой!", id))
		}
	}
	fmt.Printf("Воркер %d завершает работу.\n", id)
}

func main() {
	numWorkers := 3

	// Запуск воркеров
	for i := 1; i <= numWorkers; i++ {
		go worker(i)
	}

	// Позволяем воркерам выполнить работу
	time.Sleep(5 * time.Second)
	fmt.Println("Все воркеры завершили работу. Программа завершена.")
}
