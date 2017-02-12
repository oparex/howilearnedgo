package main

import (
		"fmt"
		"bufio"
		"os"
		"strings"
		"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter number: ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", 1)
	i, err := strconv.Atoi(text)
	if err != nil {
		fmt.Println(text, "is not a number")
		return
	}
	sum := 0
	for j := 1; j <= i; j++ {
		sum += j
	}
	fmt.Println(sum)
}