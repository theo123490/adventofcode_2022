package main

import (
	"fmt"
	"strconv"
	"strings"

	fileManipulation "theo_adventofcode.com/modules"
)

func main() {
	inputString := fileManipulation.ReadFileToStringArray("input.txt")
	inputString = fileManipulation.RotateArrayClockwise(inputString)
	inputString = extractArray(inputString)

	commandInputString := fileManipulation.ReadFileToStringArray("input_command.txt")

	stacks := createStacksFromArray(inputString)
	for _, stackItem := range stacks {
		fmt.Printf("%v \n", stackItem)
	}
	fmt.Println("\n------------------")

	for _, commandString := range commandInputString {
		command := translateCommand(commandString)
		runCommand(&stacks, command)
	}

	for _, stackItem := range stacks {
		fmt.Printf("%v \n", stackItem)
	}

	for _, stack := range stacks {
		fmt.Printf("%c", stack.items[len(stack.items)-1])
	}
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

func createStacksFromArray(stringArray []string) []stack {
	var stacks []stack = make([]stack, 0)

	for _, stackValue := range stringArray {
		for i := len(stackValue) - 1; i >= 0; i-- {
			var stackRemovalCounter int = 0
			if stackValue[i] == 32 {
				stackRemovalCounter++
			}
			stackValue = stackValue[:len(stackValue)-stackRemovalCounter]
		}
		stacks = append(stacks, stack{[]rune(stackValue)})
	}

	return stacks
}

func translateCommand(commandString string) []int {
	commmandStringArray := strings.Split(commandString, " ")
	move, _ := strconv.Atoi(commmandStringArray[1])
	from, _ := strconv.Atoi(commmandStringArray[3])
	to, _ := strconv.Atoi(commmandStringArray[5])
	var commandData []int = []int{
		move,
		from,
		to,
	}

	return commandData
}

func runCommand(stacks_ *[]stack, command []int) {
	stacks := *stacks_
	for i := 0; i < command[0]; i++ {
		stacks[command[1]-1].moveItemTo(&stacks[command[2]-1])
	}
}
