package main

import (
	"fmt"
)

func main() {
	user := map[string]string{"firstName": "Eric", "lastName": "Chiu", "city": "Boston"}
	//
	//fmt.Println(user)
	//fmt.Println(user["firstName"])
	//
	//// Iterate over a map with Range
	//for k, v := range user {
	//	fmt.Printf("key %q : value %q, ", k, v)
	//}

	// See a list of keys or values
	var keys []string
	for key := range user {
		keys = append(keys, key)
	}
	//sort.Strings(keys)
	fmt.Printf("\n %q \n", keys)

	// Allocating slices efficiently
	keysEfficient := make([]string, len(keys))
	for ke := range user {
		keysEfficient = append(keysEfficient, ke)
	}

	// Checking existence
	counts := map[string]int{"first": 1, "second": 2, "third": 3}
	//count, ok := counts["fourth"]
	//if ok {
	//	fmt.Printf("Day count is %d\n", count)
	//} else {
	//	fmt.Println("Day count not found")
	//}

	if count, ok := counts["fourth"]; ok {
		fmt.Printf("Day ount of %d\n", count)
	} else {
		fmt.Println("Day count Fourth was not found")
	}

	// Modifying Maps
	// Changing key values
	user["firstName"] = "Winnie"
	fmt.Println(user["firstName"])

	// Adding new keys
	user["favoriteColor"] = "blue"
	fmt.Println(user)
	fmt.Println(user["favoriteColor"])

	// Remove items from maps
	delete(user, "favoriteColor")
	fmt.Println(user)

	permissions := map[int]string{1: "read", 2: "write", 4: "delete", 8: "create", 16: "modify"}
	delete(permissions, 16)
	fmt.Println(permissions)

	// Remove all items in a map
	permissions = map[int]string{}

}
