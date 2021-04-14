package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	iCount := 1.0
	dCount := -1.0
	for {
		fmt.Printf("%-30v %064b\n", iCount, math.Float64bits(iCount))
		fmt.Printf("%-30v %064b\n", dCount, math.Float64bits(dCount))
		time.Sleep(1e7)
		iCount *= 2
		dCount /= 2
	}
}
