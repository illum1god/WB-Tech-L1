package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Примеры переменных разных типов, объявленных как interface{}
	var (
		intVar     interface{} = 1
		stringVar  interface{} = "Hello, World!"
		boolVar    interface{} = true
		channelVar interface{} = make(chan int)
	)

	// Определение типов переменных
	detectType(intVar)
	detectType(stringVar)
	detectType(boolVar)
	detectType(channelVar)
	detectType(3.1415) // Пример неподдерживаемого типа
}

// detectType функция для определения типа переменной типа interface{}
func detectType(v interface{}) {
	// Получение типа переменной с помощью reflect
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Int:
		fmt.Println("Тип переменной: int")
	case reflect.String:
		fmt.Println("Тип переменной: string")
	case reflect.Bool:
		fmt.Println("Тип переменной: bool")
	case reflect.Chan:
		fmt.Println("Тип переменной: channel")
	default:
		fmt.Printf("Тип переменной неопределён или не поддерживается: %s\n", t.Kind())
	}
}
