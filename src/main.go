package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var board = [][]string{{".", ".", "."}, {".", ".", "."}, {".", ".", "."}}
var isUser1Turn = true

func main() {
	fmt.Println("Hello, world.")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Usuário 1 - Digite seu nome:")
	user1, _ := reader.ReadString('\n')
	fmt.Println(user1)

	fmt.Println("Usuário 2 - Digite seu nome:")
	user2, _ := reader.ReadString('\n')
	fmt.Println(user2)

	fmt.Println("Começando o jogo")

	fmt.Println("o ->", user1)
	fmt.Println("x ->", user2)

	for {
		clearScreen()

		fmt.Println("Usuário : qual casa (1-9)?")
		printGameBoard()

		userMove, _ := reader.ReadString('\n')

		userMove = strings.ReplaceAll(userMove, "\n", "")
		userMoveInt, _ := strconv.Atoi(userMove)
		setMoveOnBoard(userMoveInt)
		if checkGameCompletion() {
			break
		}
	}

}

func checkGameCompletion() bool {
	return checkColumns() || checkRows() || checkDiagonals() || checkDrawGame()
}

func checkDrawGame() bool {
	isDrawGame := true
	for _, line := range board {
		for _, value := range line {
			isDrawGame = isDrawGame && !checkDot(value)
		}
	}
	return isDrawGame
}

func checkDiagonals() bool {
	return checkFirstDiagonal() || checkSecondDiagonal()
}

func checkFirstDiagonal() bool {
	if checkDot(board[0][0]) || checkDot(board[1][1]) || checkDot(board[2][2]) {
		return false
	}
	return board[0][0] == board[1][1] && board[1][1] == board[2][2]
}

func checkSecondDiagonal() bool {
	if checkDot(board[0][2]) || checkDot(board[1][1]) || checkDot(board[2][0]) {
		return false
	}
	return board[0][2] == board[1][1] && board[1][1] == board[2][0]
}

func checkRows() bool {
	return checkRow(0) || checkRow(1) || checkRow(2)
}

func checkRow(index int) bool {
	if checkDot(board[0][index]) || checkDot(board[1][index]) || checkDot(board[2][index]) {
		return false
	}
	return board[0][index] == board[1][index] && board[1][index] == board[2][index]
}

func checkColumns() bool {
	return checkColumn(0) || checkColumn(1) || checkColumn(2)
}

func checkColumn(index int) bool {
	if checkDot(board[index][0]) || checkDot(board[index][1]) || checkDot(board[index][2]) {
		return false
	}
	return board[index][0] == board[index][1] && board[index][1] == board[index][2]
}

func checkDot(value string) bool {
	return value == "."
}

func setMoveOnBoard(userMove int) {
	x, y := GetCoordinates(userMove)
	if isUser1Turn {
		board[x][y] = "o"
	} else {
		board[x][y] = "x"
	}
	isUser1Turn = !isUser1Turn
}

func printGameBoard() {
	for _, line := range board {
		fmt.Print("|")
		for _, value := range line {
			fmt.Print(value)
		}
		fmt.Println("|")
	}

}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func GetCoordinates(input int) (int, int) {
	if input%3 == 0 {
		return input/3 - 1, 2
	}
	return input / 3, input%3 - 1

}
