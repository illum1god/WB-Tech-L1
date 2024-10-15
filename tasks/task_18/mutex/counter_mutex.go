package main

import (
	"fmt"
	"sync"
)

// Counter структура-счетчик с мьютексом для обеспечения потокобезопасности
type Counter struct {
	mu    sync.Mutex
	value int
}

// Increment увеличивает значение счетчика на 1
func (c *Counter) Increment() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

// Value возвращает текущее значение счетчика
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	var wg sync.WaitGroup
	counter := Counter{}

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
