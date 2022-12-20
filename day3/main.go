package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var items_map map[int32]int = make(map[int32]int)
	var priority_sum int = 0
	var content_string []string = readfile_to_string_array("input.txt")

	items_map = sort_item_map(content_string)
	priority_sum = calculate_priority_sum(items_map)
	fmt.Printf("the sum of item priority is = %v \n", priority_sum)
}

func sort_item_map(content_string []string) map[int32]int {
	var items_map map[int32]int = make(map[int32]int)
	var compartments []string
	var rucksack_item int32

	for _, string_value := range content_string {
		compartments = []string{
			string_value[:len(string_value)/2],
			string_value[len(string_value)/2:],
		}

		rucksack_item = get_rucksack_item(compartments)
		items_map[rucksack_item]++
	}

	return items_map
}

func calculate_priority_sum(items_map map[int32]int) int {
	var priority string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var priority_sum int = 0
	var item_value int

	for item_key, item_amount := range items_map {
		for item_value = 1; item_value <= 52; item_value++ {
			if priority[item_value-1] == byte(item_key) {
				priority_sum += int(item_amount) * int(item_value)
			}
		}
	}

	return priority_sum
}

func get_rucksack_item(rucksacks []string) int32 {
	var validator_flag bool
	var is_item_available bool
	for _, value_checked_on := range rucksacks[0] {
		is_item_available = false
		validator_flag = true

		for rucksack_index, rucksack := range rucksacks {
			if rucksack_index == 0 {
				continue
			}
			is_item_available = check_if_item_available(value_checked_on, rucksack)
			validator_flag = is_item_available && validator_flag
		}
		if validator_flag {
			return int32(value_checked_on)
		}
	}

	return 0
}

func check_if_item_available(item int32, compartment string) bool {
	var compartment_item int32
	for _, compartment_item = range compartment {
		if compartment_item == item {
			return true
		}
	}
	return false
}

func readfile_to_string_array(input_file string) []string {

	content, err := os.ReadFile(input_file)
	if err != nil {
		log.Fatal(err)
	}

	content_str := strings.Split(string(content), "\n")

	return content_str
}
