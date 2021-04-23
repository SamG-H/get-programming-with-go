package main

import "fmt"

func displayBoard(board [8][8]rune) {
	for _, row := range board {
		for _, column := range row {
			if column == 0 {
				fmt.Printf("|%c", ' ')
				continue
			}
			fmt.Printf("|%c", column)
		}
		fmt.Printf("%c\n", '|')
	}
}

func main() {
	var board [8][8]rune

	//black is at the top
	board[0][0] = 'r'
	board[0][7] = 'r'
	board[0][1] = 'n'
	board[0][6] = 'n'
	board[0][2] = 'b'
	board[0][5] = 'b'
	board[0][3] = 'k'
	board[0][4] = 'q'
	for column := range board[1] {
		board[1][column] = 'p'
	}

	//white is at the bottom
	board[7][0] = 'R'
	board[7][7] = 'R'
	board[7][1] = 'N'
	board[7][6] = 'N'
	board[7][2] = 'B'
	board[7][5] = 'B'
	board[7][3] = 'K'
	board[7][4] = 'Q'
	for column := range board[6] {
		board[6][column] = 'P'
	}

	displayBoard(board)
}
