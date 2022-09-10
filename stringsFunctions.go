package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "Hi my name is Eric"
	b := "times times times times"
	c := "times"

	fmt.Println(strings.ToUpper(a))
	fmt.Println(strings.ToLower(strings.ToUpper(a)))
	fmt.Println(strings.Contains(a, "z"))
	fmt.Println(strings.HasPrefix(a, "H"))
	fmt.Println(strings.Count(b, "times"))
	fmt.Println(len(c))

}
