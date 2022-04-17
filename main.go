package main

import (
	"fmt"
	"golang_workshop/facade"
)

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev := facade.Reverse(input)
	doubleRev := facade.Reverse(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q\n", rev)
	fmt.Printf("reversed again: %q\n", doubleRev)
}
