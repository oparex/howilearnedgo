package main

import (
		"fmt"
		"bufio"
		"os"
		"strings"
		"strconv"
		"math"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter number: ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", 1)
	f, err := strconv.ParseFloat(text, 64)
	if err != nil {
		fmt.Println("Can't convert", text, "to float64")
		return
	}
	sum := 0.0
	for j := 1.0; j <= f; j++ {
		if math.Mod(j, 3.0) == 0.0 || math.Mod(j, 5.0) == 0.0 {
			sum += j
		}
	}
	fmt.Println(sum)
}