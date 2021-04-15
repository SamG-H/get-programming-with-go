package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64

// sensor function type
type sensor func() kelvin

func realSensor() kelvin {
	return 0
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(1001))
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

func main() {
	var offset kelvin = 5
	sensor := calibrate(realSensor, offset)
	fmt.Println(sensor())
	// changing offset doesn't change return value of sensor because 
	// offset was passed by value to calibrate when closure was created
	offset = 6
	fmt.Println(sensor())
	sensor = calibrate(fakeSensor, offset)
	fmt.Println(sensor())
	fmt.Println(sensor())
	fmt.Println(sensor())
}
