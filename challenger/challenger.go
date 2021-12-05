package challenger

import "sync"

type Challenger interface {
	Run(trackLength int, wg *sync.WaitGroup, challengerChannel chan<- Challenger)

	GetInformation() string

	GetChallengerNumber() int
}
