package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	data := "username=ekred password=123 email=eric123@gmail.com"
	s := "times time timed time times times"

	fmt.Println(strings.Join([]string{"World", "is", "round"}, ","))
	fmt.Printf("%q", strings.Split("Eric has a laptop", " "))
	fmt.Println("")
	fmt.Printf("%q", strings.Fields(data))
	fmt.Println("")
	fmt.Println(strings.ReplaceAll(s, "time", "changed"))

	fmt.Println("Convert numbers to string" + strconv.Itoa(1000))

	num, err := strconv.ParseFloat("123.1", 64)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(num)

	fmt.Println(strconv.ParseFloat("3.14", 32))

}
