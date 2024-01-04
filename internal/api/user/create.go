package user

import (
	"context"
	"log"

	"github.com/Sysleec/auth/internal/converter"
	desc "github.com/Sysleec/auth/pkg/user_v1"
)

// Create creates a new user
func (s *Server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := s.userService.Create(ctx, converter.ToUserFromDesc(req))
	if err != nil {
		return nil, err
	}
	log.Printf("created user with id %d\n", id)

	return &desc.CreateResponse{Id: id}, nil
}
