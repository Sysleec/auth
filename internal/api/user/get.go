package user

import (
	"context"
	"errors"

	"github.com/Sysleec/auth/internal/converter"
	"github.com/Sysleec/auth/internal/model"
	desc "github.com/Sysleec/auth/pkg/user_v1"
)

// Get gets a user by id
func (s *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	user, err := s.userService.Get(ctx, req.GetId())
	if err != nil {
		switch {
		case errors.Is(err, model.ErrUserNotFound):
			return nil, model.ErrUserNotFound
		default:
			return nil, err
		}
	}

	return converter.ToUserFromService(user), nil
}
