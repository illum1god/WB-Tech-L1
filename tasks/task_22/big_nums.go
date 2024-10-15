package main

import (
	"fmt"
	"math/big"
)

// OperationResult содержит результаты арифметических операций
type OperationResult struct {
	Sum        *big.Int
	Difference *big.Int
	Product    *big.Int
	Quotient   *big.Int
}

// performOperations принимает два числа типа *big.Int и выполняет основные арифметические операции
func performOperations(a, b *big.Int) OperationResult {
	var result OperationResult

	// Сумма: a + b
	result.Sum = new(big.Int).Add(a, b)

	// Разность: a - b
	result.Difference = new(big.Int).Sub(a, b)

	// Произведение: a * b
	result.Product = new(big.Int).Mul(a, b)

	// Частное: a / b (если b != 0)
	if b.Cmp(big.NewInt(0)) != 0 {
		result.Quotient = new(big.Int).Div(a, b)
	} else {
		result.Quotient = nil // Деление на ноль неопределено
	}

	return result
}

func main() {
	// Определяем минимальное значение как 2^20
	minValue := new(big.Int).Exp(big.NewInt(2), big.NewInt(20), nil)

	// Пример значений для a и b (> 2^20)
	a := new(big.Int).Add(minValue, big.NewInt(1)) // 2^20 + 1
	b := new(big.Int).Add(minValue, big.NewInt(2)) // 2^20 + 2

	// Убедимся, что и a, и b больше 2^20
	if a.Cmp(minValue) <= 0 || b.Cmp(minValue) <= 0 {
		fmt.Println("Оба числа a и b должны быть больше 2^20.")
		return
	}

	// Выполняем арифметические операции
	results := performOperations(a, b)

	// Выводим результаты
	fmt.Printf("a + b = %s\n", results.Sum.String())
	fmt.Printf("a - b = %s\n", results.Difference.String())
	fmt.Printf("a * b = %s\n", results.Product.String())
	if results.Quotient != nil {
		fmt.Printf("a / b = %s\n", results.Quotient.String())
	} else {
		fmt.Println("a / b = неопределено (деление на ноль)")
	}
}
