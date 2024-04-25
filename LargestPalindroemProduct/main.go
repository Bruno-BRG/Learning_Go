package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	rev := 0
	original := x
	for x != 0 {
		rev = rev*10 + x%10
		x = x / 10
	}
	return original == rev
}

func findHighestPalindrome() int {
	highest := 0
	for i := 999; i > 99; i--{
		for j := 999; j > 99; j--{
			if isPalindrome(i*j) && i*j > highest {
				highest = i*j
			}
		}
	}
}
