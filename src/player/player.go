package player

import (
	"bufio"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

type Player struct {
	// Define Player struct here
}

func (p *Player) GetPlayerMove() int {
	// Get player Move from Stdin
	// Check if move is valid
	// userMoveInt := 0
}
