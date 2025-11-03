package main

import "fmt"

func basicLoop(line int, column int) {
	// O(n^2) loop
	// When middle operation is not important
	for range line {
		for j := range column {

			if j%2 == 0 {
				fmt.Printf("!")
			} else {
				fmt.Printf("*")
			}
		}
		fmt.Println()
	}

	// With explicit middle item
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", i)
	}
}
