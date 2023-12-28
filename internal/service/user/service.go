package user

import (
	"github.com/Sysleec/auth/internal/repository"
	def "github.com/Sysleec/auth/internal/service"
)

var _ def.UserService = (*serv)(nil)

type serv struct {
	userRepo repository.UserRepository
}

// NewService returns a new user service
func NewService(userRepo repository.UserRepository) *serv {
	return &serv{
		userRepo: userRepo,
	}
}
