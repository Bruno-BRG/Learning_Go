package main

import (
	"errors"
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	index, err := binarySearch(numbers, 3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The value is at position: ", index+1)
	}
}

func binarySearch(numbers []int, item int) (int, error) {
	low := 0
	high := len(numbers) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := numbers[mid]

		if guess == item {
			return mid, nil
		}

		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1, errors.New("item not found in array")
}
