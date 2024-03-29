package converter

import (
	"github.com/Sysleec/auth/internal/model"
	desc "github.com/Sysleec/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ToUserFromService converts user from service to user for gRPC
func ToUserFromService(user *model.User) *desc.GetResponse {
	return &desc.GetResponse{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      desc.Role(user.Role),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}

// ToUserFromDesc converts user from gRPC to user for service
func ToUserFromDesc(user *desc.CreateRequest) *model.User {
	return &model.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Role:     model.Role(user.Role),
	}
}

func ToUserFromDescUpdate(user *desc.UpdateRequest) *model.User {
	return &model.User{
		ID:    user.Id,
		Name:  user.Name.Value,
		Email: user.Email.Value,
	}
}
