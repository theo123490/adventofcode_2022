package main

import (
	"fmt"

	fileManipulation "theo_adventofcode.com/modules"
)

func main() {
	inputString := fileManipulation.ReadFileToStringArray("input.txt")
	inputString = fileManipulation.RotateArrayClockwise(inputString)
	inputString = extractArray(inputString)
	for _, value := range inputString {
		fmt.Println(value)
	}

	stackA := stack{[]rune("abcd")}
	stackB := stack{[]rune("xyz")}
	stackA.moveItemTo(&stackB)
	stackA.moveItemTo(&stackB)
}

type stack struct {
	items []rune
}

func (currentStack *stack) moveItemTo(destinationStack *stack) {
	destinationStack.items = append(destinationStack.items, currentStack.items[len(currentStack.items)-1])
	currentStack.items = currentStack.items[:len(currentStack.items)-1]
}

func extractArray(stringArray []string) []string {
	var newStringArray []string = make([]string, 0)
	for _, stringLine := range stringArray {
		if stringLine[0] == 32 {
			continue
		}
		newStringArray = append(newStringArray, stringLine[1:])
	}
	return newStringArray
}
