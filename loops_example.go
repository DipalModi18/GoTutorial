package main

import "fmt"

func main() {
	/*
	The syntax of for loop in Go programming language is
	for [condition |  ( init; condition; increment ) | Range] {
   		statement(s);
	}
	*/

	var b int = 15
	var a int
	numbers := [6]int{1, 2, 3, 5} 
	fmt.Printf("%v", numbers)

	/* for loop execution */
	for a := 0; a < 10; a++ {
		fmt.Printf("value of a: %d\n", a)
	}
	for a < b {
		a++
		fmt.Printf("value of a: %d\n", a)
	}
	for i,x:= range numbers {
		fmt.Printf("value of x = %d at %d\n", x,i)
	}   
}

// To run: go run loops_example.go 