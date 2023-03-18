package main

import (
	"fmt"
)

func main() {
	a := 5
	b := 8
	fmt.Printf("%b %d %d\n", ^(4),6 & ^(1 << 1),^(4))

	a ^= b 
	fmt.Printf("%b %b \n", a, b)

	b ^= a 
	fmt.Printf("%b %b \n", a, b)

	a ^= b 
	fmt.Printf("%b %b \n", a, b)
}

//0100 -> 1011 