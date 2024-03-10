package user

import (
	"context"

	"github.com/Sysleec/auth/internal/converter"
	"github.com/Sysleec/auth/internal/logger"
	desc "github.com/Sysleec/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	logger.Debugf("updating user: %v", req)

	_, err := s.userService.Update(ctx, converter.ToUserFromDescUpdate(req))
	if err != nil {
		logger.Errorf("failed to update user: %s", err.Error())
		return &emptypb.Empty{}, err
	}

	logger.Infof("updated user with id %d\n", req.Id)

	return &emptypb.Empty{}, nil
}
