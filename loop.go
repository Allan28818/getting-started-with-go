package main

import "fmt"

func basicLoop(line int, column int) {
	// O(n^2) loop
	// When index is not important
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

	// When we need to know the index
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", i)
	}
}
