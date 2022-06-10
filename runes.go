package main

import (
	"fmt"
	"unicode/utf8"
)

/*
	A go string is a read-only slice of bytes
	- containers of text encoded in UTF-8
	A character is called a rune => integer that represents a Unicode code point
*/
func main() {
	//thai word that says hello
	// go string literals are UTF-8 encoded text
	const s = "สวัสดี"

	//strings are arrays of bytes, so len will show the length of bytes
	fmt.Println("Len:", len(s)) // 18 bytes

	//returns the byte value at every index
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	//range will decode the rune to its unicode value
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	//doing the samething above but with the utf8 method
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

//you can compare rune literals (chars) directly
func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
