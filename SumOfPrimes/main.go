package main

import (
   "fmt"
   "math"
)

func main(){
	var sum int
	for i := 0; i <= 2000000; i++{
		if checkPrimeNumber(i){
			sum += checkPrimeNumber()
		}
	}
	fmt.Println("Here is the total sum: %i", sum)
}

func checkPrimeNumber(num int) {
	sum1 := num
	if num < 2 {
      fmt.Println("Number must be greater than 2.")
   }
   sq_root := int(math.Sqrt(float64(num)))
   for i:=2; i<=sq_root; i++{
      if num % i == 0 {
         fmt.Println("Non Prime Number")
      }
   }
   fmt.Println("Prime Number")
   return 
}
