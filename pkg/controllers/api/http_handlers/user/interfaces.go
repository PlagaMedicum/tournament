package user

import "tournament/pkg/domain/user"

// Usecases is the interface for user usecases.
type Usecases interface {
	CreateUser(string) (uint64, error)
	GetUser(uint64) (user.User, error)
	DeleteUser(uint64) error
	FundUser(uint64, int) error
}
