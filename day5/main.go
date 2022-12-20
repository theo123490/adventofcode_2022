package main

import "fmt"

func main() {
	stackA := stack{[]rune("abcd")}
	fmt.Printf("%c", stackA.items[0])
}

type stack struct {
	items []rune
}
