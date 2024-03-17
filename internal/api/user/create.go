package user

import (
	"context"

	"github.com/Sysleec/auth/internal/converter"
	"github.com/Sysleec/auth/internal/logger"
	desc "github.com/Sysleec/auth/pkg/user_v1"
	"github.com/opentracing/opentracing-go"
)

// Create creates a new user
func (s *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "create user")
	defer span.Finish()

	id, err := s.userService.Create(ctx, converter.ToUserFromDesc(req))
	if err != nil {
		return nil, err
	}

	logger.Infof("created user with id: %d", id)

	return &desc.CreateResponse{Id: id}, nil
}
