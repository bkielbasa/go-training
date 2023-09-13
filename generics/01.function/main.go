package main

import "fmt"

func printNumbers(numbers []int) {
	fmt.Print("Numbers: ")
	for _, num := range numbers {
		fmt.Print(num, " ")
	}
	fmt.Print("\n")
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	printNumbers(numbers)
}
