package user

import (
	"context"

	"github.com/Sysleec/auth/internal/converter"
	desc "github.com/Sysleec/auth/pkg/user_v1"
)

// Get gets a user by id
func (s *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	user, err := s.userService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return converter.ToUserFromService(user), nil
}
