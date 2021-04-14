package main

import (
	"fmt"
	"math/big"
)

func main() {
	const secondsInYear = 60 * 60 * 24 * 365
	const lightSpeedKmPerSec = 299792
	lightYearInKm := big.NewInt(lightSpeedKmPerSec * secondsInYear)
	fmt.Println("Number of km in 1 light year:", lightYearInKm)

	distanceToCanisMajorInKm := big.NewInt(2.36e17)
	distanceToCanisMajorInLightYear := new(big.Int)
	distanceToCanisMajorInLightYear.Div(distanceToCanisMajorInKm, lightYearInKm)

	fmt.Println("Number of lights years required to travel to Canis Major Dwarf:", distanceToCanisMajorInLightYear)
}
