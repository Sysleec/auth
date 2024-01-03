package user

import (
	"context"

	"github.com/Sysleec/auth/internal/model"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *serv) Update(ctx context.Context, usr *model.User) (*emptypb.Empty, error) {
	_, err := s.userRepo.Update(ctx, usr)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
