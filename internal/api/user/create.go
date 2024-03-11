package user

import (
	"context"

	"github.com/Sysleec/auth/internal/converter"
	desc "github.com/Sysleec/auth/pkg/user_v1"
)

// Create creates a new user
func (s *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := s.userService.Create(ctx, converter.ToUserFromDesc(req))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{Id: id}, nil
}
