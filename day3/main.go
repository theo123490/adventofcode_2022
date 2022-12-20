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
	var priority string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var priority_sum int = 0
	var item_value int

	for _, string_value := range content_string {
		compartments = []string{
			string_value[:len(string_value)/2],
			string_value[len(string_value)/2:],
		}

		rucksack_item = get_rucksack_item(compartments)
		items_map[rucksack_item]++
	}

	for item_key, item_amount := range items_map {
		for item_value = 1; item_value <= 52; item_value++ {
			if priority[item_value-1] == byte(item_key) {
				priority_sum += int(item_amount) * int(item_value)
			}
		}
	}
	fmt.Printf("the sum of item priority is = %v \n", priority_sum)
}

func get_rucksack_item(compartments []string) int32 {
	for _, value_compartment1 := range compartments[0] {
		for _, value_compartment2 := range compartments[1] {
			if value_compartment1 == value_compartment2 {
				return int32(value_compartment1)
			}
		}
	}

	return 0
}

func readfile_to_string_array(input_file string) []string {

	content, err := os.ReadFile(input_file)
	if err != nil {
		log.Fatal(err)
	}

	content_str := strings.Split(string(content), "\n")

	return content_str
}
