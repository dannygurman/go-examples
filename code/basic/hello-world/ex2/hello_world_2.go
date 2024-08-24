package main

import (
	"fmt"
)

func main() {
	myString := "gopher"
	//fmt.Println("Hello and welcome, %s!", myString)
	fmt.Printf("Hello and welcome, %s!\n", myString)

	for i := 1; i <= 5; i++ {
		fmt.Println("i =", 100/i)
	}
}
