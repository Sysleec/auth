package user

import (
	"context"

	"github.com/Sysleec/auth/internal/logger"
	desc "github.com/Sysleec/auth/pkg/user_v1"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Delete deletes a user by id
func (s *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "delete user")
	defer span.Finish()

	_, err := s.userService.Delete(ctx, req.Id)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	logger.Infof("deleted user with id: %d", req.Id)

	return &emptypb.Empty{}, nil
}
