package main

import "fmt"

/*
	for(condition){
		statements
	}
*/
func main() {
	// Example of nested loop
	// for i := 0; i < 3; i++ {
	// 	for i := 0; i < 3; i++ {
	// 		fmt.Println(i)
	// 	}
	// }
	// loop with interator of boolean type
	// for true {
	// 	fmt.Println("Infinite Execution")
	// }

	i := 0
	for i < 3 {
		fmt.Println(i)
		i++
	}
	// i := 0
	for {
		fmt.Println(i)
		i++
		if i >= 3 {
			break
		}
	}

	num := []int{1, 2, 3, 4}
	for _, value := range num {
		fmt.Println(value)
	}

	name := []string{"Shatvik", "Pandey"}
	for _, val := range name {
		for _, value := range val {
			if value == 'e' {
				break
			}
			fmt.Print(string(value))
		}
		fmt.Println()
	}

	for i := range make([]int, 3) {
		if i == 1 {
			continue
		}
		fmt.Println(i)
	}

	// list:=range 2
}
