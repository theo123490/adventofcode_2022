package fileManipulation

import (
	"log"
	"os"
	"strings"
)

func ReadFileToStringArray(inputFile string) []string {

	content, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	contentString := strings.Split(string(content), "\n")

	return contentString
}

func RotateArrayClockwise(stringArray []string) []string {
	var newStringArray []string = make([]string, len(stringArray[0]))
	for _, row := range stringArray {
		for xIndex, value := range row {
			newStringArray[xIndex] = newStringArray[xIndex] + string(value)
		}
	}

	for yIndex, row := range newStringArray {
		newStringArray[yIndex] = reverse(row)
	}

	return newStringArray
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
