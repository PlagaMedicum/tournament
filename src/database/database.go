package database

import (
	"github.com/jackc/pgx"
	uuid "github.com/satori/go.uuid"
	"os"
	"tournament/src/errproc"
)

type (
	ID struct {
		Id 		uuid.UUID 	`json:"id"`
	}
	User struct {
		ID
		Name    string 		`json:"name"`
		Balance int    		`json:"balance"`
	}
	Tournament struct {
		ID
		Name    string   	`json:"name"`
		Deposit int      	`json:"deposit"`
		Prize   int      	`json:"prize"`
		Users   []ID		`json:"users"`
		Winner  ID 			`json:"winner"`
	}
	DB struct {
		Conn *pgx.Conn
	}
)

func (id ID) Get() (val uuid.UUID) {
	return id.Id
}

func (id ID) FromString(s string) {
	var err error
	id.Id, err = uuid.FromString(s)
	errproc.FprintErr("Error trying to convert string in uuid: %v\n", err)
}

func (db DB) Connect(applicationName string) {
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

func (db DB) InitTables() {
	_, err := db.Conn.Exec(`
		drop table if exists Users;
		drop table if exists Tournaments;
		create table Users(
			Id uuid constraint user_pk primary key default uuid_generate_v4() not null,
			Name text not null,
			Balance int not null
		);
		create table Tournaments(
			Id uuid constraint tournament_pk primary key default uuid_generate_v4() not null,
			Name text not null,
			Deposit int not null,
			Prize int not null,
			UserID text constraint unique_userid unique not null,
			Winner text not null
		);
	`)
	if err != nil {
		errproc.FprintErr("Unable to initialise tables: %v\n", err)
		os.Exit(1)
	}
}

func (db DB) Close() {
	err := db.Conn.Close()
	errproc.FprintErr("Unexpected error trying to close DB connection: %v\n", err)
}

func (db DB) CreateUser(user *User) {
	err := db.Conn.QueryRow(`
		insert into Users (Name, Balance) values
			($1, $2) returning ID;
	`, user.Name, user.Balance).Scan(&user.Id)
	errproc.HandleSQLErr("create user", err)
}

func (db DB) GetUsers() (userList []User) {
	rows, err := db.Conn.Query(`
			select * from Users
		`)
	errproc.HandleSQLErr("get rows from Users", err)

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Balance)
		if err != nil {
			errproc.FprintErr("Unexpected error trying to scan user: %v\n", err)
		}
		userList = append(userList, user)
	}
	return
}

func (db DB) GetUser(id uuid.UUID) (user User) {
	userList := db.GetUsers()
	for _, user = range userList {
		if user.Id == id {
			return
		}
	}
	return User{}
}

func (db DB) DeleteUser(id uuid.UUID) {
	userList := db.GetUsers()
	for index := range userList {
		if userList[index].Id == id {
			_, err := db.Conn.Exec(`
				delete from Users where ID = $1;
			`, id)
			errproc.HandleSQLErr("delete user", err)
			break
		}
	}
}

func (db DB) FundUser(id uuid.UUID, points int) {
	userList := db.GetUsers()
	for index := range userList {
		if userList[index].Id == id {
			userList[index].Balance += points
			_, err := db.Conn.Exec(`
				update Users set Balance = $1 where ID = $2;
			`, userList[index].Balance, id)
			errproc.HandleSQLErr("fund user", err)
			break
		}
	}
}

func (db DB) CreateTournament(tournament *Tournament) {
	err := db.Conn.QueryRow(`
		insert into Tournament (Name, Deposit, Prize, UserID, Winner) values
			($1, $2, $3, $4, $5) returning ID;
	`, tournament.Name, tournament.Deposit, tournament.Prize, tournament.Users, tournament.Winner).Scan(&tournament.Id)
	errproc.HandleSQLErr("create tournament", err)
}

func (db DB) GetTournaments() (tournamentList []Tournament) {
	rows, err := db.Conn.Query(`
			select * from Tournaments
		`)
	errproc.HandleSQLErr("get rows from Tournaments", err)
	for rows.Next() {
		var tournament Tournament
		err := rows.Scan(&tournament.Id, &tournament.Name, &tournament.Deposit, &tournament.Prize, &tournament.Users, &tournament.Winner)
		if err != nil {
			errproc.FprintErr("Unexpected error trying to scan tournament: %v\n", err)
		}
		tournamentList = append(tournamentList, tournament)
	}
	return
}

func (db DB) GetTournament(id ID) (tournament Tournament) {
	tournamentList := db.GetTournaments()
	for _, tournament = range tournamentList {
		if tournament.ID == id {
			return
		}
	}
	return Tournament{}
}

func (db DB) DeleteTournament(id uuid.UUID) {
	tournamentList := db.GetTournaments()
	for index := range tournamentList {
		if tournamentList[index].Id == id {
			_, err := db.Conn.Exec(`
				delete from Tournaments where ID = $1;
			`, id)
			errproc.HandleSQLErr("delete tournament", err)
			break
		}
	}
}

func (db DB) JoinTournament(id uuid.UUID, userID uuid.UUID) {
	tournamentList := db.GetTournaments()
	userList := db.GetUsers()
	for tindex, tournament := range tournamentList {
		if tournamentList[tindex].Id == id {
			for uindex := range userList {
				if userList[uindex].ID.Id == userID {
					if userList[uindex].Balance >= tournamentList[tindex].Deposit {
						userList[uindex].Balance -= tournamentList[tindex].Deposit
						_, err := db.Conn.Exec(`
							insert into Tournament (ID, Name, Deposit, Prize, UserID, Winner) values
							($1, $2, $3, $4, $5);
						`, tournament.Id, tournament.Name, tournament.Deposit, tournament.Prize, userID, tournament.Winner)
						errproc.HandleSQLErr("join tournament", err)
					}
					break
				}
			}
			break
		}
	}
}

func (db DB) FinishTournament(id uuid.UUID) {
	tournamentList := db.GetTournaments()
	userList := db.GetUsers()
	for _, tournament := range tournamentList {
		if tournament.Id == id {
			if b, _ := uuid.FromString("00000000-0000-0000-0000-000000000000"); tournament.Winner.Id == b  {
				var winner User
				winner.Balance = -99999
				for _, userId := range tournament.Users {
					for _, user := range userList {
						if user.Id == userId.Get() {
							if user.Balance > winner.Balance {
								winner = user
							}
							break
						}
					}
				}
				_, err := db.Conn.Exec(`
					update Tournaments set Winner = $1 where ID = $2;
					update Users set Balance = $3 where ID = $4;
				`, winner.Id, id, winner.Balance, winner.Id)
				errproc.HandleSQLErr("finish tournament", err)
			}
			break
		}
	}
}
