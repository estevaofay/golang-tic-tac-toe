package player

import (
	"bufio"
	"fmt"
	"os"
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
	// Check if move is valid
	// userMoveInt := 0
}

func (p *Player) PrintPlayerScore() {
	fmt.Printf("Player %s has as score of %d\n", p.Name, p.Score)
}
