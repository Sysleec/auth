package user

import (
	"context"

	"github.com/Sysleec/auth/internal/logger"
	desc "github.com/Sysleec/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	logger.Debugf("deleting user: %v", req)

	_, err := s.userService.Delete(ctx, req.Id)
	if err != nil {
		logger.Errorf("failed to delete user: %s", err.Error())
		return &emptypb.Empty{}, err
	}

	logger.Infof("deleted user with id %d\n", req.Id)

	return &emptypb.Empty{}, nil
}
