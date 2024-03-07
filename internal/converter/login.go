package converter

import (
	"github.com/Sysleec/auth/internal/model"
	desc "github.com/Sysleec/auth/pkg/auth_v1"
)

func ToUserClaimsFromLogin(req *desc.LoginRequest) *model.LoginClaims {
	return &model.LoginClaims{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
	}
}
