package main

import (
	"fmt"
	"os"
)

func main() {
	//множественное присваивание
	i, j := 100, 20
	fmt.Fprintln(os.Stdout, i, j)
	i, j = j, i
	fmt.Fprintln(os.Stdout, i, j)
}
