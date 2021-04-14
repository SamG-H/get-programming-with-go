package main

import (
	"fmt"
)

func main() {
	msg := `hello I am a string literal.\n`
	fmt.Printf("Yeah\n%T\n", msg)
	fmt.Printf("%c\n", msg[1])
	//strings are immutable
	//nono 
	//msg[1] = 'a'

	const smileyFace rune = 128515
	for count := smileyFace - 500; count < 129500; count++ {
		fmt.Printf("%c", count)
	}

	const bangRune rune = '!'
	const bangByte byte = '!'
	//prints numeric value of the bang char
	fmt.Printf("\n%v %[1]T\n", bangRune)

	fmt.Println(rune(bangByte) == bangRune)

	var ascii byte = 0

	for ascii < 128 {
		fmt.Printf("%c :: %[1]v\n", ascii)
		ascii++
	}
}
