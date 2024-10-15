package main

import (
	"fmt"
	"log"
)

// setBit устанавливает i-й бит числа num в значение bitVal (0 или 1).
func setBit(num int64, i uint, bitVal int) int64 {
	if i > 63 {
		log.Fatalf("Недопустимый индекс бита: %d. Допустимые значения от 0 до 63.\n", i)
	}

	switch bitVal {
	case 1:
		// Устанавливаем i-й бит в 1 с помощью побитовой операции OR.
		num |= (1 << i)
	case 0:
		// Устанавливаем i-й бит в 0 с помощью побитовой операции AND с инвертированным битовым маской.
		num &^= (1 << i)
	default:
		log.Fatalf("Недопустимое значение бита: %d. Допустимые значения: 0 или 1.\n", bitVal)
	}

	return num
}

// getBit возвращает значение i-го бита числа num (0 или 1).
func getBit(num int64, i uint) int {
	if i > 63 {
		log.Fatalf("Недопустимый индекс бита: %d. Допустимые значения от 0 до 63.\n", i)
	}

	// Сдвигаем num вправо на i позиций и применяем побитовое AND с 1, чтобы получить значение i-го бита.
	return int((num >> i) & 1)
}

func main() {
	var num int64 = 0  // Исходное число.
	var i uint = 30    // Индекс бита, который хотим изменить (начиная с 0).
	var bitVal int = 1 // Значение бита: 1 для установки в 1, 0 для установки в 0.

	fmt.Printf("Исходное число: %d (двоичное: %064b)\n", num, num)

	// Устанавливаем i-й бит в bitVal.
	num = setBit(num, i, bitVal)
	fmt.Printf("После установки бита %d в %d: %d (двоичное: %064b)\n", i, bitVal, num, num)

	// Проверяем значение i-го бита.
	currentBit := getBit(num, i)
	fmt.Printf("Значение бита %d сейчас: %d\n", i, currentBit)

	// Устанавливаем тот же бит в 0.
	bitVal = 0
	num = setBit(num, i, bitVal)
	fmt.Printf("После установки бита %d в %d: %d (двоичное: %064b)\n", i, bitVal, num, num)

	// Проверяем значение i-го бита снова.
	currentBit = getBit(num, i)
	fmt.Printf("Значение бита %d сейчас: %d\n", i, currentBit)
}
