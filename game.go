package main

import "math/rand"

func printBoard(board []string) {
    println()
    println(" " + board[0] + " | " + board[1] + " | " + board[2])
    println("---+---+---")
    println(" " + board[3] + " | " + board[4] + " | " + board[5])
    println("---+---+---")
    println(" " + board[6] + " | " + board[7] + " | " + board[8])
    println()
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