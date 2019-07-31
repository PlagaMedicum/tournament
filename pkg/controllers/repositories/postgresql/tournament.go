package postgresql

import (
	"tournament/pkg/domain"
)

func (c *PSQLController) InsertTournament(name string, deposit int, prize int) (string, error) {
	id := c.IDFactory.New()

	err := c.Handler.QueryRow(`
			insert into tournaments (name, deposit, prize) values
				($1, $2, $3) returning id;
		`, name, deposit, prize).Scan(id)
	return id.String(), err
}

func (c *PSQLController) getTournamentParticipantList(pid string) ([]string, error) {
	id := c.IDFactory.NewNullable()

	err := id.UnmarshalText([]byte(pid))
	if err != nil {
		return nil, err
	}

	r, err := c.Handler.Query(`select userid from participants where id = $1;`,
		id.GetPointer())
	if err != nil {
		return nil, err
	}

	var plist []string
	for r.Next() {
		uid := c.IDFactory.New()

		err := r.Scan(uid)
		if err != nil {
			return nil, err
		}

		plist = append(plist, uid.String())
	}

	return plist, nil
}

func (c *PSQLController) GetTournamentByID(tid string) (domain.Tournament, error) {
	id := c.IDFactory.New()
	pid := c.IDFactory.NewNullable()
	wid := c.IDFactory.NewNullable()
	t := domain.Tournament{}

	err := id.UnmarshalText([]byte(tid))
	if err != nil {
		return domain.Tournament{}, err
	}

	err = c.Handler.QueryRow(`select * from tournaments where id = $1;`,
		id).Scan(id, &t.Name, &t.Deposit, &t.Prize, pid.GetPointer(), wid.GetPointer())
	if err != nil {
		return domain.Tournament{}, err
	}

	t.ID = id.String()
	t.WinnerID = wid.String()

	if pid.IsValid() {
		t.Participants, err = c.getTournamentParticipantList(pid.String())
	}

	return t, err
}

func (c *PSQLController) DeleteTournamentByID(tid string) error {
	id := c.IDFactory.New()

	err := id.UnmarshalText([]byte(tid))
	if err != nil {
		return err
	}

	_, err = c.Handler.Exec(`
			delete from participants using tournaments 
				where tournaments.id = $1 and participants.id = tournaments.users;
		`, id)
	if err != nil {
		return err
	}

	_, err = c.Handler.Exec(`delete from tournaments where id = $1;`,
		id)
	return err
}

func (c *PSQLController) getTournamentParticipants(tournamentID string) (NullableID, error) {
	tid := c.IDFactory.New()
	pid := c.IDFactory.NewNullable()

	err := tid.UnmarshalText([]byte(tournamentID))
	if err != nil {
		return pid, err
	}

	err = c.Handler.QueryRow(`select users from tournaments where id = $1;`,
		tid).Scan(pid.GetPointer())
	return pid, err
}

func (c *PSQLController) insertParticipant(pid NullableID, userID string) (error, NullableID) {
	uid := c.IDFactory.New()

	err := uid.UnmarshalText([]byte(userID))
	if err != nil {
		return err, pid
	}

	if pid.IsValid() {
		_, err = c.Handler.Exec(`insert into participants (id, userid) values ($1, $2);`,
			pid.GetNotNullPointer(), uid)
		return err, pid
	}

	err = c.Handler.QueryRow(`insert into participants (userid) values ($1) returning id;`,
		uid).Scan(pid.GetNotNullPointer())
	return err, pid
}

func (c *PSQLController) updateTournamentParticipants(listID string, tournamentID string) error {
	pid := c.IDFactory.NewNullable()
	tid := c.IDFactory.New()

	err := pid.UnmarshalText([]byte(listID))
	if err != nil {
		return err
	}

	err = tid.UnmarshalText([]byte(tournamentID))
	if err != nil {
		return err
	}

	_, err = c.Handler.Exec(`update tournaments set users = $1 where id = $2;`,
		pid, tid)
	return err
}

func (c *PSQLController) AddUserInTournament(userID, tournamentID string) error {
	pid, err := c.getTournamentParticipants(tournamentID)
	if err != nil {
		return err
	}

	err, pid = c.insertParticipant(pid, userID)
	if err != nil {
		return err
	}

	err = c.updateTournamentParticipants(pid.String(), tournamentID)
	return err
}

func (c *PSQLController) SetWinner(winnerID, tournamentID string) error {
	wid := c.IDFactory.NewNullable()
	tid := c.IDFactory.New()

	err := wid.UnmarshalText([]byte(winnerID))
	if err != nil {
		return err
	}

	err = tid.UnmarshalText([]byte(tournamentID))
	if err != nil {
		return err
	}

	_, err = c.Handler.Exec(`update tournaments set winner = $1 where id = $2;`,
		wid.GetPointer(), tid)
	return err
}
