package main

import (
	"fmt"
)

func main() {
	ticketA := assignmentTicket{2, 4}
	ticketB := assignmentTicket{1, 5}
	fmt.Println(ticketA.StartLocation)
	fmt.Println(ticketB.EndLocation)
}

type assignmentTicket struct {
	StartLocation int
	EndLocation   int
}
