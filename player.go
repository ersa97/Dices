package dices

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	MIN_NUM_DICE = 1
	MAX_NUM_DICE = 6
)

// Player
type Player struct {
	ID        int
	Point     int
	Dices     []int
	SaveDices []int
}

func NewPlayer(id int, dices []int) Player {
	return Player{
		ID:        id,
		Point:     0,
		SaveDices: []int{},
		Dices:     dices,
	}
}

func (p *Player) Roll() {
	rand.Seed(time.Now().UnixNano())
	for i := range p.Dices {
		res := rand.Intn(MAX_NUM_DICE-MIN_NUM_DICE+1) + MIN_NUM_DICE
		p.Dices[i] = res
	}
}

func (p *Player) Print() {
	fmt.Printf("\tPemain #%d (%d): ", p.ID, p.Point)
	if len(p.Dices) > 0 {
		var str []string
		for _, v := range p.Dices {
			str = append(str, fmt.Sprintf("%d", v))
		}

		fmt.Printf("%s\n", strings.Join(str, ", "))
		return
	}
	fmt.Printf("-\n")
}

func (p *Player) Evaluate(neighbor *Player) {
	var d []int
	for i := range p.Dices {
		if p.Dices[i] == MAX_NUM_DICE {
			p.Point++
		}
		if p.Dices[i] == MIN_NUM_DICE {
			neighbor.SaveDices = append(neighbor.SaveDices, p.Dices[i])
		}
		if p.Dices[i] > MIN_NUM_DICE && p.Dices[i] < MAX_NUM_DICE {
			d = append(d, p.Dices[i])
		}
	}
	p.Dices = append(d, p.SaveDices...)
	p.SaveDices = []int{}
}
