package user

import "tournament/pkg/domain/user"

type Usecases interface {
	CreateUser(string) (uint64, error)
	GetUser(uint64) (user.User, error)
	DeleteUser(uint64) error
	FundUser(uint64, int) error
}
