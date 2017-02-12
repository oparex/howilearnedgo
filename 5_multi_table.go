package main

import (
		"fmt"
)

func main() {
	var lines [12][12]int
	for i := 1; i < 13; i++ {
		for j := 1; j < 13; j++ {
			if i - 1 == 0 {
				lines[i-1][j-1] = j
			} else if j - 1 == 0 {
				lines[i-1][j-1] = i
			} else {
				lines[i-1][j-1] = i*j
			}
		}
	}
	for i := range lines {
		fmt.Println(lines[i])
	}
}