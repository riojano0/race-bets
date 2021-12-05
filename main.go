package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/riojano0/race-bets/challenger"
	"github.com/riojano0/race-bets/pilot"
	"github.com/riojano0/race-bets/racetrack"
	"syreclabs.com/go/faker"
)

func main() {

	play := true
	var playInput string

	for play {
		startGame()
		fmt.Println("\nDo you want to play again? (y/n)")

		fmt.Scan(&playInput)
		if playInput == "y" {
			play = true
		} else {
			play = false
		}
	}

}

func startGame() {
	challengers := []challenger.Challenger{}

	for i := 0; i < 8; i++ {
		var challenger challenger.Challenger = pilot.CreateNewPilot(faker.Name().Name())
		challengers = append(challengers, challenger)
	}

	fmt.Println("****These are the challengers*****")
	for _, challenger := range challengers {
		fmt.Println(challenger.GetInformation())
	}

	fmt.Println("Select one by pilot number:")
	var selectChallenger int
	fmt.Scan(&selectChallenger)

	//Shufle the challengers
	rand.Seed(time.Now().UnixNano())
	for i := range challengers {
		j := rand.Intn(i + 1)
		challengers[i], challengers[j] = challengers[j], challengers[i]
	}

	fmt.Println("\n*********** Stand by ************")
	raceTrack := racetrack.CreateRaceTrack(challengers)

	winnerChannel := make(chan int)

	raceTrack.StartRace(winnerChannel)

	winner := <-winnerChannel
	if winner == selectChallenger {
		fmt.Printf("\n******You won! Congrats! Pilot in first position was the: %v******", winner)
	} else {
		fmt.Printf("\n******You lost! The pilot in first position was: %v******", winner)
	}

	pilot.ResetPilotNumbers()
}
