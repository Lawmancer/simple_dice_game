package main

import "fmt"

func main() {
	fmt.Println(" -- game is starting -- ")
	playerNames := []string{"Christa Hillhouse", "Shaunna Hall", "Wanda Day", "Linda Perry"}
	_, err := playGame(playerNames)
	if err != nil {
		fmt.Println("error during game.", err.Error())
	}
	fmt.Println(" -- game over man -- ")
}
