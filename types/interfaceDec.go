package main

import "fmt"

//	type Course struct {
//		Name string
//		Type string
//		Cost int
//	}
//
//	type Coursable interface {
//		PrintIt() string
//	}
//
//	func (c Course) String() string {
//		return fmt.Sprintf("Name: %s, Type: %s, Cost: %d \n", c.Name, c.Type, c.Cost)
//	}
//
//	func (c Course) PrintIt() string {
//		return fmt.Sprintf("Name: %s, Type: %s, Cost: %d \n", c.Name, c.Type, c.Cost)
//	}
//
//	type Workshop struct {
//		Course
//	}
//
//	type WorkshopO struct {
//		Course Course
//	}
func interfaceAndUsageWithTypes() {

	var course Coursable
	course = Course{
		Name: "using interface name",
		Type: "type interface",
		Cost: 9000,
	}
	fmt.Printf("%v", course.PrintIt())
}
