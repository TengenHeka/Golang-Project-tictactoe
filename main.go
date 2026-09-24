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

func main() {
    board := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

    currentPlayer := "X"

    for {
        printBoard(board)

        var position int

        fmt.Printf("Player %s, choose a position (1-9): ", currentPlayer)
        fmt.Scan(&position)
	}
}