package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println(name, age, weight)
	fmt.Println(addInt(10, 20))
	if is_alive {
		fmt.Println("That's good that you're not dead!!!")
	}
	res := divisionInt(10, 2)
	fmt.Println(res)
	fmt.Printf("Ponteiro de %d -> %p\n", my_pointer_value, my_pointer_address)
	changeValue(&my_pointer_value, 10)
	fmt.Printf("Ponteiro de %d -> %p\n", my_pointer_value, my_pointer_address)
	printer()
	basicLoop(10, 30)
	conditionals(100)
	arrays()
	Map()
	fmt.Println(factorial(3))
	controlStructures()
	goStruct()
	interfaceGo()

	handle_file()

	http.HandleFunc("/heart-beat", handler)

	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello API!!!")
}
