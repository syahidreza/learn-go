package main

import "fmt"

func main() {
	// ARRAY
	// var int_arr [3]int
	// int_arr[1] = 123
	int_arr := [...]int32{1, 2, 3}
	
	// print array value (returns int)
	fmt.Println(int_arr[0])
	fmt.Println(int_arr[1])
	fmt.Println(int_arr[2])

	// print array 0 to 3 (returns array)
	fmt.Println(int_arr[0:3])

	// for
	for i := 0; i < len(int_arr); i++ {
		fmt.Println(int_arr[i])
	}

	// for range
	for i, el := range int_arr {
		fmt.Printf("idx: %d, el: %d\n", i, el)
	}
}