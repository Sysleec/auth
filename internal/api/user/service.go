package user

import (
	"github.com/Sysleec/auth/internal/service"
	desc "github.com/Sysleec/auth/pkg/user_v1"
)

// Implementation is the user server
type Implementation struct {
	desc.UnimplementedUserV1Server
	userService service.UserService
}

// NewImplementation returns a new user server
func NewImplementation(userService service.UserService) *Implementation {
	return &Implementation{userService: userService}
}
