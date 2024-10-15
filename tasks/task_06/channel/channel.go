package main

import (
	"fmt"
	"sync"
	"time"
)

// worker функция, которая выполняет работу до получения сигнала остановки
func worker(wg *sync.WaitGroup, id int, stopChan <-chan struct{}) {
	defer wg.Done()
	fmt.Printf("Воркер %d запущен.\n", id)
	for {
		select {
		case <-stopChan:
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
	stopChan := make(chan struct{})
	numWorkers := 3
	var wg sync.WaitGroup

	// Запуск воркеров
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(&wg, i, stopChan)
	}

	// Позволяем воркерам работать некоторое время
	time.Sleep(5 * time.Second)

	// Сигнал для остановки всех воркеров
	close(stopChan)

	wg.Wait()
	fmt.Println("Все воркеры завершили работу. Программа завершена.")
}
