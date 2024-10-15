package main

import "fmt"

// createHugeString создает строку заданного размера
func createHugeString(size int) string {
	s := ""
	for i := 0; i < size; i++ {
		s += "a"
	}
	return s
}

var justString string

func someFunc() {
	v := createHugeString(1 << 10) // Создаём большую строку
	buf := make([]byte, 100)       // Создаём буфер для 100 байтов
	copy(buf, v[:100])             // Копируем первые 100 байтов
	justString = string(buf)       // Преобразуем буфер в строку
}

func main() {
	someFunc()
	fmt.Println(justString) // Выводим результат
}
