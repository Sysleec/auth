package user

import (
	"context"
	"strings"

	"github.com/Sysleec/auth/internal/converter"
	"github.com/Sysleec/auth/internal/logger"
	"github.com/Sysleec/auth/pkg/sys/validate"
	desc "github.com/Sysleec/auth/pkg/user_v1"
	"github.com/opentracing/opentracing-go"
)

// Create creates a new user
func (s *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {

	err := validate.Validate(
		ctx,
		validatePassSpaces(req.Password),
		validatePassMatching(req.Password, req.PasswordConfirm),
	)

	if err != nil {
		return nil, err
	}

	span, ctx := opentracing.StartSpanFromContext(ctx, "create user")
	defer span.Finish()

	id, err := s.userService.Create(ctx, converter.ToUserFromDesc(req))
	if err != nil {
		return nil, err
	}

	logger.Infof("created user with id: %d", id)

	return &desc.CreateResponse{Id: id}, nil
}

func validatePassSpaces(pass string) validate.Condition {
	return func(ctx context.Context) error {
		if strings.Contains(pass, " ") {
			return validate.NewValidationError("name cannot contain spaces")
		}

		return nil
	}
}

func validatePassMatching(pass, passConfirm string) validate.Condition {
	return func(ctx context.Context) error {
		if pass != passConfirm {
			return validate.NewValidationError("passwords do not match")
		}

		return nil
	}
}
