package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const noPlayersErr = "no players found"

func playGame(playerNames []string, totalDice int) (winner player, error error) {
	if len(playerNames) == 0 {
		error = errors.New(noPlayersErr)
		return
	}
	var unsortedGamers []player
	for _, name := range playerNames {
		var p player
		p.name = name
		unsortedGamers = append(unsortedGamers, p)
	}

	gamers, err := randomTurnOrder(unsortedGamers)
	sortPlayers(gamers)
	if err != nil {
		return
	}

	listPlayers(gamers)
	roundNum := 0
	for roundNum < len(gamers) {
		roundNum++
		fmt.Println("\n-----------------------")
		fmt.Println("Starting Round", roundNum)
		for i := range gamers {
			err := gamers[i].takeTurn(totalDice)
			if err != nil {
				return
			}
		}
		gamers[0].turnOrder = gamers[0].turnOrder + len(gamers) // move to end of line
		gamers = sortPlayers(gamers)
	}

	winners := getWinners(gamers)
	error = announceWinners(winners)
	if error != nil {
		return
	}

	return
}

func randomTurnOrder(gamers []player) (sortedGamers []player, err error) {
	if len(gamers) == 0 {
		err = errors.New(noPlayersErr)
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
	return gamers, err
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

func getWinners(gamers []player) (winners []player) {
	// sort by scores
	sort.Slice(gamers, func(i, j int) bool {
		return gamers[i].score < gamers[j].score
	})

	winningScore := gamers[0].score
	for _, gamer := range gamers {
		// first player will always be appended
		// players with tied scores will also be appended
		if gamer.score == winningScore {
			winners = append(winners, gamer)
		} else {
			// no more ties
			break
		}
	}

	return winners
}

// announceWinners with some prettified grammar (uses oxford comma)
func announceWinners(winners []player) error {
	numWinners := len(winners)
	winPlural := ""
	oxfordComma := ","
	and := ""
	if numWinners == 0 {
		err := errors.New("nobody won the game")
		return err
	}
	if numWinners == 1 {
		winPlural = "s"
	}
	if numWinners <= 2 {
		oxfordComma = ""
	}
	if numWinners > 1 {
		and = "and "
	}
	last := numWinners - 1
	for i, winner := range winners {
		if i != last {
			fmt.Printf("%s%s ", winner.name, oxfordComma)
		} else {
			oxfordComma = ""
			fmt.Printf("%s%s%s ", and, winner.name, oxfordComma)
		}
	}
	fmt.Printf("win%s with a final score of %d!\n", winPlural, winners[0].score)
	return nil
}
