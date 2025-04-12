package main

// MultiplyByTwoAndThree takes a pointer to a float32 number and returns
// its double and triple values. Returns 0, 0 if the pointer is nil.
func MultiplyByTwoAndThree(num *float32) (float32, float32) {
	if num == nil {
		return 0, 0
	}
	return *num * 2, *num * 3
}

// GenerateMultiples takes a base number and returns its first three multiples
// starting from 1x up to 3x.
func GenerateMultiple() (int, int, int) {

	return 1, 2, 3
}
