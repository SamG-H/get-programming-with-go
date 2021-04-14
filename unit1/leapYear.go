package main

import (
	"fmt"
	"math/rand"
)

func main() {
	year := rand.Intn(4000) + 1
	if (year % 400 == 0 || (year % 4 == 0 && year % 100 != 0)) {
		fmt.Printf("%v IS a leap year!\n", year)
	} else {
		fmt.Printf("%v is NOT a leap year!\n", year)
	}
}
