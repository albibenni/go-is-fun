package main

import "fmt"

type Animal interface {
	Speak() string
	Move() string
}

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

func (d Dog) Move() string {
	return "Running"
}

func (c Cat) Speak() string {
	return "Meow!"
}

func (c Cat) Move() string {
	return "Stalking"
}

func MakeAnimalAct(a Animal) {
	fmt.Printf("The animal says: %s\n", a.Speak())
	fmt.Printf("The animal is: %s\n", a.Move())
}

func main() {
	dog := Dog{Name: "Buddy"}
	cat := Cat{Name: "Whiskers"}

	MakeAnimalAct(dog) // Works - Dog implements Animal
	MakeAnimalAct(cat) // Works - Cat implements Animal
}
