package user

import (
	"tournament/pkg/domain/user"
)

const defaultBalance = 700

// Controller contains an implementation of postgres.user.Repository interface.
type Controller struct {
	Repository
}

// CreateUser inserts new user instance in Repository.
func (c *Controller) CreateUser(name string) (uint64, error) {
	id, err := c.InsertUser(name, defaultBalance)
	return id, err
}

// GetUser returns user instance with ID from Repository.
func (c *Controller) GetUser(id uint64) (user.User, error) {
	u, err := c.GetUserByID(id)
	return u, err
}

// DeleteUser deletes user instance with ID from Repository.
func (c *Controller) DeleteUser(id uint64) error {
	err := c.DeleteUserByID(id)
	return err
}

// FundUser adds points to balance of user's with the ID.
func (c *Controller) FundUser(id uint64, points int) error {
	u, err := c.GetUserByID(id)
	if err != nil {
		return err
	}

	u.Balance += points

	err = c.UpdateUserBalanceByID(u.ID, u.Balance)
	return err
}
