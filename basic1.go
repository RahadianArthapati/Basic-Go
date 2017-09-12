package main

import (
	"fmt"
)

func main() {
	add := func(m int) {
		return m + 1
	}
	multiply := func(m int, n int) {
		return m * n
	}
	result := add(5)
	//looping1()
	//mapping1()
	fmt.Println(result)
}
func looping1() {
	animals := []string{"Cat", "Dog", "Emu", "Warthog"}
	for i, animal := range animals {
		fmt.Println(animal, "is at index", i)
	}
	for _, animal := range animals {
		fmt.Println(animal)
	}
	for index := 0; index < len(animals); index++ {
		value := animals[index]
		fmt.Println(fmt.Sprintf("Index is %d and value is %s", index, value))
	}
}

func mapping1() {
	starWarsYears := map[string]int{
		"A New Hope":              1977,
		"The Empire Strikes Back": 1980,
		"Return of the Jedi":      1983,
		"Attack of the Clones":    2002,
		"Revenge of the Sith":     2005,
	}
	starWarsYears["The Force Awakens"] = 2015
	//fmt.Println(len(starWarsYears))
	for title, year := range starWarsYears {
		fmt.Println(title, "was released in", year)
	}
	starWarsYears["The Last Jedi"] = 2018
	code, exists := starWarsYears["The Last Jedi"]
	if exists {
		fmt.Println("The Last Jedi was released in", code)
	} else {
		fmt.Println("I don't know The Last Jedi")
	}
}
