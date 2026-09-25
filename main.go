package main

import "fmt"

func printBoard(board []string) {
	fmt.Println()
	fmt.Printf(" %s | %s | %s\n", board[0], board[1], board[2])
	fmt.Println("---+---+---")
	fmt.Printf(" %s | %s | %s\n", board[3], board[4], board[5])
	fmt.Println("---+---+---")
	fmt.Printf(" %s | %s | %s\n", board[6], board[7], board[8])
	fmt.Println()
}

func checkWinner(board []string, player string) bool {
	wins := [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
		{0, 4, 8},
		{2, 4, 6},
	}

	for _, win := range wins {
		if board[win[0]] == player &&
			board[win[1]] == player &&
			board[win[2]] == player {
			return true
		}
	}

	return false
}

func checkDraw(board []string) bool {
	for _, position := range board {
		if position != "X" && position != "O" {
			return false
		}
	}
	return true
}

func main() {
	board := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	currentPlayer := "X"

	for {
		printBoard(board)

		var position int

		fmt.Printf("Player %s, choose a position (1-9): ", currentPlayer)
		fmt.Scan(&position)

		if position < 1 || position > 9 {
			fmt.Println("Invalid position. Choose between 1 and 9.")
			continue
		}

		if board[position-1] == "X" || board[position-1] == "O" {
			fmt.Println("That position is already taken.")
			continue
		}

		board[position-1] = currentPlayer

		if checkWinner(board, currentPlayer) {
			printBoard(board)
			fmt.Printf("Player %s wins!\n", currentPlayer)
			break
		}

		if checkDraw(board) {
			printBoard(board)
			fmt.Println("It's a draw!")
			break
		}

		if currentPlayer == "X" {
			currentPlayer = "O"
		} else {
			currentPlayer = "X"
		}
	}
}
