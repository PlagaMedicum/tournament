package usecases

import (
	uuid "github.com/satori/go.uuid"
	app "tournament/pkg/boot"
	"tournament/pkg/errproc"
	"tournament/pkg/user/model"
)

func CreateUser(user *model.User) {
	err := app.DB.Conn.QueryRow(`
			insert into users (name, balance) values 
				($1, $2) returning ID;
		`, user.Name, user.Balance).Scan(&user.Id)
	errproc.HandleSQLErr("create user", err)
}

func GetUsers() (userList []model.User) {
	rows, err := app.DB.Conn.Query(`
			select * from users;
		`)
	errproc.HandleSQLErr("get rows from users", err)
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.Id, &user.Name, &user.Balance)
		if err != nil {
			errproc.FprintErr("Unexpected error trying to scan user: %v\n", err)
		}
		userList = append(userList, user)
	}
	return
}

func GetUser(id uuid.UUID) (user model.User) {
	userList := GetUsers()
	for _, user = range userList {
		if user.Id == id {
			return
		}
	}
	return model.User{}
}

func DeleteUser(id uuid.UUID) {
	userList := GetUsers()
	for index := range userList {
		if userList[index].Id == id {
			_, err := app.DB.Conn.Exec(`
					delete from users
						where id = $1;
				`, id)
			errproc.HandleSQLErr("delete user", err)
			break
		}
	}
}

func FundUser(id uuid.UUID, points int) {
	userList := GetUsers()
	for index := range userList {
		if userList[index].Id == id {
			userList[index].Balance += points
			_, err := app.DB.Conn.Exec(`
					update users set balance = $1
						where id = $2;
				`, userList[index].Balance, id)
			errproc.HandleSQLErr("fund user", err)
			break
		}
	}
}
