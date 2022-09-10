package main

import "fmt"

var g = "global"

const (
	year     = 30         // untyped
	leapYear = int64(366) // typed, will only accept int64
)

func main() {
	//i := 1023124123
	//fmt.Println(i)
	//
	//x := 71 + 89
	//fmt.Println(x)
	//
	//slice := []string{"hi", "bye"}
	//fmt.Println(slice)

	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("var a %T = %+v\n", a, a)
	fmt.Printf("var b %T = %q\n", b, b)
	fmt.Printf("var c %T = %+v\n", c, c)
	fmt.Printf("var d %T = %+v\n\n", d, d)

	//var Email string
	//var password string

	// Reassigning Variables
	i := 77
	fmt.Println(i)

	i = 79
	fmt.Println(i)

	// Global and local variables
	l := "local"
	fmt.Println(l)

	const shark = "sammy"
	fmt.Println(shark)

}
