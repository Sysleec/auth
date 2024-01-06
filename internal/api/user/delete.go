package user

import (
	"context"
	"fmt"

	desc "github.com/Sysleec/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	_, err := s.userService.Delete(ctx, req.Id)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	fmt.Printf("deleted user with id %d\n", req.Id)

	return &emptypb.Empty{}, nil
}
