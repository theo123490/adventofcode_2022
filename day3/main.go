package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var content_string = readfile_to_string("input.txt")
	fmt.Println(content_string)
}

func readfile_to_string(input_file string) string {

	content, err := os.ReadFile(input_file)
	if err != nil {
		log.Fatal(err)
	}

	content_str := string(content)

	return content_str
}
