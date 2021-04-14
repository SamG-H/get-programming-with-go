package main

import (
	"fmt"
	"math"
)

func main() {
	var age int
	fmt.Printf("%T %064[1]b\n", age)
	fmt.Printf("%T %064[1]b\n", math.MaxInt64)
	fmt.Printf("%T %064[1]X\n", math.MinInt64)
}
