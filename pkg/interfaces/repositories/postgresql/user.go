package postgresql

import (
	"github.com/jackc/pgx"
	"tournament/pkg/domain"
)

func (c *PSQLController) InsertUser(name string, balance int) (string, error) {
	var id string
	err := c.Handler.QueryRow(`
			insert into users (name, balance) values 
				($1, $2) returning id;
		`, name, balance).Scan(id)
	return id, err
}

func (c *PSQLController) scanUserRow(row *pgx.Rows) (domain.User, error) {
	u := domain.User{}

	err := row.Scan(&u.ID, &u.Name, &u.Balance)
	return u, err
}

// GetUsers returns a slice of all users in the postgresqlDB.
func (c *PSQLController) GetUsers() ([]domain.User, error) {
	rows, err := c.Handler.Query(`select * from users;`)
	if err != nil {
		return nil, err
	}

	var userList []domain.User
	for rows.Next() {
		u, err := c.scanUserRow(rows)
		if err != nil {
			return nil, err
		}

		userList = append(userList, u)
	}

	return userList, nil
}

func (c *PSQLController) GetUserByID(id string) (domain.User, error){
	u := domain.User{}

	err := c.Handler.QueryRow(`
			select from users where id = $1;
		`, id).Scan(&u.ID, &u.Name, &u.Balance)
	return u, err
}

func (c *PSQLController) DeleteUserByID(id string) error {
	_, err := c.Handler.Exec(`delete from users where id = $1;`,
		id)
	return err
}

func (c *PSQLController) UpdateUserBalanceByID(balance int, id string) error {
	_, err := c.Handler.Exec(`update users set balance = $1 where id = $2;`,
		balance, id)
	return err
}
