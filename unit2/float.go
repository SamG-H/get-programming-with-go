package main

import (
	"fmt"
)

func main() {
	var money float64 = 987.65
	fmt.Println(money)
	const third float64 = 1.0 / 3
	fmt.Printf("%5.5f\n", third)
	fmt.Printf("%.9f\n", third)
	fmt.Println(third + third + third)
	piggyBank := 0.1
	piggyBank += 0.2
	fmt.Println("Println: 0.1 + 0.2 =", piggyBank)
	fmt.Printf("Printf:  0.1 + 0.2 = %.2f\n", piggyBank)
	dime := .10
	for count := 0; count < 10; count++ {
		dime += .10
	}
	fmt.Println(dime)
}
