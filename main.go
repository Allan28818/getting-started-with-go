package main

import "fmt"

func main() {

	fmt.Println(name, age, weight)
	fmt.Println(addInt(10, 20))
	if is_alive {
		fmt.Println("That's good that you're not dead!!!")
	}
	res, _ := divisionInt(10, 2)
	fmt.Println(res)
	fmt.Printf("Ponteiro de %d -> %p\n", my_pointer_value, my_pointer_address)
	changeValue(&my_pointer_value, 10)
	fmt.Printf("Ponteiro de %d -> %p\n", my_pointer_value, my_pointer_address)
	printer()
	basicLoop(10, 30)
}
