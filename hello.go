package main

import "fmt"

func Hello(name string) (string, error) {
	greetings := fmt.Sprintf("Hello %v!", name)
	return greetings, nil
}
