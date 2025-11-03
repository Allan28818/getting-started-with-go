package main

import "fmt"

type Cat interface {
	privateSpeak() // unavailable for another packages
	PublicSpeak()  // public so it's available for another packages
}

type Animal struct {
	specie string // unavailable for another packages
	legs   int    // unavailable for another packages
	age    int    // unavailable for another packages
}

func (a Animal) privateSpeak() {
	fmt.Println("Meoooooooooooooooooooooooow")
}

func (a Animal) PublicSpeak() {
	fmt.Println("Meow from " + a.specie)
}

func interfaceGo() {
	animal := Animal{specie: "cat", legs: 4, age: 2}

	animal.privateSpeak()
	var cat Cat = animal

	cat.PublicSpeak()
}
