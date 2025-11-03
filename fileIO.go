package main

import (
	"fmt"
	"log"
	"os"
)

func handle_file() {
	file, err := os.Create("created-text-file.txt")

	if err != nil {
		log.Fatal(err)
	}

	file.WriteString("This is our recently created text file")
	file.Close()

	stream, err := os.ReadFile("./test-file.txt")

	if err != nil {
		log.Fatal(err)
	}

	string_stream := string(stream)
	fmt.Println(string_stream)
}
