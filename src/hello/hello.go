package main

import (
	"fmt"
	"rsc.io/quote/v4"
	"example.com/greetings"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())

	message := greetings.Hello("Glados")
	fmt.Println(message)
}
