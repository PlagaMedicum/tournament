package tournament

import "tournament/pkg/domain/tournament"

// Usecases is the interface for tournament usecases.
type Usecases interface {
	CreateTournament(string, int) (uint64, error)
	GetTournament(uint64) (tournament.Tournament, error)
	DeleteTournament(uint64) error
	JoinTournament(uint64, uint64) error
	FinishTournament(uint64) error
}
