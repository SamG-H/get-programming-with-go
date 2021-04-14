package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Initiating rocket launch")
	fmt.Println("Beginning countdown sequence")
	count := 10
	for count > 0 {
		if rand.Intn(100) == 1 {
			break
		}
		fmt.Printf("%v...\n", count)
		time.Sleep(time.Second)
		count--
	}
	if count == 0 {
		fmt.Println("Launch successful")
	} else {
		fmt.Println("Launch failed")
	}
}
