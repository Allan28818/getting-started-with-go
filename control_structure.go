package main

import "fmt"

func controlStructures() {
	defer FirstRun()
	defer SecondRun()
	defer ThirdRun()

	divisionInt(10, 0)
}

func FirstRun() {
	fmt.Println("First")
}
func SecondRun() {
	fmt.Println("Second")
}
func ThirdRun() {
	fmt.Println("Third")
}
