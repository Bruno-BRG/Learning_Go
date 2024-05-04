package main

import (
	"fmt"
)

type Player int

const (
	Empty Player = iota
	X
	O
)

type Board [3][3]Player

func display(board Board) {
	// i want the display to have | as the vertial line and - as the horizontal line
	// i want the display to have the numbers 1, 2, 3 on the top and 1, 2, 3 on the left side

	fmt.Println(" 0  1  2 ")
	for _, row := range board {
		for _, cell := range row {
			switch cell {
			case Empty:
				fmt.Print("| |")
			case X:
				fmt.Print("|X|")
			case O:
				fmt.Print("|O|")
			}
		}
		fmt.Println()
	}
}

func checkWinner(board Board) Player {
	for i := 0; i < 3; i++ {
		if board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			return board[i][0]
		}
		if board[0][i] == board[1][i] && board[1][i] == board[2][i] {
			return board[0][i]
		}
	}
	if board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		return board[0][0]
	}
	if board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		return board[0][2]
	}
	return Empty
}

func main() {
	var board Board
	var player Player = X
	for {
		display(board)
		fmt.Println("Player", player)
		var x int
		var y int
		fmt.Print("Enter x: ")
		fmt.Scan(&x)
		fmt.Print("Enter y: ")
		fmt.Scan(&y)
		if board[x][y] != Empty {
			fmt.Println("Cell is already occupied")
			continue
		}
		board[x][y] = player
		if winner := checkWinner(board); winner != Empty {
			display(board)
			fmt.Println("Player", winner, "wins!")
			break
		}
		if player == X {
			player = O
		} else {
			player = X
		}
	}
}
