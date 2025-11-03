package main

var my_pointer_value int = 1

var my_pointer_address *int = &my_pointer_value

func changeValue(x *int, newValue int) {
	*x = newValue
}
