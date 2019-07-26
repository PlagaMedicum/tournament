package postgresql

import (
	"github.com/jackc/pgx"
	"tournament/pkg/domain"
)

func (c *PSQLController) InsertUser(name string, balance int) (string, error) {
	id := c.IDFactory.New()

	err := c.Handler.QueryRow(`
			insert into users (name, balance) values 
				($1, $2) returning id;
		`, name, balance).Scan(id)
	return id.String(), err
}

func (c *PSQLController) scanUserRow(row *pgx.Rows) (domain.User, error) {
	u := domain.User{}
	id := c.IDFactory.New()

	err := row.Scan(&id, &u.Name, &u.Balance)
	if err != nil {
		return domain.User{}, err
	}

	u.ID = id.String()
	return u, nil
}

// GetUsers returns a slice of all users in the postgresqlDB.
func (c *PSQLController) GetUsers() ([]domain.User, error) {
	rows, err := c.Handler.Query(`select * from users;`)
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

func (c *PSQLController) GetUserByID(uid string) (domain.User, error){
	u := domain.User{}
	id := c.IDFactory.New()

	err := id.UnmarshalText([]byte(uid))
	if err != nil {
		return u, err
	}

	err = c.Handler.QueryRow(`
			select from users where id = $1;
		`, id).Scan(&id, &u.Name, &u.Balance)
	if err != nil {
		return u, err
	}

	u.ID = id.String()
	return u, nil
}

func (c *PSQLController) DeleteUserByID(uid string) error {
	id := c.IDFactory.New()

	err := id.UnmarshalText([]byte(uid))
	if err != nil {
		return err
	}

	_, err = c.Handler.Exec(`delete from users where id = $1;`,
		id)
	return err
}

func (c *PSQLController) UpdateUserBalanceByID(balance int, uid string) error {
	id := c.IDFactory.New()

	err := id.UnmarshalText([]byte(uid))
	if err != nil {
		return err
	}

	_, err = c.Handler.Exec(`update users set balance = $1 where id = $2;`,
		balance, id)
	return err
}
