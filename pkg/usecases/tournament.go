package usecases

import (
	"tournament/pkg/domain"
)

const defaultPrize = 4000

// CreateTournament inserts new tournament instance in Repository.
func (c *Controller) CreateTournament(name string, deposit int) (uint64, error) {
	id, err := c.Repository.InsertTournament(name, deposit, defaultPrize)
	return id, err
}

// GetTournament returns tournament instance with ID from Repository.
func (c *Controller) GetTournament(id uint64) (domain.Tournament, error) {
	t, err := c.Repository.GetTournamentByID(id)
	return t, err
}

// DeleteTournament deletes tournament instance with ID from Repository.
func (c *Controller) DeleteTournament(id uint64) error {
	err := c.Repository.DeleteTournamentByID(id)
	return err
}

func (c *Controller) checkUserIsParticipant(u domain.User, t domain.Tournament) error {
	for _, p := range t.Participants {
		if p == u.ID {
			return ErrParticipantExists
		}
	}

	return nil
}

// JoinTournament assigns new participant to the tournament with ID
// and updating balance of the participant.
func (c *Controller) JoinTournament(tournamentID uint64, userID uint64) error {
	t, err := c.Repository.GetTournamentByID(tournamentID)
	if err != nil {
		return err
	}

	u, err := c.Repository.GetUserByID(userID)
	if err != nil {
		return err
	}

	err = c.checkUserIsParticipant(u, t)
	if err != nil {
		return err
	}

	if u.Balance < t.Deposit {
		return ErrNotEnoughPoints
	}

	err = c.Repository.AddUserInTournament(u.ID, t.ID)
	if err != nil {
		return err
	}

	u.Balance -= t.Deposit

	err = c.Repository.UpdateUserBalanceByID(u.ID, u.Balance)
	return err
}

func (c *Controller) findWinner(participantIDs []uint64) (domain.User, error) {
	users, err := c.Repository.GetUsers()
	if err != nil {
		return domain.User{}, err
	}

	winner := domain.User{ID: 0}
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
	return winner, ErrNoParticipants
}

// FinishTournament updates winner field of the tournament with ID
// and adds prize to the winner's balance.
func (c *Controller) FinishTournament(tid uint64) error {
	t, err := c.Repository.GetTournamentByID(tid)
	if err != nil {
		return err
	}

	if t.WinnerID != 0 {
		return ErrFinishedTournament
	}

	winner, err := c.findWinner(t.Participants)
	if err != nil {
		return err
	}

	err = c.Repository.SetWinner(winner.ID, t.ID)
	if err != nil {
		return err
	}

	winner.Balance += t.Prize

	err = c.Repository.UpdateUserBalanceByID(winner.ID, winner.Balance)
	return err
}
