package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var content_string = readfile_to_string_array("input.txt")
	for _, string_value := range content_string {
		fmt.Println(string_value)
	}
}

func readfile_to_string_array(input_file string) []string {

	content, err := os.ReadFile(input_file)
	if err != nil {
		log.Fatal(err)
	}

	content_str := strings.Split(string(content), "\n")

	return content_str
}
