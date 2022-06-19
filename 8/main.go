package main

import "fmt"

var in int64 = 1000

func main() {
	fmt.Printf("%b\n", in)
	offset(in, 1, 1)

}

func offset(num int64, bit, bitPos int) {
	num = num << 1
	fmt.Printf("%b", num)
}
