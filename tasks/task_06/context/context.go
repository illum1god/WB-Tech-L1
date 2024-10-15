package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// worker функция, которая выполняет работу до получения сигнала отмены контекста
func worker(wg *sync.WaitGroup, ctx context.Context, id int) {
	defer wg.Done()
	fmt.Printf("Воркер %d запущен.\n", id)
	for {
		select {
		case <-ctx.Done():
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
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	numWorkers := 3

	// Запуск воркеров
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(&wg, ctx, i)
	}

	// Позволяем воркерам работать некоторое время
	time.Sleep(5 * time.Second)

	// Отмена контекста, сигнализируя воркерам о необходимости завершиться
	cancel()

	wg.Wait()
	fmt.Println("Все воркеры завершили работу. Программа завершена.")
}
