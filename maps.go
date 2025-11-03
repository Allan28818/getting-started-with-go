package main

import "fmt"

func Map() {
	simple_map := map[string]float64{
		"longitude": 10.1999,
		"latitude":  11.1234234,
	}

	fmt.Println(simple_map)

	car := make(map[string]map[string]string)

	car["Wheel"] = map[string]string{
		"Size": "10 inches",
	}

	fmt.Println(car["Wheel"]["Size"])

	if value, was_found := car["Wheel"]["Size"]; was_found {
		fmt.Println(value, was_found)
	}
}
