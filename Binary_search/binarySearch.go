package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("o valor esta na posicao: ", binarySearch(numbers, 3) + 1)
}

func binarySearch(numbers []int, target int) int {
	low := 0
	high := len(numbers) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := numbers[mid]

		if guess == target {
			return mid
		}

		if guess > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}
