package main

import "fmt"

func addInt(a int, b int) int {
	return a + b
}

func subInt(a int, b int) int {
	return a - b
}

func divisionInt(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("You cannot divide by %f", b)
	}

	return a / b, nil
}

func multInt(a int, b int) int {
	return a * b
}
