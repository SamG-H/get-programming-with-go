package main

import "fmt"

func encodeRot13(msg string) string {
	var result string
	for i := 0; i < len(msg); i++ {
		c := msg[i]
		if c >= 'a' && c <= 'z' {
			c = c - 13
			if c > 'z' {
				c = c - 26
			}
			if c < 'a' {
				c = c + 26
			}
		}
		result += string(c)
	}
	return result
}

func decodeRot13(msg string) string {
	var result string
	for i := 0; i < len(msg); i++ {
		c := msg[i]
		if c >= 'a' && c <= 'z' {
			c = c + 13
			if c > 'z' {
				c = c - 26
			}
			if c < 'a' {
				c = c + 26
			}
		}
		result += string(c)
	}
	return result
}

func main() {
	encodedMessage := "uv vagreangvbany fcnpr fgngvba"
	fmt.Println(decodeRot13(encodedMessage))
	fmt.Println(encodeRot13("hi international space station"))
}
