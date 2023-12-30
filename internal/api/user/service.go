package user

import (
	"github.com/Sysleec/auth/internal/service"
	desc "github.com/Sysleec/auth/pkg/user_v1"
)

type Server struct {
	desc.UnimplementedUserV1Server
	userService service.UserService
}

func NewServer(userService service.UserService) *Server {
	return &Server{userService: userService}
}
