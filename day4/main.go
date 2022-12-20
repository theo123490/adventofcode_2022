package main

import (
	"fmt"
)

func main() {
	ticketA := assignmentTicket{2, 4}
	ticketB := assignmentTicket{1, 5}
	ticketC := assignmentTicket{1, 3}

	var AInB bool = ticketA.isFullyContainedIn(ticketB)
	var AInC bool = ticketA.isFullyContainedIn(ticketC)
	fmt.Printf("ticket a is fully contained in b = %v \n", AInB)
	fmt.Printf("ticket a is fully contained in c = %v \n", AInC)
	fmt.Printf("ticket b is fully contained in c = %v \n", ticketB.isFullyContainedIn(ticketC))
	fmt.Printf("ticket b is fully contained in a = %v \n", ticketB.isFullyContainedIn(ticketA))
}

type assignmentTicket struct {
	StartLocation int
	EndLocation   int
}

func (ticket assignmentTicket) isFullyContainedIn(checkingTicket assignmentTicket) bool {
	return (ticket.StartLocation > checkingTicket.StartLocation) && (ticket.EndLocation < checkingTicket.EndLocation)
}
