package main

import "fmt"

func main() {
	// recursion
	rec(4)
	fmt.Println(fib(10))
	fmt.Println(fact(5))

}

/*
func rec(num int){
	if num==0{
		return
	}
	if num%2==0{
		fmt.Println(num+1)
		rec(num)
	}else{
		fmt.Println(num+2)
		rec(num)
	}
}
*/

func rec(num int) {
	if num <= 0 {
		return
	}
	rec(num - 1)
	rec(num - 2)
	fmt.Println(num - 1)
}

// fibonacci series
func fib(number int) int {
	if number == 0 || number == 1 {
		return number
	}
	result := fib(number-1) + fib(number-2)
	return result
}

// factorial of a number
func fact(number int) int {
	if number == 1 || number == 0 {
		return 1
	}

	if number < 0 {
		fmt.Println("Invalid Number")
		return -1
	}
	result := number * fact(number-1)
	return result
}
