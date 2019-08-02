package tournament

import (
	"database/sql"
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

func (c *Controller) getTournamentParticipantList(pid uint64) ([]uint64, error) {
	r, err := c.Query(`select userid from participants where id = $1;`,
		pid)
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

// GetTournamentByID returns from DB the tournament with tid.
func (c *Controller) GetTournamentByID(tid uint64) (tournament.Tournament, error) {
	t := tournament.Tournament{}
	pid := sql.NullInt64{}
	wid := sql.NullInt64{}

	err := c.QueryRow(`select * from tournaments where id = $1;`,
		tid).Scan(&t.ID, &t.Name, &t.Deposit, &t.Prize, &pid, &wid)
	if err != nil {
		return tournament.Tournament{}, err
	}

	if pid.Valid {
		t.Participants, err = c.getTournamentParticipantList(uint64(pid.Int64))
	}

	if wid.Valid {
		t.WinnerID = uint64(wid.Int64)
	}

	return t, err
}

// DeleteTournamentByID deletes from DB the tournament with tid.
func (c *Controller) DeleteTournamentByID(tid uint64) error {
	_, err := c.Exec(`
			delete from participants using tournaments 
				where tournaments.id = $1 and participants.id = tournaments.users;
		`, tid)
	if err != nil {
		return err
	}

	_, err = c.Exec(`delete from tournaments where id = $1;`,
		tid)
	return err
}

func (c *Controller) getTournamentParticipants(tid uint64) (uint64, error) {
	var pid sql.NullInt64

	err := c.QueryRow(`select users from tournaments where id = $1;`,
		tid).Scan(&pid)
	return uint64(pid.Int64), err
}

func (c *Controller) insertParticipant(pid, uid uint64) (error, uint64) {
	if pid != 0 {
		_, err := c.Exec(`insert into participants (id, userid) values ($1, $2);`,
			pid, uid)
		return err, pid
	}

	err := c.QueryRow(`insert into participants (userid) values ($1) returning id;`,
		uid).Scan(&pid)
	return err, pid
}

func (c *Controller) updateTournamentParticipants(pid, tid uint64) error {
	_, err := c.Exec(`update tournaments set users = $1 where id = $2;`,
		pid, tid)
	return err
}

// AddUserInTournament adds the user with uid in participant of
// tournament with tid.
func (c *Controller) AddUserInTournament(uid, tid uint64) error {
	pid, err := c.getTournamentParticipants(tid)
	if err != nil {
		return err
	}

	err, pid = c.insertParticipant(pid, uid)
	if err != nil {
		return err
	}

	err = c.updateTournamentParticipants(pid, tid)
	return err
}

// SetWinner sets the user with wid a winner of tournament with tid.
func (c *Controller) SetWinner(wid, tid uint64) error {
	_, err := c.Exec(`update tournaments set winner = $1 where id = $2;`,
		wid, tid)
	return err
}
