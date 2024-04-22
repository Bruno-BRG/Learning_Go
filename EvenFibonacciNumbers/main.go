package main

import "fmt"

func fibonacci() int {
	a, b := 0, 1
	sum := 0
	for {
		if a > 4000000 {
			break
		}
		if a%2 == 0 {
			fmt.Println(a)
			sum += a
		}
		a, b = b, a+b
	}
	return sum
}

func main() {
	evenSum := fibonacci()
	fmt.Println("Sum of even Fibonacci numbers", evenSum)
}
