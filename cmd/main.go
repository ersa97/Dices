package main

import (
	"fmt"

	dices "github.com/ersa97/dices"
)

func main() {
	const N = 3
	const M = 4

	game := dices.NewGame(generatePlayers(N, M))

	fmt.Printf("Pemain = %d, Dadu = %d\n", N, M)
	fmt.Printf("================================\n\n")
	game.Start()
}

func generatePlayers(p, d int) []dices.Player {
	var players []dices.Player
	for i := 0; i < p; i++ {
		d := make([]int, d)
		players = append(players, dices.NewPlayer(i+1, d))
	}
	return players
}
