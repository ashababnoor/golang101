package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Go proverb:", quote.Go())
	fmt.Println("Optimization truth:", quote.Opt())
	fmt.Println("For world travelers:", quote.Glass())
}
