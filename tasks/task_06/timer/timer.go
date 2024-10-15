package main

import (
	"fmt"
	"sync"
	"time"
)

// worker функция, которая выполняет работу до получения сигнала отмены контекста
func worker(wg *sync.WaitGroup, seconds time.Duration, id int) {
	defer wg.Done()
	fmt.Printf("Воркер %d запущен.\n", id)
	timer := time.NewTimer(seconds)
	for {
		select {
		case <-timer.C:
			fmt.Printf("Воркер %d завершает работу.\n", id)
			return
		default:
			// Симуляция работы
			fmt.Printf("Воркер %d выполняет задачу.\n", id)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	// Создаем контекст с возможностью отмены
	var wg sync.WaitGroup // WaitGroup для отслеживания работающих случаев
	duration := 4 * time.Second
	numWorkers := 3

	// Запуск воркеров
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(&wg, duration, i)
	}

	// Позволяем воркерам работать некоторое время
	time.Sleep(5 * time.Second)

	wg.Wait()

	fmt.Println("Все воркеры завершили работу. Программа завершена.")
}
