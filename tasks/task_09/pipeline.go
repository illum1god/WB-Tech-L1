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

// producer функция, которая читает числа из массива и отправляет их в канал.
// Она принимает контекст для обработки сигнала отмены, массив чисел,
// выходной канал для отправки чисел и WaitGroup для синхронизации завершения.
func producer(ctx context.Context, numbers []int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик WaitGroup при завершении продюсера

	for _, num := range numbers {
		select {
		case <-ctx.Done():
			// Получен сигнал отмены, прекращаем работу продюсера
			fmt.Println("Прекращаем работу продюсера.")
			return
		case out <- num:
			// Отправляем число в канал
			fmt.Printf("Произведено число: %d\n", num)
			time.Sleep(time.Millisecond * 200) // Симулируем задержку производства
		}
	}

	// После отправки всех чисел сигнализируем о завершении, закрывая канал
	close(out)
	fmt.Println("Все числа произведены. Продюсер завершил работу.")
}

// processor функция, которая читает числа из входного канала,
// удваивает их и отправляет в выходной канал.
// Она принимает контекст для обработки сигнала отмены,
// входной и выходной каналы, а также WaitGroup.
func processor(ctx context.Context, in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик WaitGroup при завершении процессора

	for {
		select {
		case <-ctx.Done():
			// Получен сигнал отмены, прекращаем работу процессора
			fmt.Println("Прекращаем работу процессора.")
			return
		case num, ok := <-in:
			if !ok {
				// Входной канал закрыт, завершение работы процессора
				fmt.Println("Входной канал закрыт. Процессор завершает работу.")
				return
			}
			// Обрабатываем число (удваиваем)
			processedNum := num * 2
			fmt.Printf("Обработано число: %d -> %d\n", num, processedNum)

			select {
			case <-ctx.Done():
				// Проверяем сигнал отмены перед отправкой
				fmt.Println("Прекращаем работу процессора перед отправкой.")
				return
			case out <- processedNum:
				// Отправляем обработанное число в выходной канал
			}
		}
	}
}

// consumer функция, которая читает числа из входного канала и выводит их в stdout.
// Она принимает контекст для обработки сигнала отмены, входной канал и WaitGroup.
func consumer(ctx context.Context, in <-chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик WaitGroup при завершении потребителя

	for {
		select {
		case <-ctx.Done():
			// Получен сигнал отмены, прекращаем работу потребителя
			fmt.Println("Прекращаем работу потребителя.")
			return
		case num, ok := <-in:
			if !ok {
				// Входной канал закрыт, завершение работы потребителя
				fmt.Println("Входной канал закрыт. Потребитель завершает работу.")
				return
			}
			// Выводим полученное число
			fmt.Printf("Потреблено число: %d\n", num)
		}
	}
}

func main() {
	// Парсим флаги командной строки
	numWorkers := flag.Int("workers", 3, "Количество воркеров (процессоров)")
	duration := flag.Int("duration", 10, "Время выполнения программы в секундах")
	flag.Parse()

	// Массив чисел для производства
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Создаем каналы для передачи данных
	produceChan := make(chan int) // Канал для производителей
	processChan := make(chan int) // Канал для процессоров

	// Создаем WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Создаем контекст с возможностью отмены
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Гарантируем вызов отмены при выходе из main

	// Обрабатываем сигналы ОС для graceful shutdown
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt) // Слушаем сигнал прерывания (Ctrl+C)

	// Запускаем продюсера
	wg.Add(1)
	go producer(ctx, numbers, produceChan, &wg)

	// Запускаем процессоров (воркеров)
	for i := 1; i <= *numWorkers; i++ {
		wg.Add(1)
		go processor(ctx, produceChan, processChan, &wg)
	}

	// Запускаем горутину для закрытия processChan после завершения всех процессоров
	go func() {
		wg.Wait()          // Ждем завершения всех процессоров и продюсера
		close(processChan) // Закрываем канал процессоров один раз
	}()

	// Запускаем потребителя
	wg.Add(1)
	go consumer(ctx, processChan, &wg)

	// Запускаем таймер для автоматической остановки программы
	go func() {
		time.Sleep(time.Duration(*duration) * time.Second)
		fmt.Printf("\nВремя выполнения программы (%d секунд) истекло. Завершаем работу...\n", *duration)
		cancel() // Вызываем отмену контекста, уведомляя горутины об остановке
	}()

	// Блокируем основной поток до получения сигнала прерывания или таймера
	select {
	case <-sigint:
		fmt.Println("\nПолучен сигнал прерывания. Завершаем работу...")
		cancel() // Вызываем отмену контекста, уведомляя горутины об остановке
	case <-ctx.Done():
		// Контекст уже отменен горутиной таймера
	}

	// Ожидаем завершения всех горутин
	wg.Wait()

	fmt.Println("Программа завершена.")
}
