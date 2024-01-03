package user

import (
	"github.com/Sysleec/auth/internal/client/db"
	"github.com/Sysleec/auth/internal/repository"
	"github.com/Sysleec/auth/internal/service"
)

var _ service.UserService = (*serv)(nil)

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

func NewMockService(deps ...interface{}) service.UserService {
	srv := serv{}

	for _, v := range deps {
		switch s := v.(type) {
		case repository.UserRepository:
			srv.userRepo = s
		}
	}
	return &srv
}

func NewMockTxService(deps ...interface{}) *serv {
	srv := serv{}

	for _, v := range deps {
		switch s := v.(type) {
		case serv:
			srv = s
		}
	}
	return &srv
}
