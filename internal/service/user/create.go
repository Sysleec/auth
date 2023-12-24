package user

import (
	"context"

	"github.com/Sysleec/auth/internal/model"
)

func (s *serv) Create(ctx context.Context, usr *model.User) (int64, error) {
	user, err := s.userRepo.Create(ctx, usr)
	if err != nil {
		return 0, err
	}

	return user, nil
}
