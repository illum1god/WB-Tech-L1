package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

// worker функция, которая читает данные из канала и выводит их в stdout.
// Она принимает контекст для обработки сигнала отмены.
func worker(ctx context.Context, id int, jobs <-chan string, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик WaitGroup при завершении воркера

	for {
		select {
		case <-ctx.Done():
			// Получен сигнал отмены, завершаем работу воркера
			fmt.Printf("Воркер %d завершает работу.\n", id)
			return
		case job, ok := <-jobs:
			if !ok {
				// Канал закрыт, завершаем работу воркера
				fmt.Printf("Воркер %d обнаружил закрытый канал и завершает работу.\n", id)
				return
			}
			// Обработка полученных данных
			fmt.Printf("Воркер %d обработал: %s\n", id, job)
			// Симулируем время обработки
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func main() {
	// Определяем флаги командной строки
	numWorkers := flag.Int("workers", 5, "Количество воркеров")
	duration := flag.Int("duration", 10, "Время выполнения программы в секундах")
	flag.Parse()

	// Создаем канал для передачи данных
	jobs := make(chan string)

	// Создаем WaitGroup для ожидания завершения всех воркеров
	var wg sync.WaitGroup

	// Создаем контекст с возможностью отмены
	ctx, cancel := context.WithCancel(context.Background())

	// Обрабатываем сигналы ОС для graceful shutdown (например, Ctrl+C)
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt) // Слушаем сигнал прерывания (Ctrl+C)

	// Запускаем N воркеров
	for i := 1; i <= *numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i, jobs, &wg)
	}

	// Горутина для постоянной записи данных в канал
	go func() {
		count := 1
		for {
			job := fmt.Sprintf("Задача %d", count)
			select {
			case <-ctx.Done():
				// Получен сигнал отмены, прекращаем отправку данных
				fmt.Println("Прекращаем отправку данных в канал.")
				close(jobs) // Закрываем канал, чтобы воркеры могли завершиться
				return
			case jobs <- job:
				// Отправляем данные в канал
				fmt.Printf("Отправлено: %s\n", job)
				count++
				time.Sleep(time.Millisecond * 300) // Интервал между отправками
			}
		}
	}()

	// Горутина для завершения программы по истечению заданного времени
	go func() {
		time.Sleep(time.Duration(*duration) * time.Second) // Ждем N секунд
		fmt.Printf("\nВремя выполнения программы (%d секунд) истекло. Завершаем работу...\n", *duration)
		cancel() // Вызываем отмену контекста, уведомляя воркеров и остановку записи
	}()

	// Блокируем основной поток до получения сигнала прерывания или истечения времени
	select {
	case <-sigint:
		fmt.Println("\nПолучен сигнал прерывания. Завершаем работу...")
	case <-ctx.Done():
		// Контекст уже отменен горутиной таймера
	}

	// Ожидаем завершения всех воркеров
	wg.Wait()

	fmt.Println("Все воркеры завершили работу. Программа завершена.")
}
