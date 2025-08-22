package main

import "fmt"

// Using errors.New()

// Custom error types
type ValidationError struct {
	Field string
	Value string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("invalid value '%s' for field '%s'", e.Value, e.Field)
}

func main() {
	err := ValidationError{
		Field: "email",
		Value: "invalid-email",
	}

	fmt.Println(err)
}
