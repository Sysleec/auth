package user

import (
	"context"

	"github.com/Sysleec/auth/internal/converter"
	desc "github.com/Sysleec/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Update updates a user
func (s *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	_, err := s.userService.Update(ctx, converter.ToUserFromDescUpdate(req))
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
