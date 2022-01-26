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

func (p *Player) PrintPlayerScore() {
	fmt.Printf("Player %s has a score of %d!\n", p.Name, p.Score)
}

func (p *Player) GetPlayerMove() int {
	userMoveInt := 0
	for {
		fmt.Println(p.Name, "what is your next move? (1-9)")
		userMove, _ := reader.ReadString('\n')
		userMove = strings.ReplaceAll(userMove, "\n", "")
		userMoveInt, _ = strconv.Atoi(userMove)
		if userMoveInt < 9 && userMoveInt > 0 {
			break
		} else {
			fmt.Println("Please input a number from 1 to 9")
		}
	}
	return userMoveInt
}
