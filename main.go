package main

import "fmt"

func main() {
	var text string
	fmt.Print("Enter your text: ")
	fmt.Scan(&text)
	fmt.Printf("%x \n", text)
}
