package main

import (
	"fmt"
	"math/big"
)

func main() {
	lightSpeed := big.NewInt(299792)
	secondsPerDay := big.NewInt(60 * 60 * 24)

	distance := new(big.Int)
	distance.SetString("240000000000000000", 10)
	fmt.Println("The Andromeda Galaxy is", distance, "km away.")

	seconds := new(big.Int)
	seconds.Div(distance, lightSpeed)

	days := new(big.Int)
	days.Div(seconds, secondsPerDay)

	fmt.Println("That is", days, "days of travel at the speed of light")
}
