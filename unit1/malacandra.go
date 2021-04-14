package main
import "fmt"

func main(){
	const hoursPerDay = 24
	const time = 28
	const distance = 56000000
	const speed = distance / (time * hoursPerDay)
	fmt.Printf("You need to travel %v kmh to get to Mars in %v days if mars was %v km away\n", speed, time, distance)
}
