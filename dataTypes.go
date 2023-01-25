package main

import "fmt"

// main function would be automatically called when you start the application
func main() {
	// var Variable_Name datatype
	var data int
	data = 8

	// another way of declaring and initialising variable
	data1 := 100 // declare(automatic) the variable and assign value to it

	// data := 1000 //redeclaring the variable

	data = 1000 //reassigning the data variable
	// data1 = false // wrong assignment value type
	Data := 10
	fmt.Println(data)
	fmt.Println(data1)
	fmt.Println(Data)

	// DataType
	// 1. Primitive
	// int, float64, string, bool, complex
	// 2. Non-Primitive
	// struct, map, chan, interface, array, slice
	// functional programming

	name := "Shatvik"
	fmt.Println(name)
}
