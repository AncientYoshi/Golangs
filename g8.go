package main

import "fmt"

func main() {
	maxMessageSize := 20
	lenMessageSize := 30
	fmt.Println("Trying to Send the Message Size with", maxMessageSize,"Length=", lenMessageSize)

	if lenMessageSize <= maxMessageSize {
		fmt.Println("Message Sent successfully")
	}else {
		fmt.Println("Message Not Sent")
	}
}