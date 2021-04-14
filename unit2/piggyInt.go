package main

import (
	"fmt"
	"math/rand"
)

func main() {
	total := 0
	for total <= 2000 {
		switch rand.Intn(3) {
			case 0:
				total += 5
			case 1:
				total += 10
			case 2:
				total += 25
		}
		fmt.Printf("$%d.%02d\n", total / 100, total % 100)
	}
}
