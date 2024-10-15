package main

import "fmt"

// Human представляет базовую структуру с полями и методами.
type Human struct {
	Name string
	Age  int
}

// Greet выводит приветственное сообщение.
func (h *Human) Greet() {
	fmt.Printf("Привет, меня зовут %s и мне %d лет.\n", h.Name, h.Age)
}

// Action встраивает структуру Human, наследуя её поля и методы.
type Action struct {
	Human      // Встраивание структуры Human
	ActionName string
}

// PerformAction выполняет действие и вызывает метод родителя.
func (a *Action) PerformAction() {
	fmt.Printf("Выполняю действие: %s\n", a.ActionName)
	a.Greet() // Вызов метода из Human
}

func main() {
	// Инициализация экземпляра Action с вложенной структурой Human
	a := Action{
		Human: Human{
			Name: "Иван",
			Age:  25,
		},
		ActionName: "Плавание",
	}
	// Выполнение действия и приветствие
	a.PerformAction()
}
