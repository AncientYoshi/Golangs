package main

import "fmt"

func main() {
	const nama = "namastay"
	const openrate = 24.5

	msg := fmt.Sprintf("Hi %v Your openrate is %.1f Nice to Go", nama, openrate)

	fmt.Println(msg)
}