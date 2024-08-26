package main

import "fmt"

func main() {
	card := newCard()
	fmt.Println(card)
}

// string is the return type
func newCard() string {
	return "card"
}
