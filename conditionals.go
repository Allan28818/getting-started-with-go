package main

import "fmt"

func conditionals(age int) {
	switch age {
	case 1:
		fmt.Println("You're a baby")
	case 10:
		fmt.Println("You're a child")
	case 15:
		fmt.Println("Teenager 🤮")
	case 20:
		fmt.Println("You're cool now")
	case 30:
		fmt.Println("Adult")
	case 50:
		fmt.Println("Hey uncle bob!")
	case 100:
		fmt.Println("Are you sure that you're alive?")
	default:
		fmt.Println("Not inside")
	}
}
