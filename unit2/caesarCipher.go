package main

import "fmt"

func main() {
	const message = "L fdph, L vdz, L frqtxhuhg."
	const msg = "LFDPHLVDZLFRQTXHUHG"
	for i := 0; i < len(msg); i++ {
		c := msg[i]
		if c >= 'A' && c <= 'Z' {
			c = c - 3
			if c < 'A' {
				c += 26
			}
		}
		if c >= 'a' && c <= 'z' {
			c = c - 3
			if c < 'a' {
				c += 26
			}
		}
		fmt.Printf("%c", c)
	}
	fmt.Println()
}
