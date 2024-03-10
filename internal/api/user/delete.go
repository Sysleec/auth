package user

import (
	"context"

	desc "github.com/Sysleec/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Delete deletes a user by id
func (s *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	_, err := s.userService.Delete(ctx, req.Id)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
