// main.go
package main

import (
	"fmt"
)

// Target определяет интерфейс, который клиенты ожидают.
type Target interface {
	Request() string
}

// Adaptee содержит существующий интерфейс, который необходимо адаптировать.
type Adaptee struct{}

// SpecificRequest — это специальный метод, который несовместим с Target.
func (a *Adaptee) SpecificRequest() string {
	return "Adaptee выполняет специфичный запрос."
}

// Adapter адаптирует Adaptee к интерфейсу Target.
type Adapter struct {
	Adaptee *Adaptee
}

// Request реализует интерфейс Target, обращаясь к специфическому методу Adaptee.
func (a *Adapter) Request() string {
	return a.Adaptee.SpecificRequest()
}

func main() {
	adaptee := &Adaptee{}
	adapter := &Adapter{Adaptee: adaptee}

	fmt.Println("Клиент: Использование интерфейса Target.")
	fmt.Println(adapter.Request())
}
