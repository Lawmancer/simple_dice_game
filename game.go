package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const wild = 4
const wildValue = 0
const rounds = 4
const totalDice = 5

func playGame(playerNames []string) (winner player, error error) {
	if len(playerNames) == 0 {
		error = errors.New("no players to play the game")
		return
	}
	var unsortedGamers []player
	for _, name := range playerNames {
		var p player
		p.name = name
		unsortedGamers = append(unsortedGamers, p)
	}

	gamers, err := randomFirstPlayer(unsortedGamers)
	if err != nil {
		error = errors.New(err.Error())
		return
	}

	listPlayers(gamers)
	roundNum := 0
	for roundNum < rounds {
		roundNum++
		fmt.Println("\n-----------------------")
		fmt.Println("Starting Round", roundNum)
		for i := range gamers {
			err := gamers[i].takeTurn(totalDice)
			if err != nil {
				error = errors.New(err.Error())
				return
			}
		}
		gamers[0].turnOrder = gamers[0].turnOrder + len(gamers) // move to end of line
		gamers = sortPlayers(gamers)
	}
	gamers = getWinnerFirst(gamers)
	winner = gamers[0]
	for _, p := range gamers {
		fmt.Printf("%s has a final score of %d\n", p.name, p.score)
	}
	fmt.Printf("\n%s wins with a final score of %d\n", winner.name, winner.score)

	return
}

func randomFirstPlayer(gamers []player) (sortedGamers []player, err error) {
	if len(gamers) == 0 {
		err = errors.New("no gamers to sort")
		return
	}
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(len(gamers))
	gamers[r].turnOrder = 1
	next := 2 // we already set 1
	for i := 0; i < len(gamers); i++ {
		if gamers[i].turnOrder == 0 {
			gamers[i].turnOrder = next
			next++
		}
	}
	return sortPlayers(gamers), err
}

func sortPlayers(gamers []player) []player {
	sort.Slice(gamers, func(i, j int) bool {
		return gamers[i].turnOrder < gamers[j].turnOrder
	})
	return gamers
}

func listPlayers(gamers []player) {
	for i, g := range gamers {
		fmt.Printf("Player %d: %s\n", i+1, g.name)
	}
	fmt.Println()
}

func getWinnerFirst(gamers []player) []player {
	sort.Slice(gamers, func(i, j int) bool {
		return gamers[i].score < gamers[j].score
	})

	return gamers
}
