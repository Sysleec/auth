package user

import (
	"github.com/Sysleec/auth/internal/service"
	desc "github.com/Sysleec/auth/pkg/user_v1"
)

// Server is the user server
type Server struct {
	desc.UnimplementedUserV1Server
	userService service.UserService
}

// NewServer returns a new user server
func NewServer(userService service.UserService) *Server {
	return &Server{userService: userService}
}
