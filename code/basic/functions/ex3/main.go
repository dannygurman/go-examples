package main

import "fmt"

func main() {
	card := newCard()
	fmt.Println(card)

	cards := []string{newCard(), newCard()}
	fmt.Println(cards)
}

// string is the return type
func newCard() string {
	return "card"
}
