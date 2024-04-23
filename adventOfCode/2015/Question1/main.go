package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	str := string(content)
	floorCount := 0 // Initialize floorCount outside the loop

	for i := 1; i < len(str); i++ {
		if str[i] == ')' {
			floorCount--

			if floorCount < 0 {
				fmt.Println("Final floor level:", str[i])
				break
			}
		} else if str[i] == '(' {
			floorCount++
		}
	}

	
}
