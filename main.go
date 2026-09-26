package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("================================")
	fmt.Println("        TIC-TAC-TOE IN GO")
	fmt.Println("================================")

	fmt.Println("1. Two Players")
	fmt.Println("2. Player vs Computer")

	var mode int

	for {
		fmt.Print("Choose game mode: ")
		fmt.Scan(&mode)

		if mode == 1 || mode == 2 {
			break
		}

		fmt.Println("Invalid choice. Please select 1 or 2.")
	}

	playerXWins := 0
	playerOWins := 0
	draws := 0

	for {
		board := []string{
			"1", "2", "3",
			"4", "5", "6",
			"7", "8", "9",
		}

		currentPlayer := "X"

		for {
			printBoard(board)

			// Computer's turn
			if mode == 2 && currentPlayer == "O" {
				position := computerMove(board)
				board[position] = "O"

				fmt.Printf("Computer chose position %d\n", position+1)

			} else {
				// Human player's turn
				var position int

				fmt.Printf("Player %s, choose a position (1-9): ", currentPlayer)
				fmt.Scan(&position)

				if position < 1 || position > 9 {
					fmt.Println("Invalid position. Choose between 1 and 9.")
					continue
				}

				if board[position-1] == "X" ||
					board[position-1] == "O" {
					fmt.Println("That position is already taken.")
					continue
				}

				board[position-1] = currentPlayer
			}

			// Check winner
			if checkWinner(board, currentPlayer) {
				printBoard(board)

				if mode == 2 && currentPlayer == "O" {
					fmt.Println("Computer wins!")
					playerOWins++
				} else {
					fmt.Printf("Player %s wins!\n", currentPlayer)

					if currentPlayer == "X" {
						playerXWins++
					} else {
						playerOWins++
					}
				}

				break
			}

			// Check draw
			if checkDraw(board) {
				printBoard(board)
				fmt.Println("It's a draw!")
				draws++
				break
			}

			// Switch player
			if currentPlayer == "X" {
				currentPlayer = "O"
			} else {
				currentPlayer = "X"
			}
		}

		// Scoreboard
		fmt.Println()
		fmt.Println("========== SCORE ==========")
		fmt.Printf("Player X: %d\n", playerXWins)
		fmt.Printf("Player O: %d\n", playerOWins)
		fmt.Printf("Draws: %d\n", draws)
		fmt.Println("===========================")

		var again string

		fmt.Print("Play again? (y/n): ")
		fmt.Scan(&again)

		if again != "y" && again != "Y" {
			fmt.Println("Thanks for playing!")
			break
		}
	}
}
