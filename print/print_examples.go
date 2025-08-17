package main

import "fmt"

func main() {
	// String formats
	str := "Hello"
	fmt.Printf("%%s: %s\n", str) // %s: Hello
	fmt.Printf("%%q: %q\n", str) // %q: "Hello"
	fmt.Printf("%%x: %x\n", str) // %x: 48656c6c6f

	// Integer formats
	num := 42
	fmt.Printf("%%d: %d\n", num) // %d: 42
	fmt.Printf("%%o: %o\n", num) // %o: 52
	fmt.Printf("%%x: %x\n", num) // %x: 2a
	fmt.Printf("%%X: %X\n", num) // %X: 2A
	fmt.Printf("%%b: %b\n", num) // %b: 101010
	fmt.Printf("%%c: %c\n", num) // %c: *

	// Float formats
	pi := 3.14159265359
	fmt.Printf("%%f: %f\n", pi)     // %f: 3.141593
	fmt.Printf("%%.2f: %.2f\n", pi) // %.2f: 3.14
	fmt.Printf("%%e: %e\n", pi)     // %e: 3.141593e+00
	fmt.Printf("%%g: %g\n", pi)     // %g: 3.14159

	// Boolean format
	isTrue := true
	fmt.Printf("%%t: %t\n", isTrue) // %t: true

	// Pointer format
	ptr := &num
	fmt.Printf("%%p: %p\n", ptr) // %p: 0xc000014078

	// General formats
	fmt.Printf("%%v: %v\n", num) // %v: 42
	fmt.Printf("%%T: %T\n", num) // %T: int
	fmt.Printf("%%%%: %%\n")     // %%: %
}
