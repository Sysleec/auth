package user

import (
	"context"

	"github.com/Sysleec/auth/internal/converter"
	"github.com/Sysleec/auth/internal/logger"
	desc "github.com/Sysleec/auth/pkg/user_v1"
)

// Get gets a user by id
func (s *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	logger.Debugf("getting user: %v", req)

	user, err := s.userService.Get(ctx, req.GetId())
	if err != nil {
		logger.Errorf("failed to get user: %s", err.Error())
		return nil, err
	}

	logger.Infof("got user - Id: %d | Name: %s | Email: %s\n", user.ID, user.Name, user.Email)

	return converter.ToUserFromService(user), nil
}
