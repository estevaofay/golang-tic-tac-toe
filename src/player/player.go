package player

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

type Player struct {
	Name  string
	Score int
	Mark  rune
}

func (p *Player) GetPlayerMove() int {
	fmt.Printf("%s what is your next move? (1-9):\n", p.Name)

	// Get player Move from Stdin
	text, _ := reader.ReadString('\n')
	text = strings.ReplaceAll(text, "\n", "")

	// Check if move is valid
	userMoveInt, _ := strconv.Atoi(text)

	// userMoveInt := 0
	if userMoveInt < 1 || userMoveInt > 9 {
		fmt.Println("Please insert a number between 1 and 9")
	}

}

func (p *Player) PrintPlayerScore() {
	fmt.Printf("Player %s has as score of %d\n", p.Name, p.Score)
}
