package main

import (
	"fmt"
	"strings"
)

// todo: implement strings.Repeat solution
const (
	PrintColor = "\033[38;5;%dm%v\033[39;49m\n"
)

func wrap(num int) {
	for j := 0; j < 256; j++ {
		fmt.Printf(PrintColor, j, j)
	}
}

func decipherText(cipherText string, keyWord string) string {
	keyIndex := 0
	var message string

	//no package no range solution no if
	for i := 0; i < len(cipherText); i++ {
		c := cipherText[i]
		c = (c-keyWord[keyIndex]+26) % 26 + 'A'
		message += string(c)
		keyIndex++
		keyIndex %= len(keyWord)
	}

	return message
}

func cipherText(msg string, keyWord string) string {
	formattedMsg := strings.ToUpper(strings.Replace(msg, " ", "", -1))

	keyIndex := 0
	var cipherText string

	//range no package solution
	for _, c := range formattedMsg {
		if keyIndex >= len(keyWord) {
			keyIndex = 0
		}
		newC := byte(c)
		newC += keyWord[keyIndex] % 'A'
		if newC > 'Z' {
			newC -= 26
		}
		cipherText += string(newC)
		keyIndex++
	}

	return cipherText
}

func main() {
	wrap(0)
	encryptedText := "CSOITEUIWUIZNSROCNKFD"
	keyWord := "GOLANG"
	fmt.Println(decipherText(encryptedText, keyWord))
	fmt.Println(encryptedText)
	fmt.Println(cipherText(decipherText(encryptedText, keyWord), keyWord))

	// check if they match
	if encryptedText == cipherText(decipherText(encryptedText, keyWord), keyWord) {
		fmt.Println("It works!")
	} else {
		fmt.Println("It didn't work...")
	}

	msg := "I want to rock and roll all day and party every night"
	encryptedText = cipherText(msg, keyWord)
	fmt.Println(encryptedText)
	fmt.Println(decipherText(encryptedText, keyWord))
}
