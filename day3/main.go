package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	var content_str string

	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	content_str = fmt.Sprintf("%s", content)

	fmt.Println(content_str)
}
