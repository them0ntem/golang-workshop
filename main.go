package main

import "fmt"

func main() {
	msg, err := Hello("World")
	if err == nil {
		fmt.Println(msg)
	}
}

func Hello(name string) (string, error) {
	greetings := fmt.Sprintf("Hello %v!", name)
	return greetings, nil
}
