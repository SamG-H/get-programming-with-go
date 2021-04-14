package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const distanceToMars = 62100000
	const secondsInADay = 60 * 60 * 24

	fmt.Printf("%-16v %4v %10v %4v\n", "Spaceline", "Days", "Trip Type", "Price")
	fmt.Println("======================================")

	for count := 0; count < 10; count++ {
		speed := rand.Intn(15) + 16

		tripLength := distanceToMars / (speed * secondsInADay)

		tripType := "One-way"
		if rand.Intn(2) == 0 {
			tripType = "Round-trip"
		}

		spaceline := "SpaceX"
		switch rand.Intn(3) {
		case 0:
			spaceline = "Virgin Galactic"
		case 1:
			spaceline = "Space Adventures"
		}

		price := speed + 20
		if tripType == "Round-trip" {
			price *= 2
		}
		fmt.Printf("%-16v %4v %10v $%4v\n", spaceline, tripLength, tripType, price)
	}
}
