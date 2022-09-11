package main

import (
	"fmt"
	"strconv"
)

func main() {
	var index int8 = 15
	var bigIndex int32
	bigIndex = int32(index)

	var big int64 = 100
	var small int8

	var x int64 = 78
	var y float64 = float64(x)

	var z int = 4
	var a float32 = float32(z)

	var str string = "42"
	var num int

	if n, err := strconv.Atoi(str); err == nil {
		num = n
	}

	fmt.Println(bigIndex)
	fmt.Printf("Index variable type: %T\n", index)
	fmt.Printf("BigIndex variable type: %T\n", bigIndex)

	fmt.Printf("Big number %T and Small number %T\n", big, small)
	fmt.Printf("Convert integer to float: %T to %T. New val: %.2f\n", x, y, y)

	fmt.Printf("Conver int to float: %T to %T with new val: %.2f\n", z, a, a)

	fmt.Printf("Convert %T to %T with new val: %d\n", str, num, num)

	val := fmt.Sprint(2333.444)
	fmt.Printf("Convert float to string, with val: %s\n", val)

	fmt.Println("Converting between strings and bytes")
	lotr := "you shall not pass"
	lotrByte := []byte(lotr)
	fmt.Println("My lotrByte", lotrByte)
	fmt.Println("My string ", lotr)

}
