package main
import (
    "fmt"
)
func isPrime(n int) bool {
    if n <= 1 {
        return false
    }
    if n <= 3 {
        return true
    }
    if n%2 == 0 || n%3 == 0 {
        return false
    }
    i := 5
    for i*i <= n {
        if n%i == 0 || n%(i+2) == 0 {
            return false
        }
        i += 6
    }
    return true
}
func sumPrimesBelowLimit(limit int) int {
    sum := 0
    for i := 2; i < limit; i++ {
        if isPrime(i) {
            sum += i
        }
    }
    return sum
}
func main() {
    limit := 2000000
    sum := sumPrimesBelowLimit(limit)
    fmt.Println("The sum of all primes below two million is:", sum)
}

print("The sum of all primes below two million is:", sumPrimesBelowLimit(2000000))
