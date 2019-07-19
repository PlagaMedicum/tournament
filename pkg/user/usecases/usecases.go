package usecases

import (
	uuid "github.com/satori/go.uuid"
	"tournament/env/errproc"
	app "tournament/pkg"
	"tournament/pkg/user/model"
)

const (
	defaultBalance = 700
)

type UserInterface interface {
	CreateUser(model.User) (uuid.UUID, error)
	GetUsers() ([]model.User, error)
	GetUser(uuid.UUID) (model.User, error)
	DeleteUser(uuid.UUID) error
	FundUser(uuid.UUID, int) error
}

type User struct {
	model.User
}

// CreateUserHandler inserts new user instance in database.
func (c *User) CreateUser(u model.User) (uuid.UUID, error) {
	u.Balance = defaultBalance

	err := app.DB.Conn.QueryRow(`
			insert into users (name, balance) values 
				($1, $2) returning id;
		`, u.Name, u.Balance).Scan(u.ID)
	return u.ID, err
}

// GetParticipants returns a slice of all users in the database.
func (c *User) GetUsers() ([]model.User, error) {
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

// GetUserHandler returns user instance with id from
// slice list of all users in database.
func (c *User) GetUser(id uuid.UUID) (model.User, error) {
	userList, err := c.GetUsers()
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

// DeleteUserHandler deletes user instance with id from database.
func (c *User) DeleteUser(id uuid.UUID) error {
	userList, err := c.GetUsers()
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
func (c *User) FundUser(id uuid.UUID, points int) error {
	userList, err := c.GetUsers()
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
