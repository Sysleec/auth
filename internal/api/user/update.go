package user

import (
	"context"
	"fmt"

	"github.com/Sysleec/auth/internal/converter"
	desc "github.com/Sysleec/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	_, err := s.userService.Update(ctx, converter.ToUserFromDescUpdate(req))
	if err != nil {
		return &emptypb.Empty{}, err
	}
	fmt.Printf("updated user with id %d", req.Id)

	return &emptypb.Empty{}, nil
}
