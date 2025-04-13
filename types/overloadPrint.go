package main

import "fmt"

type Course struct {
	Name string
	Type string
	Cost int
}
type Coursable interface {
	PrintIt() string
}

func (c Course) String() string {
	return fmt.Sprintf("Name: %s, Type: %s, Cost: %d \n", c.Name, c.Type, c.Cost)
}

func (c Course) PrintIt() string {
	return fmt.Sprintf("Name: %s, Type: %s, Cost: %d \n", c.Name, c.Type, c.Cost)
}

type Workshop struct {
	Course
}

type WorkshopO struct {
	Course Course
}

func printThem() {
	//var wPrintable Workshop
	wNonPrintable := WorkshopO{
		Course: Course{
			Name: "test",
			Type: "my type",
			Cost: 10,
		},
	}
	wPrintable := Workshop{
		Course: Course{
			Name: "test",
			Type: "my type",
			Cost: 10,
		},
	}


	fmt.Printf("%v \n", wNonPrintable)
	fmt.Printf("%v \n", wPrintable)
}
