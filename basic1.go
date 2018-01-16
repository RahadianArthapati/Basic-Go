package main

import (
	"fmt"
)

func main() {

	//looping1()
	//mapping1()
	/*
		a, b := 5, 8
		fn := func(sum int) (int, int) {
			x := sum * a / b
			y := sum - x
			return x, y
		}
		// Passing function value as an argument to another function
		SplitValues(fn)
		// Calling the function value by providing argument
		x, y := fn(20)
		fmt.Println(x, y)*/
	//array1()
	slice1()
}

func SplitValues(f func(sum int) (int, int)) {
	x, y := f(35)
	fmt.Println(x, y)
	x, y = f(50)
	fmt.Println(x, y)
}
func slice1() {
	x := make([]int, 2, 5)
	x[0] = 10
	x[1] = 20
	fmt.Println("Slice x:", x)
	fmt.Printf("Length is %d Capacity is %d\n", len(x), cap(x))
	// Create a bigger slice
	x = append(x, 30, 40, 50)
	fmt.Println("Slice x after appending data:", x)
	fmt.Printf("Length is %d Capacity is %d\n", len(x), cap(x))
	x = append(x, 60, 70, 80)
	fmt.Println("Slice x after appending data for the second time:", x)
	fmt.Printf("Length is %d Capacity is %d\n", len(x), cap(x))
}
func array1() {
	langs := [4]string{"Go", "Rust", "Scala", "Julia"}
	for i := 0; i < len(langs); i++ {
		fmt.Printf("langs[%d]:%s \n", i, langs[i])
	}

	for k, v := range langs {
		fmt.Printf("langs[%d]:%s \n", k, v)
	}
	for _, v := range langs {
		fmt.Printf(v)
	}

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
