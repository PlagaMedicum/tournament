package repositories

import (
	uuid "github.com/satori/go.uuid"
	"tournament/pkg/domain"
	database "tournament/pkg/infrastructure/postgresqlDB/model"
)

type Controller struct {
	db database.DB
}

func InsertUser(u domain.User) (string, error) {

}

func (c *Controller) GetTournaments() ([]domain.Tournament, error) {
	rows, err := c.db.Conn.Query(`
			select * from tournaments;
		`)
	if err != nil {
		return nil, err
	}

	var tournamentList []domain.Tournament
	for rows.Next() {
		var tournament domain.Tournament
		var id, wid uuid.NullUUID

		err := rows.Scan(&tournament.ID, &tournament.Name, &tournament.Deposit, &tournament.Prize, &id, &wid)
		if err != nil {
			return nil, err
		}

		if wid.Valid {
			tournament.WinnerID = wid.UUID.String()
		}

		if id.Valid {
			tournament.Users, err = c.GetTournamentParticipants(id.UUID.String())
			if err != nil {
				return nil, err
			}
		}
		tournamentList = append(tournamentList, tournament)
	}

	return tournamentList, nil
}

func (c *Controller) GetTournamentParticipants(id string) ([]string, error) {
	var d database.DB
	d.Connect()

	r, err := d.Conn.Query(`select userid from participants where id = $1;`,
		id.UUID)
	if err != nil {
		return nil, err
	}

	var plist []string
	for i := 0; r.Next(); i++ {
		var u string

		err := r.Scan(&u)
		if err != nil {
			return nil, err
		}

		plist = append(plist, u)
	}

	return plist, nil
}

func (c *Controller) SetWinner(w domain.User, t domain.Tournament) error {
	_, err := c.db.Conn.Exec(`update tournaments set winner = $1 where id = $2;`,
		w.ID, t.ID)
	if err != nil {
		return err
	}

	_, err = c.db.Conn.Exec(`update users set balance = $2 where id = $1;`,
		w.ID, w.Balance)
	return err
}
