package main

import (
	"fmt"
	"strconv"
)

func BiggestPairNumber(numbers int) int {

	numberString := strconv.Itoa(numbers)

	var pair = 0
	var biggest = 0

	for i := 1; i < len(numberString); i++ {
		num1, _ := strconv.Atoi(string(numberString[i]))
		num2, _ := strconv.Atoi(string(numberString[i-1]))

		if num1+num2 > biggest {
			biggest = num1 + num2

			pairNumber, _ := strconv.Atoi(string(numberString[i-1]) + string(numberString[i]))
			pair = pairNumber
		}
	}

	return pair // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(BiggestPairNumber(11223344))
	fmt.Println(BiggestPairNumber(89083278))
}
