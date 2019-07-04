package usecases

import (
	uuid "github.com/satori/go.uuid"
	app "tournament/pkg/boot"
	database "tournament/pkg/database/model"
	"tournament/pkg/errproc"
	"tournament/pkg/tournament/model"
	userm "tournament/pkg/user/model"
	user "tournament/pkg/user/usecases"
)

func CreateTournament(tournament *model.Tournament) {
	err := app.DB.Conn.QueryRow(`
			insert into tournaments (name, deposit, prize) values
				($1, $2, $3) returning id;
		`, tournament.Name, tournament.Deposit, tournament.Prize).Scan(&tournament.Id)
	errproc.HandleSQLErr("create tournament", err)
}

func GetTournaments() (tournamentList []model.Tournament) {
	rows, err := app.DB.Conn.Query(`
			select * from tournaments;
		`)
	errproc.HandleSQLErr("get rows from tournaments", err)
	for  rows.Next() {
		var tournament model.Tournament
		var id uuid.NullUUID
		var wid uuid.NullUUID
		err := rows.Scan(&tournament.Id, &tournament.Name, &tournament.Deposit, &tournament.Prize, &id, &wid)
		if wid.Valid {
			tournament.WinnerId = wid.UUID
		}
		errproc.FprintErr("Unexpected error trying to scan tournament: %v\n", err)
		if id.Valid {
			var d database.DB
			d.Connect("tournament-app")
			r, err := d.Conn.Query(`
				select userid from participants
					where id = $1;
			`, id.UUID)
			errproc.HandleSQLErr("get rows from participants", err)
			for i := 0; r.Next(); i++ {
				var u uuid.UUID
				err := r.Scan(&u)
				errproc.FprintErr("Unexpected error trying to scan participant: %v\n", err)
				tournament.Users = append(tournament.Users, u)
			}
		}
		tournamentList = append(tournamentList, tournament)
	}
	return
}

func GetTournament(id uuid.UUID) (tournament model.Tournament) {
	tournamentList := GetTournaments()
	for _, tournament = range tournamentList {
		if tournament.Id == id {
			return
		}
	}
	return model.Tournament{}
}

func DeleteTournament(id uuid.UUID) {
	tournamentList := GetTournaments()
	for _, tournament := range tournamentList {
		if tournament.Id == id {
			_, err := app.DB.Conn.Exec(`
				delete from participants using tournaments 
					where tournament.id = $1 and participants.id = tournaments.users;
				delete from tournaments
					where id = $1;
			`, tournament.Id)
			errproc.HandleSQLErr("delete tournament", err)
			break
		}
	}
}

func JoinTournament(id uuid.UUID, userID uuid.UUID) {
	tournamentList := GetTournaments()
	userList := user.GetUsers()
	for _, tournament := range tournamentList {
		if tournament.Id == id {
			for _, user := range userList {
				if user.Id == userID {
					if user.Balance >= tournament.Deposit {
						user.Balance -= tournament.Deposit
						var pid uuid.NullUUID
						err := app.DB.Conn.QueryRow(`
								select users from tournaments
									where id = $1;
							`, tournament.Id).Scan(&pid)
						errproc.HandleSQLErr("scan participants", err)
						if pid.Valid {
							_, err = app.DB.Conn.Exec(`
									insert into participants (id, userid) values
										($1, $2);
								`, pid, user.Id)
							errproc.HandleSQLErr("insert participant", err)
							_, err = app.DB.Conn.Exec(`
									update users set balance = $2
										where id = $1;
								`, user.Id, user.Balance)
							errproc.HandleSQLErr("update user's balance", err)
						} else {
							_, err = app.DB.Conn.Exec(`
									insert into participants (userid) values
										($1);
								`, user.Id)
							errproc.HandleSQLErr("insert participant with nil pid", err)
							_, err = app.DB.Conn.Exec(`
									update tournaments set users = participants.id from participants 
										where participants.userid = $1;
								`, user.Id)
							errproc.HandleSQLErr("update tournament's users with nil pid", err)
							_, err = app.DB.Conn.Exec(`
									update users set balance = $2
										where id = $1;
								`, user.Id, user.Balance)
							errproc.HandleSQLErr("update user's balance", err)
						}
					}
					break
				}
			}
			break
		}
	}
}

func FinishTournament(id uuid.UUID) int {
	tournamentList := GetTournaments()
	userList := user.GetUsers()
	for _, tournament := range tournamentList {
		if tournament.Id == id {
			if tournament.WinnerId.String() == "00000000-0000-0000-0000-000000000000" {
				var winner userm.User
				for _, userId := range tournament.Users {
					for _, user := range userList {
						if user.Id == userId {
							if user.Balance > winner.Balance {
								winner = user
							}
							break
						}
					}
				}
				winner.Balance += tournament.Prize
				_, err := app.DB.Conn.Exec(`
						update tournaments set winner = $1
							where id = $2;
					`, winner.Id, id)
				errproc.HandleSQLErr("set tournament's winner", err)
				_, err = app.DB.Conn.Exec(`
						update users set balance = $2
							where id = $1;
					`, winner.Id, winner.Balance)
				errproc.HandleSQLErr("update user's balance", err)
				return 200
			}
			break
		}
	}
	return 403
}
