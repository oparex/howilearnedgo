package main

import (
		"fmt"
		"bufio"
		"os"
		"strings"
		"strconv"
)

func sum(i int) int {
	sum := 0
	for j := 1; j <= i; j++ {
		sum += j
	}
	return sum
}

func prod(i int) int {
	sum := 1
	for j := 1; j <= i; j++ {
		sum *= j
	}
	return sum
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter number: ")
	nr, _ := reader.ReadString('\n')
	nr = strings.Replace(nr, "\n", "", 1)
	i, err := strconv.Atoi(nr)
	if err != nil {
		fmt.Println(nr, "is not a number")
		return
	}
	fmt.Print("Sum(s) or product(p)? (s): ")
	option, _ := reader.ReadString('\n')
	option = strings.Replace(option, "\n", "", 1)
	if option == "s" || option == "" {
		j := sum(i)
		fmt.Println("Sum of", i, "is", j)
	} else if option == "p" {
		j := prod(i)
		fmt.Println("Product of", i, "is", j)
	} else {
		fmt.Println("Your have chosen an invalid option")
	}
}