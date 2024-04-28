package main

import "fmt"

func linearSearch(int) int {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	i := 0
	for i = 0; i < len(arr); i++ {
		if arr[i] == 5 {
			return i
		}
	}
	return i
}

func main() {
	linearSearch(5)
	fmt.Println("Index of 5 is: ", linearSearch(5)+1)
}
