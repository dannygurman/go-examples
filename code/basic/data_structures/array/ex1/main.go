package main

import "fmt"

func main() {
	cards := []string{"ace of spades", newCard()}

	//append appends elements to the end of a slice and return new slice
	cards = append(cards, "six of spades")
	fmt.Println(cards)

	//iterate over array
	//for	_, card := range cards {
	for i, card := range cards {
		fmt.Println(i, card)
	}

}

// string is the return type
func newCard() string {
	return "some new card"
}
