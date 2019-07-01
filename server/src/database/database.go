package database

import (
	"github.com/jackc/pgx"
	uuid "github.com/satori/go.uuid"
	"os"
	"tournament/src/errproc"
)

type (
	ID struct {
		Id 		uuid.UUID 		`json:"id"`
	}
	User struct {
		ID
		Name    string 			`json:"name"`
		Balance int    			`json:"balance"`
	}
	Tournament struct {
		ID
		Name     string        	`json:"name"`
		Deposit  int           	`json:"deposit"`
		Prize    int           	`json:"prize"`
		Users    []uuid.UUID   	`json:"users"`
		WinnerId uuid.UUID 		`json:"winner"`
	}
	DB struct {
		Conn *pgx.Conn
	}
)

func (t *Tournament) GetUsers() (users []uuid.UUID) {
	for _, user := range t.Users {
		users = append(users, user)
	}
	return
}

func (id *ID) Get() uuid.UUID {
	return id.Id
}

func (id *ID) FromString(s string) {
	err := id.Id.Scan(s)
	errproc.FprintErr("Error trying to convert string in uuid: %v\n", err)
}

func (db *DB) Connect(applicationName string) {
	var runtimeParams map[string]string
	runtimeParams = make(map[string]string)
	runtimeParams["application_name"] = applicationName
	connConfig := pgx.ConnConfig{
		User:              "postgres",
		Password:          "postgres",
		Host:              "localhost",
		Port:              5432,
		Database:          "tournament",
		TLSConfig:         nil,
		UseFallbackTLS:    false,
		FallbackTLSConfig: nil,
		RuntimeParams:     runtimeParams,
	}
	conn, err := pgx.Connect(connConfig)
	if err != nil {
		errproc.FprintErr("Unable to establish connection: %v\n", err)
		os.Exit(1)
	}
	db.Conn = conn
}

func (db *DB) InitTables() {
	_, err := db.Conn.Exec(`
			drop table if exists users;
			drop table if exists tournaments;
			drop table if exists participants;
			create table users(
				id uuid constraint user_pk primary key default uuid_generate_v4() not null,
				name text not null,
				balance int not null
			);
			create table tournaments(
				id uuid constraint tournament_pk primary key default uuid_generate_v4() not null,
				name text not null,
				deposit int not null,
				prize int not null,
				users uuid,
				winner uuid
			);
			create table participants(
				id uuid constraint participant_pk primary key default uuid_generate_v4() not null,
				userid uuid not null
			);
			insert into users (id, name, balance) values 
				('bef80618-779e-4cbd-b776-cbd27386a902', 'Samuel Plaunik', 1200);
			insert into tournaments (id, name, deposit, prize) values
				('6bfccaa8-9e88-4401-a12e-6559e709ee17', 'tour_1', 1000, 4000);
		`)
	if err != nil {
		errproc.FprintErr("Unable to initialise tables: %v\n", err)
		os.Exit(1)
	}
}

func (db *DB) Close() {
	err := db.Conn.Close()
	errproc.FprintErr("Unable to close connection with database: %v\n", err)
}

func (db *DB) CreateUser(user *User) {
	err := db.Conn.QueryRow(`
			insert into users (name, balance) values 
				($1, $2) returning ID;
		`, user.Name, user.Balance).Scan(&user.Id)
	errproc.HandleSQLErr("create user", err)
}

func (db *DB) GetUsers() (userList []User) {
	rows, err := db.Conn.Query(`
			select * from users;
		`)
	errproc.HandleSQLErr("get rows from users", err)
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Balance)
		if err != nil {
			errproc.FprintErr("Unexpected error trying to scan user: %v\n", err)
		}
		userList = append(userList, user)
	}
	return
}

func (db *DB) GetUser(id uuid.UUID) (user User) {
	userList := db.GetUsers()
	for _, user = range userList {
		if user.Id == id {
			return
		}
	}
	return User{}
}

func (db *DB) DeleteUser(id uuid.UUID) {
	userList := db.GetUsers()
	for index := range userList {
		if userList[index].Id == id {
			_, err := db.Conn.Exec(`
					delete from users
						where id = $1;
				`, id)
			errproc.HandleSQLErr("delete user", err)
			break
		}
	}
}

func (db *DB) FundUser(id uuid.UUID, points int) {
	userList := db.GetUsers()
	for index := range userList {
		if userList[index].Id == id {
			userList[index].Balance += points
			_, err := db.Conn.Exec(`
					update users set balance = $1
						where id = $2;
				`, userList[index].Balance, id)
			errproc.HandleSQLErr("fund user", err)
			break
		}
	}
}

func (db *DB) CreateTournament(tournament *Tournament) {
	err := db.Conn.QueryRow(`
			insert into tournaments (name, deposit, prize) values
				($1, $2, $3) returning id;
		`, tournament.Name, tournament.Deposit, tournament.Prize).Scan(&tournament.Id)
	errproc.HandleSQLErr("create tournament", err)
}

func (db *DB) GetTournaments() (tournamentList []Tournament) {
	rows, err := db.Conn.Query(`
			select * from tournaments;
		`)
	errproc.HandleSQLErr("get rows from tournaments", err)
	for  rows.Next() {
		var tournament Tournament
		var id uuid.NullUUID
		var wid uuid.NullUUID
		err := rows.Scan(&tournament.Id, &tournament.Name, &tournament.Deposit, &tournament.Prize, &id, &wid)
		if wid.Valid {
			tournament.WinnerId = wid.UUID
		}
		errproc.FprintErr("Unexpected error trying to scan tournament: %v\n", err)
		if id.Valid {
			var d DB
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
			d.Close()
		}
		tournamentList = append(tournamentList, tournament)
	}
	return
}

func (db *DB) GetTournament(id uuid.UUID) (tournament Tournament) {
	tournamentList := db.GetTournaments()
	for _, tournament = range tournamentList {
		if tournament.Id == id {
			return
		}
	}
	return Tournament{}
}

func (db *DB) DeleteTournament(id uuid.UUID) {
	tournamentList := db.GetTournaments()
	for _, tournament := range tournamentList {
		if tournament.Id == id {
			_, err := db.Conn.Exec(`
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

func (db *DB) JoinTournament(id uuid.UUID, userID uuid.UUID) {
	tournamentList := db.GetTournaments()
	userList := db.GetUsers()
	for _, tournament := range tournamentList {
		if tournament.Id == id {
			for _, user := range userList {
				if user.Id == userID {
					if user.Balance >= tournament.Deposit {
						user.Balance -= tournament.Deposit
						var pid uuid.NullUUID
						err := db.Conn.QueryRow(`
								select users from tournaments
									where id = $1;
							`, tournament.Id).Scan(&pid)
						errproc.HandleSQLErr("scan participants", err)
						if pid.Valid {
							_, err = db.Conn.Exec(`
									insert into participants (id, userid) values
										($1, $2);
								`, pid, user.Id)
							errproc.HandleSQLErr("insert participant", err)
							_, err = db.Conn.Exec(`
									update users set balance = $2
										where id = $1;
								`, user.Id, user.Balance)
							errproc.HandleSQLErr("update user's balance", err)
						} else {
							_, err = db.Conn.Exec(`
									insert into participants (userid) values
										($1);
								`, user.Id)
							errproc.HandleSQLErr("insert participant with nil pid", err)
							_, err = db.Conn.Exec(`
									update tournaments set users = participants.id from participants 
										where participants.userid = $1;
								`, user.Id)
							errproc.HandleSQLErr("update tournament's users with nil pid", err)
							_, err = db.Conn.Exec(`
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

func (db *DB) FinishTournament(id uuid.UUID) int {
	tournamentList := db.GetTournaments()
	userList := db.GetUsers()
	for _, tournament := range tournamentList {
		if tournament.Id == id {
			if tournament.WinnerId.String() == "00000000-0000-0000-0000-000000000000" {
				var winner User
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
				_, err := db.Conn.Exec(`
						update tournaments set winner = $1
							where id = $2;
					`, winner.Id, id)
				errproc.HandleSQLErr("set tournament's winner", err)
				_, err = db.Conn.Exec(`
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
