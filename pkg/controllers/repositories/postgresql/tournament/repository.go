package tournament

import (
	"tournament/pkg/controllers/repositories/postgresql"
	"tournament/pkg/domain/tournament"
	"tournament/pkg/domain/user"
)

func (c *Controller) scanUserRow(row postgresql.Rows) (user.User, error) {
	u := user.User{}

	err := row.Scan(&u.ID, &u.Name, &u.Balance)
	if err != nil {
		return u, err
	}

	return u, nil
}

// GetUsers returns a slice of all users from DB.
func (c *Controller) GetUsers() ([]user.User, error) {
	rows, err := c.Query(`select * from users;`)
	if err != nil {
		return nil, err
	}

	var users []user.User
	for rows.Next() {
		u, err := c.scanUserRow(rows)
		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}

// GetUserByID returns from DB the user with uid.
func (c *Controller) GetUserByID(uid uint64) (user.User, error){
	u := user.User{}

	err := c.QueryRow(`select * from users where id = $1;`,
		uid).Scan(&u.ID, &u.Name, &u.Balance)
	if err != nil {
		return u, err
	}

	return u, nil
}

// UpdateUserBalanceByID updates balance of the user with uid in DB.
func (c *Controller) UpdateUserBalanceByID(uid uint64, balance int) error {
	_, err := c.Exec(`update users set balance = $1 where id = $2;`,
		balance, uid)
	return err
}

// InsertTournament creates in DB the new tournament with the name,
// deposit and prize.
func (c *Controller) InsertTournament(name string, deposit int, prize int) (uint64, error) {
	var id uint64
	
	err := c.QueryRow(`
			insert into tournaments (name, deposit, prize) values
				($1, $2, $3) returning id;
		`, name, deposit, prize).Scan(&id)
	return id, err
}

func (c *Controller) getTournamentParticipantList(tid uint64) ([]uint64, error) {
	r, err := c.Query(`select userid from participants where tournamentid = $1;`,
		tid)
	if err == c.ErrNoRows() {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var plist []uint64
	for r.Next() {
		var uid uint64

		err := r.Scan(&uid)
		if err != nil {
			return nil, err
		}

		plist = append(plist, uid)
	}

	return plist, nil
}

func (c *Controller) getTournamentWinner(tid uint64) (uint64, error) {
	var wid uint64

	err := c.QueryRow(`select userid from winners where tournamentid = $1;`,
		tid).Scan(&wid)
	if err == c.ErrNoRows() {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}

	return wid, nil
}

// GetTournamentByID returns from DB the tournament with tid.
func (c *Controller) GetTournamentByID(tid uint64) (tournament.Tournament, error) {
	t := tournament.Tournament{}

	err := c.QueryRow(`select * from tournaments where id = $1;`,
		tid).Scan(&t.ID, &t.Name, &t.Deposit, &t.Prize)
	if err != nil {
		return tournament.Tournament{}, err
	}

	t.Participants, err = c.getTournamentParticipantList(tid)
	if err != nil {
		return t, err
	}

	t.WinnerID, err = c.getTournamentWinner(tid)
	return t, err
}

// DeleteTournamentByID deletes from DB the tournament with tid.
func (c *Controller) DeleteTournamentByID(tid uint64) error {
	_, err := c.Exec(`
			delete from participants where participants.tournamentid = $1;
		`, tid)
	if err != nil {
		return err
	}

	_, err = c.Exec(`
			delete from winners where winners.tournamentid = $1;
		`, tid)
	if err != nil {
		return err
	}

	_, err = c.Exec(`delete from tournaments where id = $1;`,
		tid)
	return err
}

// AddUserInTournament adds the user with uid in participant of
// tournament with tid.
func (c *Controller) AddUserInTournament(uid, tid uint64) error {
	_, err := c.Exec(`insert into participants (tournamentid, userid) values ($1, $2);`,
		tid, uid)
	return err
}

// SetWinner sets the user with wid a winner of tournament with tid.
func (c *Controller) SetWinner(wid, tid uint64) error {
	_, err := c.Exec(`insert into winners (tournamentid, userid) values ($1, $2);`,
		tid, wid)
	return err
}
