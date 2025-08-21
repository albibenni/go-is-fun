package main

import "fmt"

type Example struct {
	Hello string
	SomeN int
}

type ExampleWithEmbedded struct {
	Example
	Another string
}

type ExampleWithNested struct {
	Example Example
	Another string
}

func main() {
	myExample := Example{
		Hello: "Hello albibenni",
		SomeN: 10,
	}
	fmt.Printf("-------\n")
	fmt.Printf("-------%s\n", myExample.Hello)
	fmt.Printf("-------%d\n", myExample.SomeN)
	fmt.Printf("-------%v\n", myExample)

	fmt.Printf("-------\n")
	anonymous := struct {
		Name string
		Age  int
	}{
		Name: "albibenni",
		Age:  10,
	}

	fmt.Printf("-------%v\n", anonymous)
	fmt.Printf("-------\n")
	exampleWithEmbedded := ExampleWithEmbedded{
		Example: myExample,
		Another: "some",
	}

	exampleWithNested := ExampleWithEmbedded{
		Example: myExample,
		Another: "SOme",
	}
	fmt.Printf("-------%v\n", exampleWithEmbedded)
	fmt.Printf("accessing with Example prop-------%v\n", exampleWithEmbedded.Example)
	fmt.Printf("accessing Example prop directly without Example field -------%v\n", exampleWithEmbedded.Hello)
	fmt.Printf("-------%v\n", exampleWithNested)

}
