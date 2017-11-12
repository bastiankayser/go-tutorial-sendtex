package main

import "fmt"

func foo() {
	defer fmt.Println("Done")
	defer fmt.Println("Are we done?")
	fmt.Println("Doing so stuff, who knos what?")
}

func main() {
	foo()
}