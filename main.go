package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a string")
		return
	}

	input := os.Args[1]

	fmt.Printf("You entered %v string\n", input)
}
