package converter

import (
	"github.com/Sysleec/auth/internal/model"
	desc "github.com/Sysleec/auth/pkg/auth_v1"
)

func ToUserClaimsFromLogin(req *desc.LoginRequest) *model.UserClaims {
	return &model.UserClaims{
		Username: req.GetUsername(),
	}
}
