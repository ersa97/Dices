package dices

import (
	"fmt"
	"strings"
)

type Game struct {
	Players []Player
}

func NewGame(players []Player) Game {
	return Game{
		Players: players,
	}
}

func (g *Game) Start() {
	count := 1
	for {
		g.roll(count)
		g.evaluate()
		fmt.Printf("\n================================\n\n")
		if !g.isNext() {
			break
		}
		count++
	}
}

func (g *Game) roll(f int) {
	fmt.Printf("Game ke %d:\n", f)
	for i := range g.Players {
		g.Players[i].Roll()
	}
	for i := range g.Players {
		g.Players[i].Print()
	}
}

func (g *Game) evaluate() {
	fmt.Printf("\nSetelah Evaluasi:\n")
	for i := range g.Players {
		var neighbor *Player
		if i == len(g.Players)-1 {
			neighbor = &g.Players[0]
		} else {
			neighbor = &g.Players[i+1]
		}
		g.Players[i].Evaluate(neighbor)
	}
	// Reappend SaveDices into Dices
	g.Players[0].Dices = append(g.Players[0].Dices, g.Players[0].SaveDices...)
	for i := range g.Players {
		g.Players[i].Print()
	}
}

func (g *Game) isNext() bool {
	count := 0
	for _, p := range g.Players {
		if len(p.Dices) > 0 {
			count++
		}
	}
	if count >= 2 {
		return true
	}
	g.win()
	return false
}

func (g *Game) win() {
	highest := 0
	winners := []string{}

	for _, p := range g.Players {
		name := fmt.Sprintf("Pemain %d", p.ID)
		if p.Point > highest {
			highest = p.Point
			winners = []string{name}
			continue
		}
		if p.Point == highest {
			winners = append(winners, name)
			continue
		}
	}

	fmt.Printf("\nPemenang:\n%s\nDengan point (%d)\n", strings.Join(winners, ", "), highest)
}
