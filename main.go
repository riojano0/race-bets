package main

import (
	"math/rand"

	"github.com/riojano0/race-bets/challenger"
	"github.com/riojano0/race-bets/pilot"
	"github.com/riojano0/race-bets/racetrack"
	"syreclabs.com/go/faker"
)

func main() {
	challengers := []challenger.Challenger{}

	for i := 0; i < 12; i++ {
		var challenger challenger.Challenger = pilot.CreateNewPilot(faker.Name().Name())
		challengers = append(challengers, challenger)
	}

	//Shufle the challengers
	for i := range challengers {
		j := rand.Intn(i + 1)
		challengers[i], challengers[j] = challengers[j], challengers[i]
	}

	raceTrack := racetrack.CreateRaceTrack(challengers)

	raceTrack.StartRace()
}
