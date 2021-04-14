package main

import "fmt"

func main() {
	msg := "Hola Estación Espacial Internacional"

	for _, c := range msg {
		if c >= 'a' && c <= 'z' {
			c += 13
			if c > 'z' {
				c -= 26
			}
		} else if c >= 'A' && c <= 'Z' {
			c += 13
			if c > 'Z' {
				c -= 26
			}
		}
		fmt.Printf("%c", c)
	}
	fmt.Println()
}
