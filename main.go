package main

import "fmt"

const wild = 4
const wildValue = 0

func main() {
	fmt.Println(" -- game is starting -- ")
	fourNonBlondes := []string{"Christa Hillhouse", "Shaunna Hall", "Wanda Day", "Linda Perry"}
	totalDice := 5
	_, err := playGame(fourNonBlondes, totalDice)
	if err != nil {
		fmt.Println("error during game.", err.Error())
	}

	fmt.Println(" -- game over man -- ")
}