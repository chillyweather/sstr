package main

import (
	"fmt"
	"strings"

	"github.com/atotto/clipboard"
)

func main() {
	text, err := clipboard.ReadAll()
	if err != nil {
		fmt.Println("Error reading from clipboard:", err)
		return
	}

	result := strings.Replace(text, "\n", "", -1)

	err = clipboard.WriteAll(result)

	if err != nil {
		fmt.Println("Error copying to clipboard:", err)
	} else {
		fmt.Println("Text copied to clipboard!")
	}
}
