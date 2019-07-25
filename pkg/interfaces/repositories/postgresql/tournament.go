package postgresql

import (
	"github.com/jackc/pgx"
	"tournament/pkg/domain"
	"tournament/pkg/infrastructure/postgresqlDB"
)

func (c *PSQLController) InsertTournament(name string, deposit int, prize int) (string, error){
	var id string
	err := c.Handler.QueryRow(`
			insert into tournaments (name, deposit, prize) values
				($1, $2, $3) returning id;
		`, name, deposit, prize).Scan(id)
	return id, err
}

func (c *PSQLController) scanTournamentRow(row *pgx.Rows) (domain.Tournament, error) {
	t := domain.Tournament{}
	participantsID := c.IDType.NewNullable()
	winnerID := c.IDType.NewNullable()

	err := row.Scan(&t.ID, &t.Name, &t.Deposit, &t.Prize, &participantsID, &winnerID)
	if err != nil {
		return t, err
	}

	if winnerID.IsValid() {
		t.WinnerID = winnerID.String()
	}

	if participantsID.IsValid() {
		t.Participants, err = c.GetTournamentParticipantList(participantsID.String())
	}

	return t, err
}

func (c *PSQLController) GetTournaments() ([]domain.Tournament, error) {
	rows, err := c.Handler.Query(`
			select * from tournaments;
		`)
	if err != nil {
		return nil, err
	}

	var tournamentList []domain.Tournament
	for rows.Next() {
		t, err := c.scanTournamentRow(rows)
		if err != nil {
			return nil, err
		}

		tournamentList = append(tournamentList, t)
	}

	return tournamentList, nil
}

func (c *PSQLController) GetTournamentByID(id string) (domain.Tournament, error){
	row, err := c.Handler.Query(`
			select * from tournaments where id = $1;
		`, id)
	if err != nil {
		return domain.Tournament{}, err
	}

	t, err := c.scanTournamentRow(row)
	return t, err
}

func (c *PSQLController) GetTournamentParticipantList(id string) ([]string, error) {
	var d postgresqlDB.DB
	d.Connect()

	r, err := d.Conn.Query(`select userid from participants where id = $1;`,
		id)
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

func (c *PSQLController) DeleteTournamentByID(id string) error {
	_, err := c.Handler.Exec(`
				delete from participants using tournaments 
					where tournament.id = $1 and participants.id = tournaments.users;
				delete from tournaments
					where id = $1;
			`, id)
	return err
}

func (c *PSQLController) GetTournamentParticipants(id string) (NullableID, error) {
	pid := c.IDType.NewNullable()

	err := c.Handler.QueryRow(`select users from tournaments where id = $1;`,
		id).Scan(&pid)
	return pid, err
}

func (c *PSQLController) AddUserInTournament(userID, tournamentID string) error {
	pid, err := c.GetTournamentParticipants(tournamentID)
	if err != nil {
		return err
	}

	if pid.IsValid() {
		_, err = c.Handler.Exec(`insert into participants (id, userid) values ($1, $2);`,
			pid, userID)
		return err
	}

	_, err = c.Handler.Exec(`insert into participants (userid) values ($1);`,
		userID)
	if err != nil {
		return err
	}

	_, err = c.Handler.Exec(`
				update tournaments set users = participants.id from participants 
					where participants.userid = $1 and id = $2;
			`, userID, tournamentID)
	return err
}

func (c *PSQLController) SetWinner(winnerID, tournamentID string) error {
	_, err := c.Handler.Exec(`update tournaments set winner = $1 where id = $2;`,
		winnerID, tournamentID)
	return err
}
