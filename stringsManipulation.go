package main

import (
	"fmt"
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

}
