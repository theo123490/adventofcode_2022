package main

import (
	"fmt"

	fileManipulation "theo_adventofcode.com/modules"
)

func main() {
	inputString := fileManipulation.ReadFileToStringArray("input.txt")
	for _, value := range inputString {
		fmt.Println(value)
	}
	inputString = fileManipulation.RotateArrayClockwise(inputString)
	for _, value := range inputString {
		fmt.Println(value)
	}
	stackA := stack{[]rune("abcd")}
	stackB := stack{[]rune("xyz")}
	stackA.moveItemTo(&stackB)
	stackA.moveItemTo(&stackB)
	for _, item := range stackA.items {
		fmt.Printf("%c", item)
	}
	fmt.Println("\n----")
	for _, item := range stackB.items {
		fmt.Printf("%c", item)
	}
}

type stack struct {
	items []rune
}

func (currentStack *stack) moveItemTo(destinationStack *stack) {
	destinationStack.items = append(destinationStack.items, currentStack.items[len(currentStack.items)-1])
	currentStack.items = currentStack.items[:len(currentStack.items)-1]
}
