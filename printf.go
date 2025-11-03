package main

import "fmt"

func printer() {
	fmt.Println("Name size", len(name))
	fmt.Println(name + " has ben concatenated to us!!!")

	fmt.Printf("PI -> %.3f \n", pi)
	fmt.Printf("Type of PI -> %T \n", pi)
	fmt.Printf("Are you alive? %t \n", is_alive)
	fmt.Printf("Binary of 10 is -> %b \n", 10)
	fmt.Printf("Character 33 in ASCII table -> %c \n", 33)
	fmt.Printf("x in hexadecimal -> %x \n", 15)
}
