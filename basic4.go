package main

import (
	"fmt"
	"sort"
)

func main() {
	map_example2()

}
func map_example1() {
	langs := map[string]string{
		"EL": "Greek",
		"EN": "English",
		"ES": "Spanish",
		"FR": "French",
		"HI": "Hindi",
	}
	if lan, ok := langs["EN"]; ok {
		fmt.Println(lan)
	}
	langs["ES"] = "Estonia"
	lan1 := langs["ES"]
	fmt.Println(lan1)

}
func map_example2() {
	chapts := make(map[int]string)
	// Add data as key/value pairs
	chapts[1] = "Beginning Go"
	chapts[2] = "Go Fundamentals"
	chapts[3] = "Structs and Interfaces"
	// Slice for specifying the order of the map
	var keys []int
	// Appending keys of the map
	for k := range chapts {
		keys = append(keys, k)
	}
	// Ints sorts a slice of ints in increasing order.
	sort.Ints(keys)
	// Iterate over the map with an order
	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", chapts[k])
	}
}
