package postgresql

import (
	"tournament/pkg/domain"
)

// InsertUser in DB creates a new user with the name and balance.
func (c *PSQLController) InsertUser(name string, balance int) (uint64, error) {
	var id uint64

	err := c.Database.QueryRow(`
			insert into users (name, balance) values 
				($1, $2) returning id;
		`, name, balance).Scan(&id)
	return id, err
}

func (c *PSQLController) scanUserRow(row Rows) (domain.User, error) {
	u := domain.User{}

	err := row.Scan(&u.ID, &u.Name, &u.Balance)
	if err != nil {
		return u, err
	}

	return u, nil
}

// GetUsers returns a slice of all users from DB.
func (c *PSQLController) GetUsers() ([]domain.User, error) {
	rows, err := c.Database.Query(`select * from users;`)
	if err != nil {
		return nil, err
	}

	var users []domain.User
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
func (c *PSQLController) GetUserByID(uid uint64) (domain.User, error){
	u := domain.User{}

	err := c.Database.QueryRow(`select * from users where id = $1;`,
		uid).Scan(&u.ID, &u.Name, &u.Balance)
	if err != nil {
		return u, err
	}

	return u, nil
}

// DeleteUserByID deletes from DB the user with uid.
func (c *PSQLController) DeleteUserByID(uid uint64) error {
	_, err := c.Database.Exec(`delete from users where id = $1;`,
		uid)
	return err
}

// UpdateUserBalanceByID updates balance of the user with uid in DB.
func (c *PSQLController) UpdateUserBalanceByID(uid uint64, balance int) error {
	_, err := c.Database.Exec(`update users set balance = $1 where id = $2;`,
		balance, uid)
	return err
}
