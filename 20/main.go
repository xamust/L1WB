package main

import (
	"fmt"
	"strings"
)

var (
	input_first = "now dog sun"
	input_sec   = "Разработать программу, которая переворачивает слова в строке"
)

func main() {
	fmt.Printf("Input %s, revert: %s\n", input_first, revert(input_first))
	fmt.Printf("Input %s, revert: %s\n", input_sec, revert(input_sec))
}

func revert(input string) string {
	temp := strings.Split(input, " ")
	for i, j := 0, len(temp)-1; i < len(temp)/2; i, j = i+1, j-1 {
		temp[i], temp[j] = temp[j], temp[i]
	}
	return strings.Join(temp, " ")
}
