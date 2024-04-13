// i want a code that prints out the numbers of fibonacci up until the 1000th number

package main

import (
	"fmt"
)

// i want it to only print the last number

func main() {
	var a, b, c, count int
	a = 0
	b = 1
	count = 0
	for count < 1000 {
		c = a + b
		a = b
		b = c
		count++
	}
	fmt.Println(c)
}
