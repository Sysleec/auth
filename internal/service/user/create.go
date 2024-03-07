package user

import (
	"context"

	"github.com/Sysleec/auth/internal/model"
)

func (s *serv) Create(ctx context.Context, usr *model.User) (int64, error) {
	var id int64

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.userRepo.Create(ctx, usr)
		if errTx != nil {
			return errTx
		}

		_, errTx = s.userRepo.Get(ctx, id)
		if errTx != nil {
			return errTx
		}

		return nil

	})

	if err != nil {
		return 0, err
	}

	return id, nil
}
