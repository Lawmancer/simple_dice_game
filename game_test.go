package main

import (
	"testing"
)

// integration test for unexpected errors or panic
func TestPlayGameReturnsPlayer(t *testing.T) {
	playerNames := []string{"Freddie Mercury", "Brian May", "Roger Taylor", "John Deacon "}
	totalDice := 4
	_, err := playGame(playerNames, totalDice)
	if err != nil {
		t.Errorf("Received unexpected error during play: %v", err)
	}
}

func TestPlayGameErrWithoutNames(t *testing.T) {
	expect := noPlayersErr
	totalDice := 6
	var playerNames []string
	_, err := playGame(playerNames, totalDice)
	if err == nil {
		t.Error("Did not receive expected error:")
	} else if err.Error() != expect {
		t.Errorf("Did not get expected error. Expected: %v, got: %v", expect, err.Error())
	}
}

// no real way to test random order, but at least ensure we get everyone back
func TestRandomTurnOrderReturnsFour(t *testing.T) {
	expect := 4
	gamers := getFourGamers()
	res, err := randomTurnOrder(gamers)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(res) != expect {
		t.Errorf("Did not received expected number of players back. Expected %v, got %v", expect, len(res))
	}
}

func TestRandomTurnOrderErrWithZeroGamers(t *testing.T) {
	var gamers []player
	_, err := randomTurnOrder(gamers)
	if err == nil {
		t.Errorf("Expected error")
	} else if err.Error() != noPlayersErr {
		t.Errorf("Expected error: %v", err)
	}
}

func TestSortPlayersReturnsCorrectOrder(t *testing.T) {
	expect := 4
	gamers := getFourGamers()
	for i := 3; i >= 0; i-- {
		gamers[i].turnOrder = (i - 4) * -1
	}
	res := sortPlayers(gamers)
	if len(res) != expect {
		t.Errorf("Did not received expected number of players back. Expected %v, got %v", expect, len(res))
	}
	if gamers[0].name != "John Deacon" {
		t.Errorf("Expected John Deacon to be first")
	}
	if gamers[1].name != "Roger Taylor" {
		t.Errorf("Expected Roger Taylor to be second")
	}
	if gamers[2].name != "Brian May" {
		t.Errorf("Expected Brian May to be third")
	}
	if gamers[3].name != "Freddie Mercury" {
		t.Errorf("Expected Freddie Mercury to be last")
	}

}

func TestGetWinnersReturnsOne(t *testing.T) {
	gamers := getFourGamers()
	winners := getWinners(gamers)
	if len(winners) != 1 {
		t.Errorf("Did not received one winner")
	}
	if winners[0].name != "Freddie Mercury" {
		t.Errorf("Freddie Mercury did not win.")
	}
}

func TestGetWinnersReturnsTie(t *testing.T) {
	gamers := getFourGamers()
	gamers[2].score = 0
	winners := getWinners(gamers)
	if len(winners) != 2 {
		t.Errorf("Did not received four players back")
	}
	if winners[0].name != "Freddie Mercury" {
		t.Errorf("Freddie Mercury did not tie for win.")
	}
	if winners[1].name != "Roger Taylor" {
		t.Errorf("Roger Taylor did not tie for win.")
	}
}

// helper
func getFourGamers() []player {
	queen := []string{"Freddie Mercury", "Brian May", "Roger Taylor", "John Deacon"}
	var gamers []player
	for i, name := range queen {
		gamer := player{
			name,
			0,
			i,
		}
		gamers = append(gamers, gamer)
	}

	return gamers
}
