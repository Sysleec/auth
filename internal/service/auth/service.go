package login

import (
	"github.com/Sysleec/auth/internal/client/db"
	"github.com/Sysleec/auth/internal/repository"
	"github.com/Sysleec/auth/internal/service"
)

type serverAuth struct {
	loginRepository repository.AuthRepository
	txManager       db.TxManager
}

func NewService(
	loginRepository repository.AuthRepository,
	txManager db.TxManager,
) service.AuthService {
	return &serverAuth{
		loginRepository: loginRepository,
		txManager:       txManager,
	}
}
