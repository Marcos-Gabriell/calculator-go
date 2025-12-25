package main

import "fmt"

func main() {

	fmt.Println("Enter 2 numbers to sum:")
	var num1, num2 int
	fmt.Scan(&num1, &num2)

	sum := num1 + num2
	fmt.Printf("%d + %d = %d\n", num1, num2, sum)
}
