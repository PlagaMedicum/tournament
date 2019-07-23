package usecases

import (
	"os/user"
	"tournament/pkg/domain"
	"tournament/pkg/infrastructure/errproc"
)

const (
	defaultPrize = 4000
)

const (
	defaultBalance = 700
)

type Repository interface {
	InsertUser(domain.User) (string, error)
	GetTournaments() ([]domain.Tournament, error)
	GetTournamentParticipants(string) ([]string, error)
	SetWinner(domain.User, domain.Tournament) error
}

type RepositoryInteractor interface {

	CreateTournament(domain.Tournament) (string, error)
	GetTournament(string) (domain.Tournament, error)
	DeleteTournament(string) error
	JoinTournament(string, string) error
	FinishTournament(string) error
	CreateUser(domain.User) (string, error)
	GetUsers() ([]domain.User, error)
	GetUser(string) (domain.User, error)
	DeleteUser(string) error
	FundUser(string, int) error
}

type Controller struct {
	r Repository
}

// CreateUser inserts new user instance in postgresqlDB.
func (c *Controller) CreateUser(u domain.User) (string, error) {
	u.Balance = defaultBalance

	err := app.DB.Conn.QueryRow(`
			insert into users (name, balance) values 
				($1, $2) returning id;
		`, u.Name, u.Balance).Scan(u.ID)
	return u.ID.String(), err
}

// GetUsers returns a slice of all users in the postgresqlDB.
func (c *Controller) GetUsers() ([]domain.User, error) {
	rows, err := app.DB.Conn.Query(`select * from users;`)
	if err != nil {
		return nil, err
	}

	var userList []domain.User
	for rows.Next() {
		var user domain.User

		err := rows.Scan(&user.ID, &user.Name, &user.Balance)
		if err != nil {
			return nil, err
		}

		userList = append(userList, user)
	}

	return userList, nil
}

// GetUser returns user instance with ID from
// slice list of all users in postgresqlDB.
func (c *Controller) GetUser(id string) (domain.User, error) {
	userList, err := c.GetUsers()
	if err != nil {
		return domain.User{}, err
	}

	for _, u := range userList {
		if u.ID.String() == id {
			return u, nil
		}
	}

	return domain.User{}, errproc.NoUserWithID
}

// DeleteUser deletes user instance with ID from postgresqlDB.
func (c *Controller) DeleteUser(id string) error {
	userList, err := c.GetUsers()
	if err != nil {
		return err
	}

	for index := range userList {
		if userList[index].ID.String() != id {
			continue
		}

		_, err := app.DB.Conn.Exec(`delete from users where id = $1;`,
			id)
		return err
	}

	return errproc.NoUserWithID
}

// FundUser adds points to balance of user's with the ID.
func (c *Controller) FundUser(id string, points int) error {
	userList, err := c.GetUsers()
	if err != nil {
		return err
	}

	for index := range userList {
		if userList[index].ID.String() != id {
			continue
		}

		userList[index].Balance += points

		_, err := app.DB.Conn.Exec(`update users set balance = $1 where id = $2;`,
			userList[index].Balance, id)
		return err
	}

	return errproc.NoUserWithID
}

// CreateTournament inserts new tournament instance in postgresqlDB.
func (c *Controller) CreateTournament(t domain.Tournament) (string, error) {
	t.Prize = defaultPrize

	err := app.DB.Conn.QueryRow(`
			insert into tournaments (name, deposit, prize) values
				($1, $2, $3) returning id;
		`, t.Name, t.Deposit, t.Prize).Scan(t.ID)
	return t.ID.String(), err
}

// GetTournament returns tournament instance with ID from
// slice list of all tournaments in postgresqlDB.
func (c *Controller) GetTournament(id string) (domain.Tournament, error) {
	tournamentList, err := c.r.GetTournaments()
	if err != nil {
		return domain.Tournament{}, err
	}

	for _, tournament := range tournamentList {
		if tournament.ID.String() == id {
			return tournament, nil
		}
	}

	return domain.Tournament{}, errproc.NoTournamentWithID
}

// DeleteTournament deletes tournament instance with ID from postgresqlDB.
func (c *Controller) DeleteTournament(id string) error {
	tournamentList, err := getTournaments()
	if err != nil {
		return err
	}

	for _, tournament := range tournamentList {
		if tournament.ID.String() != id {
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

func (c *Controller) addUserInTournament(u domain.User, t domain.Tournament) error {
	u.Balance -= t.Deposit
	var pid string

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

// JoinTournament assigns new participant to the tournament
// and updating balance of the participant.
func (c *Controller) JoinTournament(id string, userID string) error {
	tlist, err := c.r.GetTournaments()
	if err != nil {
		return err
	}

	for _, t := range tlist {
		if t.ID.String() != id {
			continue
		}

		usr := user.User{}

		ulist, err := c.GetUsers()
		if err != nil {
			return err
		}

		for _, u := range ulist{
			if u.ID.String() != userID {
				continue
			}

			if u.Balance >= t.Deposit {
				err := c.addUserInTournament(u, t)
				return err
			} else {
				return errproc.NotEnoughPoints
			}
		}

		return errproc.NoUserWithID
	}

	return errproc.NoTournamentWithID
}

func (c *Controller) findWinner(pIDList []string)  (domain.User, error) {
	usr := domain.User{}

	uList, err := c.GetUsers()
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

// FinishTournament updates winner field of the tournament.
// Adds prize to the winner's balance.
func (c *Controller) FinishTournament(id string) error {
	tlist, err := c.r.GetTournaments()
	if err != nil {
		return err
	}

	for _, t := range tlist {
		if t.ID.String() != id {
			continue
		}

		if t.WinnerID == id.Null() {
			winner, err := c.findWinner(t.GetParticipants())
			if err != nil {
				return err
			}

			winner.Balance += t.Prize
			err = c.r.SetWinner(winner, t)
			return err
		}
		break
	}

	return errproc.NoTournamentWithID
}
