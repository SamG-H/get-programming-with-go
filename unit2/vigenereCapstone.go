package main

import (
	"fmt"
)

func main() {
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	keyWord := "GOLANG"
	keyIndex := 0
	var message string

	// no package no range solution no if
	for i := 0; i < len(cipherText); i++ {
		c := cipherText[i]
		c = (c - keyWord[keyIndex] + 26) % 26 + 'A'
		message += string(c)
		keyIndex++
		keyIndex %= len(keyWord)
	}

	//range no package solution
	/*for _, c := range cipherText {
		newC := byte(c)
		if keyIndex > 5 {
			keyIndex = 0
		}

		newC -= keyWord[keyIndex] % 'A'
		if newC < 'A' {
			newC += 26
		}
		message += string(newC)
		keyIndex++
	}*/

	//package solution
	/*for _, c := range cipherText {
		
	}*/

	fmt.Println(message)
}
