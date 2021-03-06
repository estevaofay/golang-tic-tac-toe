package board

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"tictactoe/player"
)

type Board struct {
	Board      [][]rune
	Player1    player.Player
	Player2    player.Player
	PlayerTurn int
}

var isUser1Turn = true
var reader = bufio.NewReader(os.Stdin)

func PlayTheGame() {

	fmt.Println("Welcome to tic-tac-toe!")
	fmt.Println()
	fmt.Println("Player 1 - Type your name:")

	player1 := player.Player{
		Name:  readName(),
		Score: 0,
		Mark:  'x',
	}

	fmt.Println()
	fmt.Println("Player 2 - Type your name:")

	player2 := player.Player{
		Name:  readName(),
		Score: 0,
		Mark:  'o',
	}

	board := Board{
		Board:      [][]rune{{'.', '.', '.'}, {'.', '.', '.'}, {'.', '.', '.'}},
		Player1:    player1,
		Player2:    player2,
		PlayerTurn: 1,
	}

	fmt.Println("Starting game")

	for {
		board.clearScreen()
		board.printGameBoard()
		userMoveInt := -1

		if board.PlayerTurn == 1 {
			userMoveInt = board.Player1.GetPlayerMove()
			board.PlayerTurn = 2
		} else {
			userMoveInt = board.Player2.GetPlayerMove()
			board.PlayerTurn = 1
		}

		board.setMoveOnBoard(userMoveInt)
		if board.isGameComplete() {
			if !board.checkDrawGame() {
				if board.PlayerTurn == 1 {
					board.Player2.Score++
				} else {
					board.Player1.Score++
				}
			}

			board.clearScreen()
			board.printGameBoard()
			board.printBoardScore()
			fmt.Println("Do you wish play another round? (y/n)")
			continuePlaying, _ := reader.ReadString('\n')
			continuePlaying = strings.ReplaceAll(continuePlaying, "\n", "")
			if continuePlaying == "n" {
				break
			} else {
				board.clearScreen()
				board.clearBoard()

			}
		}
	}
}

func (b *Board) clearBoard() {
	b.Board = [][]rune{{'.', '.', '.'}, {'.', '.', '.'}, {'.', '.', '.'}}
}

func readName() string {
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.ReplaceAll(name, "\n", "")
	return name
}

func (b *Board) printBoardScore() {
	b.Player1.PrintPlayerScore()
	b.Player2.PrintPlayerScore()
}

func (b *Board) isGameComplete() bool {
	return b.checkColumns() || b.checkRows() || b.checkDiagonals() || b.checkDrawGame()
}

func (b *Board) checkDrawGame() bool {
	isDrawGame := true
	for _, line := range b.Board {
		for _, value := range line {
			isDrawGame = isDrawGame && !b.checkDot(value)
		}
	}
	return isDrawGame
}

func (b *Board) checkDiagonals() bool {
	return b.checkFirstDiagonal() || b.checkSecondDiagonal()
}

func (b *Board) checkFirstDiagonal() bool {
	if b.checkDot(b.Board[0][0]) || b.checkDot(b.Board[1][1]) || b.checkDot(b.Board[2][2]) {
		return false
	}
	return b.Board[0][0] == b.Board[1][1] && b.Board[1][1] == b.Board[2][2]
}

func (b *Board) checkSecondDiagonal() bool {
	if b.checkDot(b.Board[0][2]) || b.checkDot(b.Board[1][1]) || b.checkDot(b.Board[2][0]) {
		return false
	}
	return b.Board[0][2] == b.Board[1][1] && b.Board[1][1] == b.Board[2][0]
}

func (b *Board) checkRows() bool {
	return b.checkRow(0) || b.checkRow(1) || b.checkRow(2)
}

func (b *Board) checkRow(index int) bool {
	if b.checkDot(b.Board[0][index]) || b.checkDot(b.Board[1][index]) || b.checkDot(b.Board[2][index]) {
		return false
	}
	return b.Board[0][index] == b.Board[1][index] && b.Board[1][index] == b.Board[2][index]
}

func (b *Board) checkColumns() bool {
	return b.checkColumn(0) || b.checkColumn(1) || b.checkColumn(2)
}

func (b *Board) checkColumn(index int) bool {
	if b.checkDot(b.Board[index][0]) || b.checkDot(b.Board[index][1]) || b.checkDot(b.Board[index][2]) {
		return false
	}
	return b.Board[index][0] == b.Board[index][1] && b.Board[index][1] == b.Board[index][2]
}

func (b *Board) checkDot(value rune) bool {
	return value == '.'
}

func (b *Board) setMoveOnBoard(userMove int) {
	x, y := GetCoordinates(userMove)
	if isUser1Turn {
		b.Board[x][y] = 'o'
	} else {
		b.Board[x][y] = 'x'
	}
	isUser1Turn = !isUser1Turn
}

func (b *Board) printGameBoard() {
	fmt.Println("Board:")
	fmt.Println()
	for _, line := range b.Board {
		fmt.Print("|")
		for _, value := range line {
			fmt.Print(string(value))
		}
		fmt.Println("|")
	}

	fmt.Println()
}

func (b *Board) clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func GetCoordinates(input int) (int, int) {
	return (input - 1) / 3, (input - 1) % 3
}
