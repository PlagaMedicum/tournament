package usecases

import (
	uuid "github.com/satori/go.uuid"
	database "tournament/database/model"
	"tournament/env/errproc"
	app "tournament/pkg"
	"tournament/pkg/tournament/model"
	userm "tournament/pkg/user/model"
	user "tournament/pkg/user/usecases"
)

const (
	nullUUID     = "00000000-0000-0000-0000-000000000000"
	defaultPrize = 4000
)

type TournamentInterface interface {
	CreateTournament(model.Tournament) (uuid.UUID, error)
	GetTournament(uuid.UUID) (model.Tournament, error)
	DeleteTournament(uuid.UUID) error
	JoinTournament(uuid.UUID, uuid.UUID) error
	FinishTournament(uuid.UUID) error
}

type Tournament struct {
	model.Tournament
}

// CreateTournamentHandler inserts new tournament instance in database.
func (c *Tournament) CreateTournament(t model.Tournament) (uuid.UUID, error) {
	t.Prize = defaultPrize

	err := app.DB.Conn.QueryRow(`
			insert into tournaments (name, deposit, prize) values
				($1, $2, $3) returning id;
		`, t.Name, t.Deposit, t.Prize).Scan(t.ID)
	return t.ID, err
}

func getTournamentParticipants(id uuid.NullUUID) ([]uuid.UUID, error) {
	var d database.DB
	d.Connect()

	r, err := d.Conn.Query(`select userid from participants where id = $1;`,
		id.UUID)
	if err != nil {
		return nil, err
	}

	var plist []uuid.UUID
	for i := 0; r.Next(); i++ {
		var u uuid.UUID

		err := r.Scan(&u)
		if err != nil {
			return nil, err
		}

		plist = append(plist, u)
	}

	return plist, nil
}

func getTournaments() ([]model.Tournament, error) {
	rows, err := app.DB.Conn.Query(`
			select * from tournaments;
		`)
	if err != nil {
		return nil, err
	}

	var tournamentList []model.Tournament
	for rows.Next() {
		var tournament model.Tournament
		var id, wid uuid.NullUUID

		err := rows.Scan(&tournament.ID, &tournament.Name, &tournament.Deposit, &tournament.Prize, &id, &wid)
		if err != nil {
			return nil, err
		}

		if wid.Valid {
			tournament.WinnerID = wid.UUID
		}

		if id.Valid {
			tournament.Users, err = getTournamentParticipants(id)
			if err != nil {
				return nil, err
			}
		}
		tournamentList = append(tournamentList, tournament)
	}

	return tournamentList, nil
}

// GetTournamentHandler returns tournament instance with ID from
// slice list of all tournaments in database.
func (c *Tournament) GetTournament(id uuid.UUID) (model.Tournament, error) {
	tournamentList, err := getTournaments()
	if err != nil {
		return model.Tournament{}, err
	}

	for _, tournament := range tournamentList {
		if tournament.ID == id {
			return tournament, nil
		}
	}

	return model.Tournament{}, errproc.NoTournamentWithID
}

// DeleteTournamentHandler deletes tournament instance with ID from database.
func (c *Tournament) DeleteTournament(id uuid.UUID) error {
	tournamentList, err := getTournaments()
	if err != nil {
		return err
	}

	for _, tournament := range tournamentList {
		if tournament.ID != id {
			continue
		}

		_, err := app.DB.Conn.Exec(`
				delete from participants using tournaments 
					where tournament.id = $1 and participants.id = tournaments.users;
				delete from tournaments
					where id = $1;
			`, tournament.ID)
		return err
	}

	return errproc.NoTournamentWithID
}

func addUserInTournament(u userm.User, t model.Tournament) error {
	u.Balance -= t.Deposit
	var pid uuid.NullUUID

	err := app.DB.Conn.QueryRow(`select users from tournaments where id = $1;`,
		t.ID).Scan(&pid)
	if err != nil {

		return err
	}

	if pid.Valid {
		_, err = app.DB.Conn.Exec(`insert into participants (id, userid) values ($1, $2);`,
			pid, u.ID)
		if err != nil {
			return err
		}
	} else {
		_, err := app.DB.Conn.Exec(`insert into participants (userid) values ($1);`,
			u.ID)
		if err != nil {
			return err
		}

		_, err = app.DB.Conn.Exec(`
				update tournaments set users = participants.id from participants 
					where participants.userid = $1;
			`, u.ID)
		if err != nil {
			return err
		}
	}

	_, err = app.DB.Conn.Exec(`update users set balance = $2 where id = $1;`,
		u.ID, u.Balance)

	return err
}

// JoinTournamentHandler assigns new participant to the tournament
// and updating balance of the participant.
func (c *Tournament) JoinTournament(id uuid.UUID, userID uuid.UUID) error {
	tlist, err := getTournaments()
	if err != nil {
		return err
	}

	for _, t := range tlist {
		if t.ID != id {
			continue
		}

		usr := user.User{}

		ulist, err := usr.GetUsers()
		if err != nil {
			return err
		}

		for _, u := range ulist{
			if u.ID != userID {
				continue
			}

			if u.Balance >= t.Deposit {
				err := addUserInTournament(u, t)
				return err
			} else {
				return errproc.NotEnoughPoints
			}
		}

		return errproc.NoUserWithID
	}

	return errproc.NoTournamentWithID
}

func findWinner(pIDList []uuid.UUID)  (userm.User, error) {
	usr := user.User{}

	uList, err := usr.GetUsers()
	if err != nil {
		return userm.User{}, err
	}

	var winner userm.User
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

func setWinner(w userm.User, t model.Tournament) error {
	w.Balance += t.Prize

	_, err := app.DB.Conn.Exec(`update tournaments set winner = $1 where id = $2;`,
		w.ID, t.ID)
	if err != nil {
		return err
	}

	_, err = app.DB.Conn.Exec(`update users set balance = $2 where id = $1;`,
		w.ID, w.Balance)
	return err
}

// FinishTournamentHandler updates winner field of the tournament.
// Adds prize to the winner's balance.
func (c *Tournament) FinishTournament(id uuid.UUID) error {
	tlist, err := getTournaments()
	if err != nil {
		return err
	}

	for _, t := range tlist {
		if t.ID != id {
			continue
		}

		if t.WinnerID.String() == nullUUID {
			winner, err := findWinner(c.GetParticipants())
			if err != nil {
				return err
			}

			err = setWinner(winner, t)
			return err
		}
		break
	}

	return errproc.NoTournamentWithID
}
