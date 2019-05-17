package main

import "testing"

// integration - ensure no panic
func TestTakeTurnNoPanic(t *testing.T) {
	gamer := getPlayer()
	totalDice := 7
	err := gamer.takeTurn(totalDice)
	if err != nil {
		t.Errorf("Unexpected error while taking turn.")
	}
}

func TestTakeTurnErrWithNoDice(t *testing.T) {
	expect := noDiceErr
	gamer := getPlayer()
	err := gamer.takeTurn(0)
	if err == nil {
		t.Errorf("Did not get expected error while taking turn.")
	} else if err.Error() != expect {
		t.Errorf("Did not get expected error while taking turn. Expected: %v, got: %v", expect, err.Error())
	}
}

func TestChooseReturnsCorrectChoices(t *testing.T) {
	gamer := getPlayer()
	testTables := []struct {
		input  []int
		expect []int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, []int{1, 2, 4}},
		{[]int{3, 5, 6}, []int{3}},
		{[]int{6}, []int{6}},
	}

	for _, table := range testTables {
		res, err := gamer.choose(table.input)
		if err != nil {
			t.Errorf("Unexpexted error while choosing dice.")
		}
		if intArrayEquals(res, table.expect) == false {
			t.Errorf("Did not choose wisely. Wanted %v and Got %v", table.expect, res)
		}
	}
}

func TestChooseErrWithNoDice(t *testing.T) {
	expect := noDiceErr
	gamer := getPlayer()
	var empty []int
	_, err := gamer.choose(empty)
	if err == nil {
		t.Errorf("Did not get expected error while taking turn.")
	} else if err.Error() != expect {
		t.Errorf("Did not get expected error while taking turn. Expected: %v, got: %v", expect, err.Error())
	}

}

func TestTotalChoicesDoesCorrectMath(t *testing.T) {
	gamer := getPlayer()
	testTables := []struct {
		input  []int
		expect int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, 17}, // 4 is 0
		{[]int{3, 5, 6}, 14},
		{[]int{6, 5, 3}, 14},
		{[]int{}, 0},
	}
	for _, table := range testTables {
		res := gamer.totalChoices(table.input)
		if res != table.expect {
			t.Errorf("Did not total correctly. Wanted %v and Got %v", table.expect, res)
		}
	}
}

// helper
func getPlayer() player {
	player := player{
		"Freddie Mercury",
		1,
		10,
	}
	return player
}

// helper
func intArrayEquals(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
