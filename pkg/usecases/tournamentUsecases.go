package usecases

import (
	"tournament/pkg/domain"
)

const defaultPrize   = 4000

// CreateTournament inserts new tournament instance in Repository.
func (c *Controller) CreateTournament(name string, deposit int) (string, error) {
	id, err := c.Repository.InsertTournament(name, deposit, defaultPrize)
	return id, err
}

// GetTournament returns tournament instance with ID from Repository.
func (c *Controller) GetTournament(id string) (domain.Tournament, error) {
	t, err := c.Repository.GetTournamentByID(id)
	return t, err
}

// DeleteTournament deletes tournament instance with ID from Repository.
func (c *Controller) DeleteTournament(id string) error {
	t, err := c.Repository.GetTournamentByID(id)
	if err != nil {
		return err
	}

	err = c.Repository.DeleteTournamentByID(t.ID)
	return err
}

func (c *Controller) checkUserIsParticipant(u domain.User, t domain.Tournament) error {
	participants, err := c.Repository.GetTournamentParticipantList(t.ID)
	if err != nil {
		return err
	}

	for _, participantID := range participants {
		if participantID == u.ID{
			return ErrParticipantExists
		}
	}

	return nil
}

// JoinTournament assigns new participant to the tournament with id
// and updating balance of the participant.
func (c *Controller) JoinTournament(tournamentID string, userID string) error {
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

	if u.Balance >= t.Deposit {
		err := c.Repository.AddUserInTournament(u.ID, t.ID)
		if err != nil {
			return err
		}

		u.Balance -= t.Deposit
		err = c.Repository.UpdateUserBalanceByID(u.Balance, u.ID)
		return err
	}
	return ErrNotEnoughPoints
}

func (c *Controller) findWinner(pIDList []string)  (domain.User, error) {
	uList, err := c.Repository.GetUsers()
	if err != nil {
		return domain.User{}, err
	}

	var winner domain.User
	for _, pid := range pIDList {
		for _, u := range uList {
			if u.ID == pid {
				if u.Balance > winner.Balance {
					winner = u
				}

				break
			}
		}
	}

	return winner, nil
}

// FinishTournament updates winner field of the tournament with id.
// Adds prize to the winner's balance.
func (c *Controller) FinishTournament(id string) error {
	t, err := c.Repository.GetTournamentByID(id)
	if err != nil {
		return err
	}

	if t.WinnerID == c.IDType.Null() {
		winner, err := c.findWinner(t.GetParticipants())
		if err != nil {
			return err
		}

		winner.Balance += t.Prize

		err = c.Repository.SetWinner(winner.ID, t.ID)
		if err != nil {
			return err
		}

		err = c.Repository.UpdateUserBalanceByID(winner.Balance, winner.ID)
		return err
	}
	return ErrFinishedTournament
}
