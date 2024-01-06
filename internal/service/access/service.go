package access

import (
	"github.com/Sysleec/auth/internal/client/db"
	"github.com/Sysleec/auth/internal/repository"
	"github.com/Sysleec/auth/internal/service"
)

var _ service.AccessService = (*servAccess)(nil)

type servAccess struct {
	accessRepo repository.AccessRepository
	txManager  db.TxManager
}

// NewService returns a new user service
func NewService(accessRepo repository.AccessRepository,
	txManager db.TxManager) *servAccess {
	return &servAccess{
		accessRepo: accessRepo,
		txManager:  txManager,
	}
}
