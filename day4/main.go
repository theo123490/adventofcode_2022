package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var contentString []string = readFileToStringArray("input.txt")

	for _, value := range contentString {
		fmt.Println(value)
	}

	ticketA, _ := formatToTicket("2-4")
	ticketB, _ := formatToTicket("1-5")
	ticketC, _ := formatToTicket("1-3")

	var AInB bool = ticketA.isFullyContainedIn(ticketB)
	var AInC bool = ticketA.isFullyContainedIn(ticketC)
	fmt.Printf("ticket a is fully contained in b = %v \n", AInB)
	fmt.Printf("ticket a is fully contained in c = %v \n", AInC)
	fmt.Printf("ticket b is fully contained in c = %v \n", ticketB.isFullyContainedIn(ticketC))
	fmt.Printf("ticket b is fully contained in a = %v \n", ticketB.isFullyContainedIn(ticketA))
	fmt.Printf("ticket 2-4 is fully contained in 3-8 = %v \n", isFormatMatchupEitherContained("2-4,3-8"))
	fmt.Printf("ticket 2-4 is fully contained in 1-8 = %v \n", isFormatMatchupEitherContained("2-4,1-8"))
	fmt.Printf("ticket 4-12 is fully contained in 6-6 = %v \n", isFormatMatchupEitherContained("4-12,6-6"))
}

type assignmentTicket struct {
	StartLocation int
	EndLocation   int
}

func (ticket assignmentTicket) isFullyContainedIn(checkingTicket assignmentTicket) bool {
	return (ticket.StartLocation >= checkingTicket.StartLocation) && (ticket.EndLocation <= checkingTicket.EndLocation)
}

func readFileToStringArray(inputFile string) []string {

	content, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	contentString := strings.Split(string(content), "\n")

	return contentString
}

func formatToTicket(formatString string) (assignmentTicket, error) {
	var valueString []string = strings.Split(formatString, "-")
	startValue, errStart := strconv.Atoi(valueString[0])
	endValue, errEnd := strconv.Atoi(valueString[1])

	if errStart != nil {
		var emptyTicket assignmentTicket
		return emptyTicket, errStart
	}
	if errEnd != nil {
		var emptyTicket assignmentTicket
		return emptyTicket, errEnd
	}

	return assignmentTicket{startValue, endValue}, nil
}

func isFormatMatchupEitherContained(matchupString string) bool {
	var matchupStringArray []string = strings.Split(matchupString, ",")
	matchupTicketA, _ := formatToTicket(matchupStringArray[0])
	matchupTicketB, _ := formatToTicket(matchupStringArray[1])

	isAinB := matchupTicketA.isFullyContainedIn(matchupTicketB)
	isBinA := matchupTicketB.isFullyContainedIn(matchupTicketA)

	return (isAinB || isBinA)
}
