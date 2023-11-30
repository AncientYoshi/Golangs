package main

import "fmt"

func main() {
	const SecondinMinute = 60
	const MinuteInHour = 60
	const SecondinHour = MinuteInHour * SecondinMinute

	fmt.Println("Second In Hour:",SecondinHour)
}