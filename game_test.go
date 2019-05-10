package main

import (
	"testing"
)

func TestPlayGameReturnsPlayer(t *testing.T) {
	playerNames := []string{"Freddie Mercury", "Brian May", "Roger Taylor", "John Deacon "}
	_, err := playGame(playerNames)
	if err != nil {
		t.Errorf("Received unexpected error during play: %v", err)
	}
}

func TestMainErrWithoutNames(t *testing.T) {
	expect := "no players to play the game"
	var playerNames []string
	_, err := playGame(playerNames)
	if err == nil {
		t.Error("Did not receive expected error:")
	} else if err.Error() != expect {
		t.Errorf("Did not get expected error. Expected: %v, got: %v", expect, err.Error())
	}
}

func TestRandomFirstPlayerReturnsFour(t *testing.T) {
	expect := 4
	gamers := getFourGamers()
	res, err := randomFirstPlayer(gamers)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(res) != expect {
		t.Errorf("Did not received expected number of players back. Expected %v, got %v", expect, len(res))
	}
}

func TestRandomFirstPlayerErrWithZeroGamers(t *testing.T) {
	var gamers []player
	_, err := randomFirstPlayer(gamers)
	if err == nil {
		t.Errorf("Expected error")
	} else if err.Error() != "no gamers to sort" {
		t.Errorf("Expected error: %v", err)
	}
}

func TestSortPlayersReturnsFour(t *testing.T) {
	expect := 4
	gamers := getFourGamers()
	res := sortPlayers(gamers)
	if len(res) != expect {
		t.Errorf("Did not received expected number of players back. Expected %v, got %v", expect, len(res))
	}
}

func TestGetWinnerFirst(t *testing.T) {
	gamers := getFourGamers()
	sortedGamers := sortPlayers(gamers)
	if len(sortedGamers) != 4 {
		t.Errorf("Did not received four players back")
	}
	if sortedGamers[0].name != "Freddie Mercury" {
		t.Errorf("Freddie Mercury did not win.")
	}
}

func TestListPlayersDoesNotPanic(t *testing.T) {
	gamers := getFourGamers()
	listPlayers(gamers)
}

// helper
func getFourGamers() []player {
	playerNames := []string{"Freddie Mercury", "Brian May", "Roger Taylor", "John Deacon "}
	var gamers []player
	for i, name := range playerNames {
		gamer := player{
			name,
			0,
			i,
		}
		gamers = append(gamers, gamer)
	}

	return gamers
}
