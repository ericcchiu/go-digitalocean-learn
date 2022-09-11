package main

import "fmt"

func main() {
	x := 5
	y := 6

	isEven := 4%2 == 0 && 2%2 == 0

	fmt.Println("x == y:", x == y)
	fmt.Println("x != y:", x != y)
	fmt.Println("x < y:", x < y)
	fmt.Println("x > y:", x > y)
	fmt.Println("x <= y:", x <= y)
	fmt.Println("x >= y:", x >= y)
	fmt.Println("Number is even: ", isEven)
}
