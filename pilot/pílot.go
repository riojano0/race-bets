package pilot

import (
	"math/rand"
	"strconv"
	"sync"

	"github.com/riojano0/race-bets/challenger"
	"syreclabs.com/go/faker"
)

type Car struct {
	Velocity int
	Name     string
}

type Pilot struct {
	Car         Car
	Name        string
	PilotNumber int
}

var initialPilotNumber int = 0

func CreateNewPilot(name string) *Pilot {
	min := 50
	max := 100
	carVelocity := rand.Intn(max-min) + min

	initialPilotNumber += 1

	return &Pilot{
		Car: Car{
			Name:     faker.App().Name(),
			Velocity: carVelocity,
		},
		Name:        name,
		PilotNumber: initialPilotNumber,
	}
}

func (pilot *Pilot) Run(trackLength int, wg *sync.WaitGroup, challengerChannel chan<- challenger.Challenger) {

	lengthReach := 0
	for lengthReach < trackLength {
		// How is not 'parallel' is concurrent, the car with the major velocity not always is the first
		lengthReach += pilot.Car.Velocity
	}

	challengerChannel <- pilot

	wg.Done()
}

func (pilot *Pilot) GetInformation() string {
	return "Pilot Name: " + pilot.Name + " Pilot Number: " + strconv.Itoa(pilot.PilotNumber) + " Car Name: " + pilot.Car.Name
}
