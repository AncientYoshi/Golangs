package main

import "fmt"

func concat(s1 string, s2 string) string {
	return s1 + s2
}

func main() {
	fmt.Println(concat("hello", "world"))
	fmt.Println(concat("Mingalrpr, ","Myanmar"))
	fmt.Println(concat("Bingalr, ","San Francisco"))
	

}