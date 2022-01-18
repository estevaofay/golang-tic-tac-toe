package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var board = [][]string{{".", ".", "."}, {".", ".", "."}, {".", ".", "."}}

func main() {
	// fmt.Println("Hello, world.")
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

	fmt.Println("Usuário 1: qual casa (1-9)?")
	userMove, _ := reader.ReadString('\n')

	userMove = strings.ReplaceAll(userMove, "\n", "")
	userMoveInt, _ := strconv.Atoi(userMove)

	fmt.Println(userMoveInt)

	printGameBoard()

	fmt.Println(board[0][0])

}

func printGameBoard() {
	fmt.Println(board)
}

func GetCoordinates(input int) (int, int) {
	return 1, 2
}
