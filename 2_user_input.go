package main

import (
		"fmt"
		"bufio"
		"os"
		"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", 1)
	if text == "Alice" || text == "Bob" {
		fmt.Println("Hello", text)
	}
}