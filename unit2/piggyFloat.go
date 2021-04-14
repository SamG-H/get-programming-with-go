package main

import (
	"fmt"
	"math/rand"
)

func main() {
	piggyBank := 0.0
	count := 0
	for piggyBank <= 20 {
		switch rand.Intn(3) {
			case 0:
				piggyBank += 0.05
			case 1:
				piggyBank += 0.10
			case 2:
				piggyBank += 0.25
		}
		fmt.Printf("$%4.2f\n", piggyBank)
		count++
	}
}
