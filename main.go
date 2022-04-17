package main

import (
	"fmt"
	"golang_workshop/facade"
)

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev, revErr := facade.Reverse(input)
	doubleRev, doubleRevErr := facade.Reverse(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q, err: %v\n", rev, revErr)
	fmt.Printf("reversed again: %q, err: %v\n", doubleRev, doubleRevErr)
}
