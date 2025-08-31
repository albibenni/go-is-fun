package main

import (
	"cmp"
	"fmt"
)

// 1. Basic Generic Function
func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// 2. Generic Struct - A simple container
type Container[T any] struct {
	Value T
}

func (c *Container[T]) Set(value T) {
	c.Value = value
}

func (c *Container[T]) Get() T {
	return c.Value
}

// 3. Generic Struct - A stack data structure
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}

	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item, true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// 4. Generic Interface - Define behavior for types
type Printable[T any] interface {
	Print() string
}

// A struct that implements the generic interface
type Person struct {
	Name string
	Age  int
}

func (p Person) Print() string {
	return fmt.Sprintf("Person: %s, Age: %d", p.Name, p.Age)
}

// Another struct implementing the same interface
type Product struct {
	Name  string
	Price float64
}

func (p Product) Print() string {
	return fmt.Sprintf("Product: %s, Price: $%.2f", p.Name, p.Price)
}

// 5. Function that works with the generic interface
func PrintItem[T Printable[T]](item T) {
	fmt.Println(item.Print())
}

// 6. Generic interface with type constraints
type Numeric interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

type Calculator[T Numeric] struct {
	value T
}

func (c *Calculator[T]) Add(x T) T {
	c.value += x
	return c.value
}

func (c *Calculator[T]) Multiply(x T) T {
	c.value *= x
	return c.value
}

func (c *Calculator[T]) GetValue() T {
	return c.value
}

// 7. Generic function with multiple type parameters
func Pair[T, U any](first T, second U) struct {
	First  T
	Second U
} {
	return struct {
		First  T
		Second U
	}{First: first, Second: second}
}

// 8. Generic slice operations
func Map[T, U any](slice []T, fn func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

func Filter[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	fmt.Println("=== Go Generics Examples ===\n")

	// 1. Basic generic function
	fmt.Println("1. Basic Generic Function:")
	fmt.Printf("Max(10, 20) = %d\n", Max(10, 20))
	fmt.Printf("Max(3.14, 2.71) = %.2f\n", Max(3.14, 2.71))
	fmt.Printf("Max(\"apple\", \"banana\") = %s\n\n", Max("apple", "banana"))

	// 2. Generic container
	fmt.Println("2. Generic Container:")
	intContainer := Container[int]{Value: 42}
	fmt.Printf("Int container: %d\n", intContainer.Get())

	stringContainer := Container[string]{}
	stringContainer.Set("Hello, Generics!")
	fmt.Printf("String container: %s\n\n", stringContainer.Get())

	// 3. Generic stack
	fmt.Println("3. Generic Stack:")
	stack := Stack[string]{}
	stack.Push("first")
	stack.Push("second")
	stack.Push("third")

	fmt.Print("Popping items: ")
	for !stack.IsEmpty() {
		if item, ok := stack.Pop(); ok {
			fmt.Printf("%s ", item)
		}
	}
	fmt.Println("\n")

	// 4. & 5. Generic interface
	fmt.Println("4. & 5. Generic Interface:")
	person := Person{Name: "Alice", Age: 30}
	product := Product{Name: "Laptop", Price: 999.99}

	PrintItem(person)
	PrintItem(product)
	fmt.Println()

	// 6. Generic struct with constraints
	fmt.Println("6. Generic Calculator with Constraints:")
	calc := Calculator[float64]{value: 10.0}
	fmt.Printf("Initial value: %.2f\n", calc.GetValue())
	fmt.Printf("After adding 5.5: %.2f\n", calc.Add(5.5))
	fmt.Printf("After multiplying by 2: %.2f\n", calc.Multiply(2))
	fmt.Println()

	// 7. Multiple type parameters
	fmt.Println("7. Multiple Type Parameters:")
	pair1 := Pair("name", 25)
	pair2 := Pair(true, 3.14159)
	fmt.Printf("String-Int pair: %+v\n", pair1)
	fmt.Printf("Bool-Float pair: %+v\n\n", pair2)

	// 8. Generic slice operations
	fmt.Println("8. Generic Slice Operations:")
	numbers := []int{1, 2, 3, 4, 5}

	// Map: square each number
	squared := Map(numbers, func(x int) int { return x * x })
	fmt.Printf("Original: %v\n", numbers)
	fmt.Printf("Squared: %v\n", squared)

	// Map: convert to strings
	strings := Map(numbers, func(x int) string { return fmt.Sprintf("num-%d", x) })
	fmt.Printf("As strings: %v\n", strings)

	// Filter: get even numbers
	evens := Filter(numbers, func(x int) bool { return x%2 == 0 })
	fmt.Printf("Even numbers: %v\n", evens)
}
