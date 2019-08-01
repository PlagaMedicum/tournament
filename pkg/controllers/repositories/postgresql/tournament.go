package postgresql

import (
	"tournament/pkg/domain"
)

// InsertTournament creates in DB the new tournament with the name,
// deposit and prize.
func (c *PSQLController) InsertTournament(name string, deposit int, prize int) (uint64, error) {
	var id uint64
	
	err := c.Database.QueryRow(`
			insert into tournaments (name, deposit, prize) values
				($1, $2, $3) returning id;
		`, name, deposit, prize).Scan(&id)
	return id, err
}

func (c *PSQLController) getTournamentParticipantList(pid uint64) ([]uint64, error) {
	r, err := c.Database.Query(`select userid from participants where id = $1;`,
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
func (c *PSQLController) GetTournamentByID(tid uint64) (domain.Tournament, error) {
	t := domain.Tournament{}
	var pid uint64

	err := c.Database.QueryRow(`select * from tournaments where id = $1;`,
		tid).Scan(&t.ID, &t.Name, &t.Deposit, &t.Prize, &pid, &t.WinnerID)
	if err != nil {
		return domain.Tournament{}, err
	}

	if pid != 0 {
		t.Participants, err = c.getTournamentParticipantList(pid)
	}

	return t, err
}

// DeleteTournamentByID deletes from DB the tournament with tid.
func (c *PSQLController) DeleteTournamentByID(tid uint64) error {
	_, err := c.Database.Exec(`
			delete from participants using tournaments 
				where tournaments.id = $1 and participants.id = tournaments.users;
		`, tid)
	if err != nil {
		return err
	}

	_, err = c.Database.Exec(`delete from tournaments where id = $1;`,
		tid)
	return err
}

func (c *PSQLController) getTournamentParticipants(tid uint64) (uint64, error) {
	var pid uint64

	err := c.Database.QueryRow(`select users from tournaments where id = $1;`,
		tid).Scan(pid)
	return pid, err
}

func (c *PSQLController) insertParticipant(pid, uid uint64) (error, uint64) {
	if pid != 0 {
		_, err := c.Database.Exec(`insert into participants (id, userid) values ($1, $2);`,
			pid, uid)
		return err, pid
	}

	err := c.Database.QueryRow(`insert into participants (userid) values ($1) returning id;`,
		uid).Scan(&pid)
	return err, pid
}

func (c *PSQLController) updateTournamentParticipants(pid, tid uint64) error {
	_, err := c.Database.Exec(`update tournaments set users = $1 where id = $2;`,
		pid, tid)
	return err
}

// AddUserInTournament adds the user with uid in participant of
// tournament with tid.
func (c *PSQLController) AddUserInTournament(uid, tid uint64) error {
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
func (c *PSQLController) SetWinner(wid, tid uint64) error {
	_, err := c.Database.Exec(`update tournaments set winner = $1 where id = $2;`,
		wid, tid)
	return err
}
