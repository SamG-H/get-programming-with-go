package main

import "fmt"

const (
	PrintColor = "\033[38;5;%dm%-9v\033[39;49m"
	tempFmt = "%.1f"
)

type getDataFunc func(int) (col1, col2, col3 string, col1color, col2color, col3color int)

type kelvin float64

func (k kelvin) fahrenheit() fahrenheit {
	return k.celsius().fahrenheit()
}

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

func (k kelvin) color() int {
	switch {
	case k <= 273.15:
		return 39
	case k >= 373.1:
		return 1
	default:
		return 2
	}
}

type fahrenheit float64

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * (5.0 / 9.0))
}

func (f fahrenheit) kelvin() kelvin {
	return f.celsius().kelvin()
}

func (f fahrenheit) color() int {
	switch {
	case f <= 32:
		return 39
	case f >= 212:
		return 1
	default:
		return 2
	}
}

type celsius float64

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit(c*(9.0/5.0) + 32.0)
}

func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}

func (c celsius) color() int {
	switch {
	case c <= 0:
		return 39
	case c >= 100:
		return 1
	default:
		return 2
	}
}

func drawHeader(title1, title2, title3 string) {
	drawRowSeparator()
	fmt.Printf("| %-9v| %-9v| %-9v|\n", title1, title2, title3)
	drawRowSeparator()
}

func drawRowSeparator() {
	fmt.Printf("==================================\n")
}

func drawRow(column1, column2, column3 string, color1, color2, color3 int) {
	fmt.Printf("| "+PrintColor, color1, column1)
	fmt.Printf("| "+PrintColor, color2, column2)
	fmt.Printf("| "+PrintColor, color3, column3)
	fmt.Printf("|\n")
}

func fToCToK(rowNum int) (col1, col2, col3 string, fColor, cColor, kColor int) {
	f := fahrenheit((rowNum * 5) - 40)
	c := f.celsius()
	k := f.kelvin()
	col1 = fmt.Sprintf(tempFmt, f)
	col2 = fmt.Sprintf(tempFmt, c)
	col3 = fmt.Sprintf(tempFmt, k)
	fColor = f.color()
	cColor = c.color()
	kColor = k.color()
	return col1, col2, col3, fColor, cColor, kColor
}

func cToFToK(rowNum int) (col1, col2, col3 string, cColor, fColor, kColor int) {
	// when declaring return types with names it is like you are
	// initializing them
	// if you name return types you can do a naked return
	// and just type in return (like implicit return in Ruby)
	c := celsius((rowNum * 5) - 40)
	f := c.fahrenheit()
	k := c.kelvin()
	col1 = fmt.Sprintf(tempFmt, c)
	col2 = fmt.Sprintf(tempFmt, f)
	col3 = fmt.Sprintf(tempFmt, k)
	cColor = c.color()
	fColor = f.color()
	kColor = k.color()
	return col1, col2, col3, cColor, fColor, kColor
}

// add names to multi returns & single return if value isn't obvious
// function names should convey meaning
func kToFToC(rowNum int) (col1, col2, col3 string, kColor, fColor, cColor int) {
	k := kelvin(rowNum * 5)
	f := k.fahrenheit()
	c := k.celsius()
	col1 = fmt.Sprintf(tempFmt, k)
	col2 = fmt.Sprintf(tempFmt, f)
	col3 = fmt.Sprintf(tempFmt, c)
	kColor = k.color()
	fColor = f.color()
	cColor = c.color()
	return col1, col2, col3, kColor, fColor, cColor
}

func drawTable(hdr1, hdr2, hdr3 string, numRows int, getData getDataFunc) {
	drawHeader(hdr1, hdr2, hdr3)
	for i := 0; i < numRows; i++ {
		col1, col2, col3, color1, color2, color3 := getData(i)
		drawRow(col1, col2, col3, color1, color2, color3)
	}
	drawRowSeparator()
}

func main() {
	drawTable("°F", "°C", "K", 66, fToCToK)
	drawTable("°C", "°F", "K", 37, cToFToK)
	drawTable("K", "°F", "°C", 130, kToFToC)
}
