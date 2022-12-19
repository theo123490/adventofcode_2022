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
	var rucksack_item int32
	var items_map map[int32]int = make(map[int32]int)
	for _, string_value := range content_string {
		fmt.Println("----")
		compartments = []string{
			string_value[:len(string_value)/2],
			string_value[len(string_value)/2:],
		}

		for _, value_compartment1 := range compartments[0] {
			for _, value_compartment2 := range compartments[1] {
				if value_compartment1 == value_compartment2 {
					rucksack_item = value_compartment1
				}
			}
		}

		items_map[rucksack_item]++
		fmt.Println(items_map)
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
