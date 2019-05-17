# simple_dice_game

A simple simulation of players participating in a dice game with the below rules. This is a simple simulation of 
players, so game trusts players to report rolls. Later version might decouple game from simulation, use channel to 
wait for turn. Simulation could be cli, api, or script. Give players more logic that changes with scores (more risk, 
less risk.)

## Rounds
* Players take 1 turn per round
* First Round: start with a random first player, then other players in any order
* Other Rounds: Each player goes first once, game ends when each player has gone first once

## Turns
* On each player's turn they roll five dice
* They must choose at least one die to keep
* They may choose more dice to keep
* Score is face value, except 4's are counted as 0
* Score accumulates through rounds

## Goals
* Each player is seeking the lowest total score
* Simple logic: players keep 1s, 2s, and 4s


