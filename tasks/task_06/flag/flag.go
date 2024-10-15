package main

import (
	"fmt"
	"sync"
	"time"
)

// worker функция, которая выполняет работу до изменения флага завершения
func worker(wg *sync.WaitGroup, id int, running *bool) {
	defer wg.Done()
	fmt.Printf("Воркер %d запущен.\n", id)
	for {
		// Проверка условия завершения
		if !*running {
			fmt.Printf("Воркер %d завершает работу.\n", id)
			return
		}

		// Симуляция работы
		fmt.Printf("Воркер %d выполняет задачу.\n", id)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	var (
		numWorkers = 3
		running    = true
		wg         sync.WaitGroup
	)

	// Запуск воркеров
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(&wg, i, &running)
	}

	// Позволяем воркерам работать некоторое время
	time.Sleep(5 * time.Second)

	// Изменяем флаг завершения
	running = false

	wg.Wait()
	fmt.Println("Все воркеры завершили работу. Программа завершена.")
}
