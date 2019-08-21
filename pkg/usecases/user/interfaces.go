package user

import "tournament/pkg/domain/user"

// Repository is interface of database repository methods
// used in user usecases.
type Repository interface {
	InsertUser(string, int) (uint64, error)
	GetUserByID(uint64) (user.User, error)
	DeleteUserByID(uint64) error
	UpdateUserBalanceByID(uint64, int) error
}
