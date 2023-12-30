package converter

import (
	"github.com/Sysleec/auth/internal/model"
	modelRepo "github.com/Sysleec/auth/internal/repository/user/model"
)

// ToUserFromRepo converts user from repository to user for gRPC
func ToUserFromRepo(user *modelRepo.User) *model.User {
	return &model.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      model.Role(user.Role),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
