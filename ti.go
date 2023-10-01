package main

import (
	"fmt"
)

var board [3][3]string
var currentPlayer string

func initializeBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = " "
		}
	}
}

func printBoard() {
	fmt.Println("  0 1 2")
	for i := 0; i < 3; i++ {
		fmt.Printf("%d %s|%s|%s\n", i, board[i][0], board[i][1], board[i][2])
		if i < 2 {
			fmt.Println("  -----")
		}
	}
}

func makeMove(row, col int) bool {
	if row >= 0 && row < 3 && col >= 0 && col < 3 && board[row][col] == " " {
		board[row][col] = currentPlayer
		return true
	}
	return false
}

func checkWinner() string {
	for i := 0; i < 3; i++ {
		if board[i][0] != " " && board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			return board[i][0]
		}
		if board[0][i] != " " && board[0][i] == board[1][i] && board[1][i] == board[2][i] {
			return board[0][i]
		}
	}
	if board[0][0] != " " && board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		return board[0][0]
	}
	if board[0][2] != " " && board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		return board[0][2]
	}
	return " "
}

func main() {
	initializeBoard()
	currentPlayer = "X"
	fmt.Println("欢迎来到井字游戏！")
	fmt.Println("玩家 X 先开始。")

	for {
		printBoard()
		fmt.Printf("玩家 %s，请输入行和列 (例如：0 1): ", currentPlayer)
		var row, col int
		_, err := fmt.Scanf("%d %d", &row, &col)
		if err != nil || row < 0 || row > 2 || col < 0 || col > 2 {
			fmt.Println("请输入有效的行和列。")
			continue
		}
		if makeMove(row, col) {
			winner := checkWinner()
			if winner != " " {
				printBoard()
				fmt.Printf("玩家 %s 获胜！恭喜！\n", winner)
				break
			} else {
				if currentPlayer == "X" {
					currentPlayer = "O"
				} else {
					currentPlayer = "X"
				}
			}
		} else {
			fmt.Println("该位置已经被占用，请选择一个空位置。")
		}
	}
}
