package main

import "fmt"

func main() {
	i := 10
	j := &i
	*j = 100
	fmt.Println(j)
	fmt.Println(i)

	var k int
	k = 10
	var l *int // only declaration no value assigned to l
	l = &k
	*l = 100
	fmt.Println(*l)
	fmt.Println(k)

	//m := new(int)   declaration + assignment

	name := new(string)
	fmt.Println(name) // name = random address
	*name = "golang"
	fmt.Println(*name) // name = golang

}
