package main

import "fmt"

func main() {
	// Пример слайса
	slice := []int{1, 2, 3, 4, 5}
	i := 2 // индекс элемента, который нужно удалить

	// Удаление i-го элемента
	slice = append(slice[:i], slice[i+1:]...)

	// Вывод результата
	fmt.Println(slice) // [1 2 4 5]
}
