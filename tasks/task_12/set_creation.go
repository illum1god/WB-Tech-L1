package main

import "fmt"

// Set представляет собой множество строк
type Set struct {
	elements map[string]struct{}
}

// NewSet создает новое множество
func NewSet() *Set {
	return &Set{
		elements: make(map[string]struct{}),
	}
}

// Add добавляет элемент в множество
func (s *Set) Add(item string) {
	s.elements[item] = struct{}{}
}

// Contains проверяет, содержит ли множество данный элемент
func (s *Set) Contains(item string) bool {
	_, exists := s.elements[item]
	return exists
}

// ToSlice возвращает элементы множества в виде среза строк
func (s *Set) ToSlice() []string {
	keys := make([]string, 0, len(s.elements))
	for key := range s.elements {
		keys = append(keys, key)
	}
	return keys
}

func main() {
	// Исходная последовательность строк
	strings := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создание нового множества
	set := NewSet()

	// Добавление элементов в множество
	for _, s := range strings {
		set.Add(s)
	}

	// Вывод элементов множества
	fmt.Println("Элементы множества:", set.ToSlice())
}
