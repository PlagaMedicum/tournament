package usecases

import (
	"tournament/pkg/domain"
)

const defaultBalance = 700

// CreateUser inserts new user instance in Repository.
func (c *Controller) CreateUser(name string) (string, error) {
	id, err := c.Repository.InsertUser(name, defaultBalance)
	return id, err
}

// GetUser returns user instance with ID from Repository.
func (c *Controller) GetUser(id string) (domain.User, error) {
	u, err := c.Repository.GetUserByID(id)
	return u, err
}

// DeleteUser deletes user instance with ID from Repository.
func (c *Controller) DeleteUser(id string) error {
	err := c.Repository.DeleteUserByID(id)
	return err
}

// FundUser adds points to balance of user's with the ID.
func (c *Controller) FundUser(id string, points int) error {
	u, err := c.Repository.GetUserByID(id)
	if err != nil {
		return err
	}

	u.Balance += points

	err = c.Repository.UpdateUserBalanceByID(u.ID, u.Balance)
	return err
}
