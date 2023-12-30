package user

import (
	"github.com/Sysleec/auth/internal/client/db"
	"github.com/Sysleec/auth/internal/repository"
	def "github.com/Sysleec/auth/internal/service"
)

var _ def.UserService = (*serv)(nil)

type serv struct {
	userRepo  repository.UserRepository
	txManager db.TxManager
}

// NewService returns a new user service
func NewService(userRepo repository.UserRepository,
	txManager db.TxManager) *serv {
	return &serv{
		userRepo:  userRepo,
		txManager: txManager,
	}
}
