package racetrack

import (
	"fmt"
	"math/rand"
	"sync"

	"github.com/riojano0/race-bets/challenger"
)

type RaceTrack struct {
	Challengers []challenger.Challenger
	Track       Track
}

type Track struct {
	Length int
}

func createTrack() Track {
	min := 1000
	max := 10000
	trackLength := rand.Intn(max-min) + min
	return Track{
		Length: trackLength,
	}
}

func CreateRaceTrack(challengers []challenger.Challenger) *RaceTrack {
	track := createTrack()

	return &RaceTrack{
		Challengers: challengers,
		Track:       track,
	}
}

func (racetrack *RaceTrack) StartRace() {

	wg := &sync.WaitGroup{}
	challengerChannel := make(chan challenger.Challenger)

	fmt.Println("Starting Race on track with length:", racetrack.Track.Length)
	wg.Add(2)
	go challengersRun(racetrack, wg, challengerChannel)
	go AnnounceWinners(racetrack, wg, challengerChannel)

	wg.Wait()
}

func challengersRun(racetrack *RaceTrack, wg *sync.WaitGroup, challengerChannel chan challenger.Challenger) {

	wtChallenger := &sync.WaitGroup{}
	trackLength := racetrack.Track.Length

	for _, challenger := range racetrack.Challengers {
		wtChallenger.Add(1)
		go challenger.Run(trackLength, wtChallenger, challengerChannel)
	}

	wtChallenger.Wait()
	close(challengerChannel)

	wg.Done()
}

func AnnounceWinners(racetrack *RaceTrack, wg *sync.WaitGroup, challengerChannel <-chan challenger.Challenger) {

	positionSlotInuse := 0

	for challenger := range challengerChannel {
		positionSlotInuse += 1
		fmt.Println("Finish in position", positionSlotInuse, "\n\t", challenger.GetInformation())
	}

	wg.Done()
}
