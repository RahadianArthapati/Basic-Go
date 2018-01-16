package main

import (
	"fmt"
	"strings"
)

func main() {
	//pointer1()
	pointer2()
	structs1()
}
func pointer1() {
	fruit := "banana"
	giveMePear(fruit)
	fmt.Println(fruit) // Outputs: banana
	giveMeKiwi(&fruit)
	fmt.Println(fruit) // Outputs: kiwi
}
func pointer2() {
	counter := &Counter{}
	fmt.Println(counter.Count) // Outputs: 0

	counter.Increment()
	fmt.Println(counter.Count) // Outputs: 0

	counter.IncrementWithPointer()
	fmt.Println(counter.Count) // Outputs: 1”
}
func structs1() {
	episodeIV := Movie{
		Title:       "Star Wars: A New Hope",
		Rating:      5.0,
		ReleaseYear: 1977,
	}
	episodeIV.Actors = []string{
		"Mark Hamill",
		"Harrison Ford",
		"Carrie Fisher",
	}
	//fmt.Println(episodeIV.Title, "has a rating of", episodeIV.Rating)
	fmt.Println(episodeIV.Title, "has actors such as", episodeIV.Actors)
	fmt.Println(episodeIV.DisplayTitle() + " starring by " + episodeIV.DisplayActors())
}
func giveMePear(fruit string) {
	fruit = "pear"
}
func giveMeKiwi(fruit *string) {
	*fruit = "kiwi"
}

type Movie struct {
	Actors      []string
	Rating      float32
	ReleaseYear int
	Title       string
}

func (movie Movie) DisplayTitle() string {
	return fmt.Sprintf("%s (%d)", movie.Title, movie.ReleaseYear)
}
func (movie Movie) DisplayActors() string {
	s := []string{}
	for _, actor := range movie.Actors {
		s = append(s, actor)
	}
	return strings.Join(s, ", ")
}

type Counter struct {
	Count int
}

func (c Counter) Increment() {
	c.Count++
}

func (c *Counter) IncrementWithPointer() {
	c.Count++
}
