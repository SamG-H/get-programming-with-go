package main

import (
	"fmt"
)

type coordinate struct {
	d, m, s float64
	h       rune
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

type location struct {
	lat, long float64
}

func newLocation(lat, long coordinate) location {
	return location{lat.decimal(), long.decimal()}
}

func main() {
	spirit := newLocation(coordinate{d: 14, m: 34, s: 6.2, h: 'S'}, coordinate{d: 175, m: 28, s: 21.5, h: 'E'})
	opportunity := newLocation(coordinate{1, 56, 46.3, 'S'}, coordinate{354, 28, 24.2, 'E'})
	curiosity := newLocation(coordinate{4, 35, 22.2, 'S'}, coordinate{137, 26, 30.1, 'E'})
	insight := newLocation(coordinate{4, 30, 0, 'N'}, coordinate{135, 54, 0, 'E'})
	fmt.Printf("%-12s%-20v\n", "Spirit", spirit)
	fmt.Printf("%-12s%-20v\n", "Opportunity", opportunity)
	fmt.Printf("%-12s%-20v\n", "Curiosity", curiosity)
	fmt.Printf("%-12s%-20v\n", "InSight", insight)
}
