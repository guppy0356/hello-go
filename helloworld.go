package main

import "fmt"

// Greeting says 'Hello world'.
func Greeting() string {
	msg := "Hello world\n"
	return msg
}

func main() {
	msg := Greeting()
	fmt.Printf(msg)
}
