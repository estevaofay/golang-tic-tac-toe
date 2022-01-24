package player

import "fmt"

type Player struct {
	Name  string
	Score int
	Mark  rune
}

func (p *Player) PrintPlayerScore() {
	fmt.Printf("Player %s has a score of %d!\n", p.Name, p.Score)
}
