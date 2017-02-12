package main

import (
	"fmt"

	"github.com/evadnoob/randomizer"
)

func main() {

	wordfile := new(randomizer.WordFile).Exclude("the")
	wordfile.Open("war_and_peace_leo_tolstoy.txt")

	word := wordfile.Random()

	fmt.Printf("%s\n", word)
}
