package main

import "fmt"

func main() {
	// number := 10
	var number int
	number = 10
	fmt.Println(number)

	//store a function as a value to a variable
	// var variable_name datatype
	var getInt func(int) int
	getInt = func(x int) int {
		fmt.Println("In a 1st function")
		return 20 + x
	}
	g(getInt, 8)
	// getInt = func(x int) int {
	// 	fmt.Println("In a 2nd function")
	// 	return 10 + x
	// }
	var y func()
	y = func() {
		fmt.Println(9)
	}
	y()

	j := func(x int) int {
		//nameless function or anonymous function
		fmt.Println("In a 2st function")
		return 20 + x
	}(10)

	// declaring and calling function same time
	fmt.Println(j)
}

//closure
func g(getInt func(int) int, u int) {
	j := getInt(78)
	fmt.Print(j)
}

//function is a first class member/variable in golang: we can use a function as a type
