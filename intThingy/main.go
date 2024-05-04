package main

import "fmt"

func numberParser(x int) {
	for i := 0; i < x; i++ {
		fmt.Println(i)
	}
}

func main() {
	numberParser(10)
}
