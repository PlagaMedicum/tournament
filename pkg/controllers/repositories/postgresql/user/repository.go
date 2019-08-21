package user

import (
	"tournament/pkg/controllers/repositories/postgresql"
	"tournament/pkg/domain/user"
)

// Controller contains an implementation of postgresql.Database interface.
type Controller struct {
	postgresql.Database
}

// InsertUser in DB creates a new user with the name and balance.
func (c *Controller) InsertUser(name string, balance int) (uint64, error) {
	var id uint64

	err := c.QueryRow(`
			insert into users (name, balance) values 
				($1, $2) returning id;
		`, name, balance).Scan(&id)
	return id, err
}

func (c *Controller) scanUserRow(row postgresql.Rows) (user.User, error) {
	u := user.User{}

	err := row.Scan(&u.ID, &u.Name, &u.Balance)
	if err != nil {
		return u, err
	}

	return u, nil
}

// GetUsers returns a slice of all users from DB.
func (c *Controller) GetUsers() ([]user.User, error) {
	rows, err := c.Query(`select * from users;`)
	if err != nil {
		return nil, err
	}

	var users []user.User
	for rows.Next() {
		u, err := c.scanUserRow(rows)
		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}

// GetUserByID returns from DB the user with uid.
func (c *Controller) GetUserByID(uid uint64) (user.User, error) {
	u := user.User{}

	err := c.QueryRow(`select * from users where id = $1;`,
		uid).Scan(&u.ID, &u.Name, &u.Balance)
	if err != nil {
		return u, err
	}

	return u, nil
}

// DeleteUserByID deletes from DB the user with uid.
func (c *Controller) DeleteUserByID(uid uint64) error {
	_, err := c.Exec(`
			delete from participants where participants.userid = $1;
		`, uid)
	if err != nil {
		return err
	}

	_, err = c.Exec(`
			delete from winners where winners.userid = $1;
		`, uid)
	if err != nil {
		return err
	}

	_, err = c.Exec(`delete from users where id = $1;`,
		uid)
	return err
}

// UpdateUserBalanceByID updates balance of the user with uid in DB.
func (c *Controller) UpdateUserBalanceByID(uid uint64, balance int) error {
	_, err := c.Exec(`update users set balance = $1 where id = $2;`,
		balance, uid)
	return err
}
