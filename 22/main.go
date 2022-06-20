package main

import (
	"fmt"
)

func main() {

	newBigInt1 := 1 << 21
	newBigInt2 := 1 << 22

	fmt.Println(newBigInt1, "\n", newBigInt2, newBigInt2*newBigInt1)
}
