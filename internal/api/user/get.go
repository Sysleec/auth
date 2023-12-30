package user

import (
	"context"
	"fmt"

	"github.com/Sysleec/auth/internal/converter"
	desc "github.com/Sysleec/auth/pkg/user_v1"
)

// Get gets a user by id
func (s *Server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	user, err := s.userService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	fmt.Printf("got user: %+v\n", user)

	return converter.ToUserFromService(user), nil
}
