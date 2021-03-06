package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const noDiceErr = "no dice provided"

type player struct {
	name      string
	turnOrder int
	score     int
}

func (p *player) takeTurn(numDice int) error {
	if numDice == 0 {
		return errors.New(noDiceErr)
	}

	fmt.Printf("\n%s is taking their turn.\n", p.name)
	rand.Seed(time.Now().UnixNano()) // seed once per turn
	diceRemaining := numDice
	for diceRemaining > 0 {
		dice := p.roll(diceRemaining)
		chosen, err := p.choose(dice)
		if err != nil {
			return errors.New(err.Error())
		}

		total := p.totalChoices(chosen)
		p.score += total
		diceRemaining -= len(chosen)
		fmt.Printf("  … score goes up %d (total: %d)\n", total, p.score)
	}

	return nil
}

func (p *player) roll(d int) []int {
	var dice []int
	for d > 0 {
		rnd := rand.Intn(5) + 1
		dice = append(dice, rnd)
		d--
	}
	return dice
}

func (p *player) choose(dice []int) (choices []int, err error) {
	if len(dice) == 0 {
		err = errors.New(noDiceErr)
		return
	}

	p.reportRolls(dice)
	for _, die := range dice {
		if die == wild || die == 1 || die == 2 {
			choices = append(choices, die)
			fmt.Printf("  … keeps: %d\n", die)
		}
	}
	if len(choices) == 0 {
		var low int
		for i := 0; i < len(choices); i++ {
			// first iteration is always first low
			if dice[i] < dice[low] || i == 0 {
				low = i
			}
		}
		choices = append(choices, dice[low])
		fmt.Printf("  … keeps: %d\n", dice[low])
	}
	return
}

func (p *player) reportRolls(dice []int) {
	fmt.Printf("%s rolls: ", p.name)
	for _, d := range dice {
		fmt.Printf("%d ", d)
	}
	fmt.Printf("\n")
}

func (p *player) totalChoices(chosen []int) (total int) {
	for _, d := range chosen {
		if d == wild {
			total += wildValue
		} else {
			total += d
		}
	}

	return
}
