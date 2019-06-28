package database

import (
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/pgtype"
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
		Name    string 				`json:"name"`
		Balance int    				`json:"balance"`
	}
	Tournament struct {
		ID
		Name    string   			`json:"name"`
		Deposit int      			`json:"deposit"`
		Prize   int      			`json:"prize"`
		Users   []uuid.UUID			`json:"users"`
		Winner  ID 					`json:"winner"`
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
				winner text not null
			);
			create table participants(
				id uuid constraint participant_pk primary key default uuid_generate_v4() not null,
				tournamentid uuid not null,
				userid uuid not null
			);
			insert into users (id, name, balance) values 
				('bef80618-779e-4cbd-b776-cbd27386a902', 'Samuel Plaunik', 1200);
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
			select * from Users
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
					delete from users where id = $1;
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
					update users set balance = $1 where id = $2;
				`, userList[index].Balance, id)
			errproc.HandleSQLErr("fund user", err)
			break
		}
	}
}

func (db *DB) CreateTournament(tournament *Tournament) {
	err := db.Conn.QueryRow(`
			insert into tournaments (name, deposit, prize, winner) values
				($1, $2, $3, $4) returning id;
		`, tournament.Name, tournament.Deposit, tournament.Prize, tournament.Winner.Id).Scan(&tournament.Id)
	errproc.HandleSQLErr("create tournament", err)
}

func (db *DB) GetTournaments() (tournamentList []Tournament) {
	rows, err := db.Conn.Query(`
			select * from tournaments
		`)
	errproc.HandleSQLErr("get rows from tournaments", err)
	for  rows.Next() {
		var (
			tournament Tournament
			id pgtype.UUID
		)
		err := rows.Scan(&tournament.Id, &tournament.Name, &tournament.Deposit, &tournament.Prize, &id, &tournament.Winner.Id)
		if err != nil {
			errproc.FprintErr("Unexpected error trying to scan tournament: %v\n", err)
		}
		r, err := db.Conn.Query(`
				select userid from participants where id = $1
			`, id)
		errproc.HandleSQLErr("get rows from participants", err)
		for i := 0; r.Next(); i++ {
			err := rows.Scan(&tournament.Users[i])
			if err != nil {
				errproc.FprintErr("Unexpected error trying to scan participant: %v\n", err)
			}
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
				delete from tournaments where id = $1;
				delete from participants where tournamentid = $1;
			`, id)
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
						var pid pgtype.UUID
						rows, err := db.Conn.Query(`
								select id from participants where tournamentid = $1
							`, )
						errproc.HandleSQLErr("get participants id", err)
						err = rows.Scan(&pid)
						errproc.FprintErr("Unexpected error trying to scan participants id: %v\n", err)
						_, err = db.Conn.Exec(`
								insert into participants (id, userid, tournamentid) values
									($1, $2, $3);
								update users set balance = $4 where id = $2;
							`, pid, user.Id, tournament.Id, user.Balance)
						errproc.HandleSQLErr("join tournament", err)
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
			var b ID
			b.FromString("00000000-0000-0000-0000-000000000000")
			if tournament.Winner.Id == b.Id {
				var winner User
				winner.Balance = -99999
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
				_, err := db.Conn.Exec(`
						update tournaments set winner = $1 where id = $2;
						update users set balance = $3 where id = $4;
					`, winner.Id, id, winner.Balance, winner.Id)
				errproc.HandleSQLErr("finish tournament", err)
				return 200
			}
			break
		}
	}
	return 403
}
