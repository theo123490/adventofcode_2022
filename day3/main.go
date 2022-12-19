package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var content_string = readfile_to_string_array("input.txt")
	var compartments []string
	for _, string_value := range content_string {
		fmt.Println("----")
		compartments = []string{
			string_value[:len(string_value)/2],
			string_value[len(string_value)/2:],
		}
		fmt.Println(compartments)
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
