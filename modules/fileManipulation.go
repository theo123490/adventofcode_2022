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
