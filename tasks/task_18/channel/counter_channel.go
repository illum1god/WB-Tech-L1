package main

import (
	"fmt"
	"sync"
)

// CounterManager управляет счетчиком через канал
type CounterManager struct {
	increment chan struct{}
	value     chan int
}

// NewCounterManager создает новый CounterManager и запускает обработчик
func NewCounterManager() *CounterManager {
	cm := &CounterManager{
		increment: make(chan struct{}),
		value:     make(chan int),
	}
	go cm.run()
	return cm
}

// run обрабатывает входящие запросы на инкремент и запросы значения
func (cm *CounterManager) run() {
	counter := 0
	for {
		select {
		case <-cm.increment:
			counter++
		case cm.value <- counter:
			// Отправляем текущее значение счетчика
		}
	}
}

// Increment отправляет сигнал на инкремент счетчика
func (cm *CounterManager) Increment() {
	cm.increment <- struct{}{}
}

// Value получает текущее значение счетчика
func (cm *CounterManager) Value() int {
	return <-cm.value
}

func main() {
	var wg sync.WaitGroup
	counter := NewCounterManager()

	// Запускаем 1000 горутин для инкрементации счетчика
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
	finalValue := counter.Value()
	fmt.Printf("Итоговое значение счетчика: %d\n", finalValue)
}
