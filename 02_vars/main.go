package main

import "fmt"

func main() {

	// INTEGER:
	// var int_num int16 = 32767 // the result is -32768
	var int_num int = 32767
	int_num = int_num + 1
	fmt.Println(int_num)

	var int_num_1 int8 = 3
	var int_num_2 int8 = 2
	fmt.Println(int_num_1 / int_num_2) // result is 1
	fmt.Println(int_num_1 % int_num_2) // result is 1

	// FLOAT:
	var float_num float64 = 12345678.9
	fmt.Println(float_num)

	var float_num_1 float32 = 3.0
	var float_num_2 float32 = 2.0
	fmt.Println(float_num_1 / float_num_2)
	// fmt.Println(float_num_1 % float_num_2) // can't do this on float

	var float_num_32 float32 = 10.1
	var int_num_32 int32 = 2
	// var result float32 = float_num_32 + int_num_32 // can't do this
	var result float32 = float_num_32 + float32(int_num_32)
	fmt.Println(result)

	// STRING:
	var my_string_1 string = "Hi, mom!"
	var my_string_2 string = `Hi, 
	mom!`
	var my_string_3 string = "Hi," + " " + "mom!"
	fmt.Println(my_string_1)
	fmt.Println(my_string_2)
	fmt.Println(my_string_3)
	fmt.Println(len(my_string_1))

	// BOOLEAN:
	var my_boolean bool = false // true / false
	fmt.Println(my_boolean)
	
	// OTHER WAY TO DECLARE VAR
	my_string := "Hi, mom!"
	fmt.Println(my_string)

	my_int := 45
	fmt.Println(my_int)

	my_float := 325.423
	fmt.Println(my_float)
	
	// CONSTANT
	const my_const string = "Hi, mom!"
	// my_const = "Hi, dad!" // can't do this
	fmt.Println(my_const)
}