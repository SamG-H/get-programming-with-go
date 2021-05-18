package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width         = 80
	height        = 15
	ansiEscapeSeq = "\033c\x0c"
	blackSquare   = "\xE2\xAC\x9B"
	whiteSqure    = "\xE2\xAC\x9C"
	greenSquare   = "\xF0\x9F\x9F\xA9"
)

type Universe [][]bool

func (u Universe) Show() {
	for _, row := range u {
		for _, column := range row {
			switch {
			case column:
				fmt.Printf(greenSquare)
			default:
				fmt.Printf(blackSquare)
			}
		}
		fmt.Printf("\n")
	}
}

func (u Universe) Seed() {
	for _, row := range u {
		for i := range row {
			if rand.Intn(4) == 1 {
				row[i] = true
			}
		}
	}
}

func (u Universe) TestSeed() {
	for i, row := range u {
		for j := range row {
			switch {
			case i == 1 && j == 1:
				row[j] = true
			case i == 2 && j == 2:
				row[j] = true
			case i == 1 && j == 2:
				row[j] = true
			//case i == 1 && j == 0:
			//	row[j] = true
			}
		}
	}
}

func (u Universe) Alive(x, y int) bool {
	y = (height + y) % height
	x = (width + x) % width
	return u[y][x]
}

func (u Universe) Neighbors(x, y int) int {
	var neighbors int

	for i := y - 1; i <= y + 1; i++ {
		for j := x - 1; j <= x + 1; j++ {
			if i == y && j == x {
				continue;
			}
			if u.Alive(j, i) {
				neighbors++
			}
		}
	}
	return neighbors
}

func (u Universe) Next(x, y int) bool {
	n := u.Neighbors(x, y)
	alive := u.Alive(x, y)
	if n < 4 && n > 1 && alive {
		return true
	} else if n == 3 && !alive {
		return true
	} else {
		return false
	}
}

func Step(a, b Universe) {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			b[i][j] = a.Next(j, i)
		}
	}
}

func MakeUniverse() Universe {
	u := make(Universe, height)
	for i, _ := range u {
		u[i] = make([]bool, width)
	}
	return u
}

func main() {
	fmt.Println(ansiEscapeSeq)
	rand.Seed(time.Now().UTC().UnixNano())
	newUniverse := MakeUniverse()
	nextUniverse := MakeUniverse()
	newUniverse.Seed()
	for {
		newUniverse.Show()
		time.Sleep(50 * time.Millisecond)
		fmt.Println(ansiEscapeSeq)
		Step(newUniverse, nextUniverse)
		newUniverse, nextUniverse = nextUniverse, newUniverse
	}
}
