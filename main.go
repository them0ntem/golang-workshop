package main

import (
	"errors"
	"fmt"
)

func main() {
	msg, err := Hello("World")
	if err == nil {
		fmt.Println(msg)
	}
}

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("supplied name is empty")
	}
	greetings := fmt.Sprintf("Hello %v!", name)
	return greetings, nil
}
