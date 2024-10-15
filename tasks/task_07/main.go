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
// Она принимает контекст для обработки сигнала отмены, а также
// записывает обработанные задачи в безопасную для конкурентного доступа карту.
func worker(ctx context.Context, id int, jobs <-chan string, wg *sync.WaitGroup, results map[string]int, mu *sync.Mutex) {
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

			// Запись результата в карту с использованием мьютекса для обеспечения безопасности
			mu.Lock() // Захватываем мьютекс перед записью
			results[job] = id
			mu.Unlock() // Освобождаем мьютекс после записи
		}
	}
}

func main() {
	// Парсим флаг для количества воркеров
	numWorkers := flag.Int("workers", 5, "Количество воркеров")
	// Парсим флаг для длительности выполнения программы
	duration := flag.Int("duration", 10, "Время выполнения программы в секундах")
	flag.Parse()

	// Создаем канал для передачи данных
	jobs := make(chan string)

	// Создаем WaitGroup для ожидания завершения всех воркеров и отправителя задач
	var wg sync.WaitGroup

	// Создаем контекст с возможностью отмены
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Гарантируем вызов отмены при выходе из main

	// Обрабатываем сигналы ОС для graceful shutdown
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt) // Слушаем сигнал прерывания (Ctrl+C)

	// Создаем карту для хранения результатов обработки задач
	results := make(map[string]int)
	// Создаем мьютекс для безопасного доступа к карте в конкурентной среде
	var mu sync.Mutex

	// Запускаем N воркеров
	for i := 1; i <= *numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i, jobs, &wg, results, &mu)
	}

	// Запускаем горутину для постоянной записи данных в канал
	wg.Add(1)
	go func() {
		defer wg.Done()
		count := 1
		for {
			select {
			case <-ctx.Done():
				// Получен сигнал отмены, прекращаем отправку данных
				fmt.Println("Прекращаем отправку данных в канал.")
				close(jobs) // Закрываем канал, чтобы воркеры могли завершиться
				return
			default:
				// Создаем новую задачу
				job := fmt.Sprintf("Задача %d", count)
				select {
				case jobs <- job:
					// Отправляем данные в канал
					fmt.Printf("Отправлено: %s\n", job)
					count++
					time.Sleep(time.Millisecond * 300) // Интервал между отправками
				case <-ctx.Done():
					// Если контекст отменен во время отправки, завершаем работу
					fmt.Println("Прекращаем отправку данных в канал.")
					close(jobs)
					return
				}
			}
		}
	}()

	// Запускаем горутину для завершения программы по истечению заданного времени
	wg.Add(1)
	go func() {
		defer wg.Done()
		// Ждем указанное количество секунд
		time.Sleep(time.Duration(*duration) * time.Second)
		fmt.Printf("\nВремя выполнения программы (%d секунд) истекло. Завершаем работу...\n", *duration)
		cancel() // Вызываем отмену контекста, уведомляя воркеров и остановку отправки
	}()

	// Блокируем основной поток до получения сигнала прерывания или истечения времени
	select {
	case <-sigint:
		fmt.Println("\nПолучен сигнал прерывания. Завершаем работу...")
		cancel() // Вызываем отмену контекста, уведомляя воркеров об остановке
	case <-ctx.Done():
		// Контекст уже отменен горутиной таймера
	}

	// Ожидаем завершения всех горутин
	wg.Wait()

	// Выводим результаты обработки задач
	fmt.Println("\nРезультаты обработки задач:")
	mu.Lock() // Захватываем мьютекс перед чтением карты
	for job, workerID := range results {
		fmt.Printf("%s обработана воркером %d\n", job, workerID)
	}
	mu.Unlock() // Освобождаем мьютекс после чтения

	fmt.Println("\nВсе воркеры завершили работу. Программа завершена.")
}
