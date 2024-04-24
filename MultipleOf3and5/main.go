package main
import "fmt"
func isDivisibleByThreeOrFive(number int) bool {
	return number%3 == 0 || number%5 == 0
}
func generateData() []int {
	var data []int
	for i := 1; i < 1000; i++ {
		if isDivisibleByThreeOrFive(i) {
			data = append(data, i)
		}
	}
	return data
}
func printData(data []int) int {
	var sum int
	for _, number := range data {
		sum += number
		fmt.Println(number)
	}
	return sum
}
func main() {
	data := generateData()
	totalSum := printData(data)
	fmt.Println("here is the total sum:", totalSum)
}