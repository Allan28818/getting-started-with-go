package main

import "fmt"

func arrays() {
	var my_arr [10]int

	my_arr[0] = 10
	my_arr[1] = 20
	my_arr[2] = 30
	my_arr[3] = 40
	my_arr[4] = 50
	my_arr[5] = 60
	my_arr[6] = 70
	my_arr[7] = 80
	my_arr[8] = 90
	my_arr[9] = 100

	for i := range len(my_arr) {
		fmt.Println(i, my_arr[i])
	}
	for index, value := range my_arr {
		fmt.Println(index, value)
	}

	secondary_array := [5]int{1, 2, 3, 4, 5}

	fmt.Println(secondary_array)
	fmt.Println(secondary_array[3:5])

	sliced := secondary_array[2:]
	sliced2 := secondary_array[:2]
	fmt.Println(sliced2, sliced)

	copy(sliced, sliced2)
	fmt.Println(sliced)
}
