package main

import "fmt"

func addInt(a int, b int) int {
	return a + b
}

func subInt(a int, b int) int {
	return a - b
}

func divisionInt(a int, b int) int {
	defer func() {
		fmt.Println("error", recover())
	}()

	if b == 0 {
		panic("Your cannot divide by zero")
	}

	return a / b
}

func multInt(a int, b int) int {
	return a * b
}
