package user

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *serv) Delete(ctx context.Context, id int64) (*emptypb.Empty, error) {
	_, err := s.userRepo.Delete(ctx, id)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
