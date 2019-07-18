package usecases

import (
	uuid "github.com/satori/go.uuid"
	"tournament/env/errproc"
	app "tournament/pkg"
	"tournament/pkg/user/model"
)

// CreateUser inserts new user instance in database.
func CreateUser(user *model.User) error {
	err := app.DB.Conn.QueryRow(`
			insert into users (name, balance) values 
				($1, $2) returning id;
		`, user.Name, user.Balance).Scan(&user.ID)
	return err
}

// GetUsers returns a slice of all users in the database.
func GetUsers() ([]model.User, error) {
	rows, err := app.DB.Conn.Query(`select * from users;`)
	if err != nil {
		return nil, err
	}

	var userList []model.User
	for rows.Next() {
		var user model.User

		err := rows.Scan(&user.ID, &user.Name, &user.Balance)
		if err != nil {
			return nil, err
		}

		userList = append(userList, user)
	}

	return userList, nil
}

// GetUser returns user instance with id from
// slice list of all users in database.
func GetUser(id uuid.UUID) (model.User, error) {
	userList, err := GetUsers()
	if err != nil {
		return model.User{}, err
	}

	for _, u := range userList {
		if u.ID == id {
			return u, nil
		}
	}

	return model.User{}, errproc.NoUserWithID
}

// DeleteUser deletes user instance with id from database.
func DeleteUser(id uuid.UUID) error{
	userList, err := GetUsers()
	if err != nil {
		return err
	}

	for index := range userList {
		if userList[index].ID != id {
			continue
		}

		_, err := app.DB.Conn.Exec(`delete from users where id = $1;`,
			id)
		return err
	}

	return errproc.NoUserWithID
}

// FundUser adds points to balance of user's with the id.
func FundUser(id uuid.UUID, points int) error {
	userList, err := GetUsers()
	if err != nil {
		return err
	}

	for index := range userList {
		if userList[index].ID != id {
			continue
		}

		userList[index].Balance += points

		_, err := app.DB.Conn.Exec(`update users set balance = $1 where id = $2;`,
			userList[index].Balance, id)
		return err
	}

	return errproc.NoUserWithID
}
