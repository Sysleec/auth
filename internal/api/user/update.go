package user

import (
	"context"

	"github.com/Sysleec/auth/internal/converter"
	"github.com/Sysleec/auth/internal/logger"
	desc "github.com/Sysleec/auth/pkg/user_v1"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Update updates a user
func (s *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "update user")
	defer span.Finish()

	_, err := s.userService.Update(ctx, converter.ToUserFromDescUpdate(req))
	if err != nil {
		return &emptypb.Empty{}, err
	}

	logger.Infof("updated user with id: %d", req.Id)

	return &emptypb.Empty{}, nil
}
