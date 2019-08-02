package tournament

import (
	"tournament/pkg/domain/tournament"
	"tournament/pkg/domain/user"
	"tournament/pkg/usecases"
)

const defaultPrize = 4000

// CreateTournament inserts new tournament instance in Repository.
func (c *Controller) CreateTournament(name string, deposit int) (uint64, error) {
	id, err := c.InsertTournament(name, deposit, defaultPrize)
	return id, err
}

// GetTournament returns tournament instance with ID from Repository.
func (c *Controller) GetTournament(id uint64) (tournament.Tournament, error) {
	t, err := c.GetTournamentByID(id)
	return t, err
}

// DeleteTournament deletes tournament instance with ID from Repository.
func (c *Controller) DeleteTournament(id uint64) error {
	err := c.DeleteTournamentByID(id)
	return err
}

func (c *Controller) checkUserIsParticipant(u user.User, t tournament.Tournament) error {
	for _, p := range t.Participants {
		if p == u.ID {
			return usecases.ErrParticipantExists
		}
	}

	return nil
}

// JoinTournament assigns new participant to the tournament with ID
// and updating balance of the participant.
func (c *Controller) JoinTournament(tournamentID uint64, userID uint64) error {
	t, err := c.GetTournamentByID(tournamentID)
	if err != nil {
		return err
	}

	u, err := c.GetUserByID(userID)
	if err != nil {
		return err
	}

	err = c.checkUserIsParticipant(u, t)
	if err != nil {
		return err
	}

	if u.Balance < t.Deposit {
		return usecases.ErrNotEnoughPoints
	}

	err = c.AddUserInTournament(u.ID, t.ID)
	if err != nil {
		return err
	}

	u.Balance -= t.Deposit

	err = c.UpdateUserBalanceByID(u.ID, u.Balance)
	return err
}

func (c *Controller) findWinner(participantIDs []uint64) (user.User, error) {
	users, err := c.GetUsers()
	if err != nil {
		return user.User{}, err
	}

	winner := user.User{ID: 0}
	for _, pid := range participantIDs {
		for _, u := range users {
			if (u.ID != pid) || (u.Balance < winner.Balance) {
				continue
			}

			winner = u
			break
		}
	}

	if winner.ID != 0 {
		return winner, nil
	}
	return winner, usecases.ErrNoParticipants
}

// FinishTournament updates winner field of the tournament with ID
// and adds prize to the winner's balance.
func (c *Controller) FinishTournament(tid uint64) error {
	t, err := c.GetTournamentByID(tid)
	if err != nil {
		return err
	}

	if t.WinnerID != 0 {
		return usecases.ErrFinishedTournament
	}

	winner, err := c.findWinner(t.Participants)
	if err != nil {
		return err
	}

	err = c.SetWinner(winner.ID, t.ID)
	if err != nil {
		return err
	}

	winner.Balance += t.Prize

	err = c.UpdateUserBalanceByID(winner.ID, winner.Balance)
	return err
}
