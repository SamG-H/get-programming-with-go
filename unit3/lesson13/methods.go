package main

import "fmt"

type celsius float64
type fahrenheit float64
type kelvin float64

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * (5.0 / 9.0))
}

func (f fahrenheit) kelvin() kelvin {
	return f.celsius().kelvin()
}

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c*(9.0/5.0) + 32))
}

func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

func (k kelvin) fahrenheit() fahrenheit {
	return k.celsius().fahrenheit()
}

func main() {
	var c celsius = 100.0
	var f fahrenheit = 212.0
	var k kelvin = 373.15

	fmt.Println("Water boils at these temperatures")
	fmt.Println("C to F:", c.fahrenheit())
	fmt.Println("K to F:", k.fahrenheit())
	fmt.Println("F to C:", f.celsius())
	fmt.Println("K to C:", k.celsius())
	fmt.Println("F to K:", f.kelvin())
	fmt.Println("C to K:", c.kelvin())

	var zeroValue kelvin
	fmt.Println("Kelvin zero value:", temp)
}
