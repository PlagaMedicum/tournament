package usecases

import (
	"tournament/pkg/domain"
)

const defaultBalance = 700

// CreateUser inserts new user instance in Repository.
func (c *Controller) CreateUser(name string) (uint64, error) {
	id, err := c.Repository.InsertUser(name, defaultBalance)
	return id, err
}

// GetUser returns user instance with ID from Repository.
func (c *Controller) GetUser(id uint64) (domain.User, error) {
	u, err := c.Repository.GetUserByID(id)
	return u, err
}

// DeleteUser deletes user instance with ID from Repository.
func (c *Controller) DeleteUser(id uint64) error {
	err := c.Repository.DeleteUserByID(id)
	return err
}

// FundUser adds points to balance of user's with the ID.
func (c *Controller) FundUser(id uint64, points int) error {
	u, err := c.Repository.GetUserByID(id)
	if err != nil {
		return err
	}

	u.Balance += points

	err = c.Repository.UpdateUserBalanceByID(u.ID, u.Balance)
	return err
}
