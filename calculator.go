package main

import "fmt"

func main() {
	resultSum := sum(6, 4)
	fmt.Println("This sum is:", resultSum)

	resultSubtraction := sub(6, 4)
	fmt.Println("This subtration is:", resultSubtraction)

	resultMultiplication := mult(6, 4)
	fmt.Println("This multiplicaiton is:", resultMultiplication)

	resultDivision := div(8, 4)
	fmt.Println("This divison is:", resultDivision)
}

func sum(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func mult(x, y int) int {
	return x * y
}

func div(x, y int) int {
	return x / y
}
