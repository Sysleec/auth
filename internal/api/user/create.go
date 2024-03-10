package user

import (
	"context"

	"github.com/Sysleec/auth/internal/converter"
	"github.com/Sysleec/auth/internal/logger"
	desc "github.com/Sysleec/auth/pkg/user_v1"
)

// Create creates a new user
func (s *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	logger.Debugf("creating user: %v", req)

	id, err := s.userService.Create(ctx, converter.ToUserFromDesc(req))
	if err != nil {
		logger.Errorf("failed to create user: %s", err.Error())
		return nil, err
	}

	logger.Infof("created user with id %d\n", id)

	return &desc.CreateResponse{Id: id}, nil
}
