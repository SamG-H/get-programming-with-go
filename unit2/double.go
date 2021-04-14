package main

import (
	"fmt"
)

func main() {
	var count uint64 = 1
	fmt.Printf("%-20v %-16v %-64v\n", "Decimal", "Hexadecimal", "Binary")
	for i := 0; i < 102; i++ {
		fmt.Printf("=")
	}
	fmt.Printf("\n")
	for {
		if (count <= 0) {
			count--
			fmt.Printf("%20v %16[1]X %064[1]b\n", count)
			break
		}
		//time.Sleep(time.Second)
		fmt.Printf("%20v %16[1]X %064[1]b\n", count)
		count *= 2
		//count++
	}
}
