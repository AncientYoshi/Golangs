package main

import "fmt"

func main() {
	messagefromDenny := []string{
		"Mana have limitations",
		"Do You Love Me",
		"Hello Hi Kya nw NAme ka Hlwn Paing",
		"Pleased To Meet you",
	}
	numMessages := float64(len(messagefromDenny))
	costperMessage := .02

	var totalCost = costperMessage * numMessages

	fmt.Printf("Denny spent %.2f on text message todays.\n", totalCost)
}