package main

import (
	"errors"
	"fmt"
)

func main() {
	const printValue string = "Hi, mom!"
	echo(printValue)

	sum := add(3, 7)
	echo(sum)


	numerator := 11
	denominator := 9
	result, remain, error := div(numerator, denominator)

	switch {
		case error != nil:
			fmt.Println(error.Error())
		case remain == 0:
			fmt.Printf("%v / %v = %v \n", numerator, denominator, result)
		default:
			fmt.Printf("%v / %v = %v \n", numerator, denominator, result)
			fmt.Printf("Remain: %v \n", remain)
	}
}

// PROCEDURE
func echo(printValue any) {
	fmt.Println(printValue)
}

// FUNCTION
func add(first int, second int) int {
	return first + second
}

// MULTIPLE RESULTS WITH ERROR 
func div(numerator int, denominator int) (int, int, error) {
	var err error

	if denominator == 0 {
		err = errors.New("can't divide by 0")
		return 0, 0, err
	}	

	var result int = numerator / denominator
	var remain int = numerator % denominator

	return result, remain, err
}