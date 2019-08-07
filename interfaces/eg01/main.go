package main

import (
	"fmt"
	_ "fmt"
)

// VowelsFinder định nghĩa 1 interface
type VowelsFinder interface {
	FindVowels() []rune
}

// MyString alias type string
type MyString string

// FindVowels MyString implements VowelsFinder
func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

func main() {
	name := MyString("Sam Anderson")
	var v VowelsFinder
	v = name // Điều này được chấp nhận vì MyString implement interface VowelsFinder
	fmt.Printf("Vowels are %c", v.FindVowels())
}
