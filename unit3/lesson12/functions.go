package main

import "fmt"

func kelvinToCelsius(k float64) float64 {
	c := k - 273.15
	return c
}

func celsiusToFahrenheit(c float64) float64 {
	f := (c * 9.0 / 5.0) + 32.0
	return f
}

func kelvinToFahrenheit(k float64) float64 {
	c := kelvinToCelsius(k)
	f := celsiusToFahrenheit(c)
	return f
}

func main() {
	kelvin := 294.0
	celsius := kelvinToCelsius(kelvin)
	fahrenheit := celsiusToFahrenheit(celsius)
	fmt.Println(kelvin, "° K is", celsius, "° C")
	fmt.Println(celsius, "° C is", fahrenheit, "° F")
	fmt.Println(kelvin, "° K is", kelvinToFahrenheit(kelvin), "° F")

}
