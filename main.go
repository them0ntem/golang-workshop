package main

import "fmt"

func main() {
	msg, err := Hello("World")
	if err == nil {
		fmt.Println(msg)
	}
}
