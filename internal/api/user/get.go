package user

import (
	"context"
	"errors"

	"github.com/Sysleec/auth/internal/converter"
	"github.com/Sysleec/auth/internal/logger"
	"github.com/Sysleec/auth/internal/model"
	desc "github.com/Sysleec/auth/pkg/user_v1"
	"github.com/opentracing/opentracing-go"
)

// Get gets a user by id
func (s *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "get user")
	defer span.Finish()

	user, err := s.userService.Get(ctx, req.GetId())
	if err != nil {
		switch {
		case errors.Is(err, model.ErrUserNotFound):
			return nil, model.ErrUserNotFound
		default:
			return nil, err
		}
	}

	logger.Infof("got user with id: %d", user.ID)

	return converter.ToUserFromService(user), nil
}
