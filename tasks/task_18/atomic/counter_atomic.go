package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// AtomicCounter структура-счетчик с использованием атомарных операций
type AtomicCounter struct {
	value int64
}

// Increment увеличивает значение счетчика на 1 атомарно
func (c *AtomicCounter) Increment() {
	atomic.AddInt64(&c.value, 1)
}

// Value возвращает текущее значение счетчика атомарно
func (c *AtomicCounter) Value() int64 {
	return atomic.LoadInt64(&c.value)
}

func main() {
	var wg sync.WaitGroup
	counter := AtomicCounter{}

	// Запускаем 1000 горутин для инкрементации счетчика
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
	fmt.Printf("Итоговое значение счетчика: %d\n", counter.Value())
}
