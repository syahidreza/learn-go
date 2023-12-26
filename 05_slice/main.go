package main

import "fmt"

func main()  {
	// like array, but dont have to init size of the element 
	var arr_cars = [4]string{"brio", "ayla", "agya", "swift"} // array
	var slice_cars = []string{"brio", "ayla", "agya", "swift"} // slice

	fmt.Printf("arr_cars len: %d\n", len(arr_cars)) // length of current elements
	fmt.Printf("arr_cars cap: %d\n", cap(arr_cars)) // maximum capacity 

	fmt.Printf("slice_cars len: %d\n", len(slice_cars)) // length of current elements
	fmt.Printf("slice_cars cap: %d\n", cap(slice_cars)) // maximum capacity 

	// slice from array
	var arr_bikes = [4]string{"beat", "vario", "scoopy", "genio"}
	var slice_bikes = arr_bikes[0:3]

	fmt.Printf("arr_bikes len: %d\n", len(arr_bikes)) // length of current elements
	fmt.Printf("arr_bikes cap: %d\n", cap(arr_bikes)) // maximum capacity 

	fmt.Printf("slice_bikes len: %d\n", len(slice_bikes)) // length of current elements
	fmt.Printf("slice_bikes cap: %d\n", cap(slice_bikes)) // maximum capacity (4, because it's a reference of arr_bikes)

	// append slice
	var new_slice_bikes = append(slice_bikes, "filano")
	fmt.Println(new_slice_bikes) // [beat vario scoopy filano]

	// the array also changed
	fmt.Println(arr_bikes) // [beat vario scoopy filano]
}