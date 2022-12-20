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
	var containedCounter int = 0
	var overlappedCounter int = 0
	for _, formatedMatchup := range contentString {
		if isFormatMatchupEitherContained(formatedMatchup) {
			containedCounter++
		}
		if isFormatMatchupOveralapped(formatedMatchup) {
			overlappedCounter++
		}
	}
	fmt.Println(containedCounter)
	fmt.Println(overlappedCounter)
}

type assignmentTicket struct {
	StartLocation int
	EndLocation   int
}

func (ticket assignmentTicket) isFullyContainedIn(checkingTicket assignmentTicket) bool {
	return (ticket.StartLocation >= checkingTicket.StartLocation) && (ticket.EndLocation <= checkingTicket.EndLocation)
}

func (ticket assignmentTicket) isOverlapped(checkingTicket assignmentTicket) bool {
	var isStartIncluded bool = checkingTicket.StartLocation <= ticket.StartLocation && ticket.StartLocation <= checkingTicket.EndLocation
	var isEndIncluded bool = checkingTicket.StartLocation <= ticket.EndLocation && ticket.EndLocation <= checkingTicket.EndLocation
	var isCheckingStartIncluded bool = ticket.StartLocation <= checkingTicket.StartLocation && checkingTicket.StartLocation <= ticket.EndLocation
	var isCheckingEndIncluded bool = ticket.StartLocation <= checkingTicket.EndLocation && checkingTicket.EndLocation <= ticket.EndLocation

	return isStartIncluded || isEndIncluded || isCheckingEndIncluded || isCheckingStartIncluded
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

func isFormatMatchupOveralapped(matchupString string) bool {
	var matchupStringArray []string = strings.Split(matchupString, ",")
	matchupTicketA, _ := formatToTicket(matchupStringArray[0])
	matchupTicketB, _ := formatToTicket(matchupStringArray[1])

	isAinB := matchupTicketA.isOverlapped(matchupTicketB)
	isBinA := matchupTicketB.isOverlapped(matchupTicketA)

	return (isAinB || isBinA)
}
