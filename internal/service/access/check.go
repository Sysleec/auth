package access

import (
	"context"
	"errors"
	"strings"

	"github.com/Sysleec/auth/internal/config/env"
	"github.com/Sysleec/auth/internal/utils"
	"google.golang.org/grpc/metadata"
)

const (
	authPrefix = "Bearer "
)

var accessibleRoles map[string]int32

func (s *servAccess) Check(ctx context.Context, endpointAddress string) error {

	jwt, err := env.NewJWTConfig()
	if err != nil {
		return err
	}

	accessTokenSecretKey := jwt.AccessTokenSecretKey()

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return errors.New("metadata is not provided")
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return errors.New("authorization header is not provided")
	}

	if !strings.HasPrefix(authHeader[0], authPrefix) {
		return errors.New("invalid authorization header format")
	}

	accessToken := strings.TrimPrefix(authHeader[0], authPrefix)

	claims, err := utils.VerifyToken(accessToken, []byte(accessTokenSecretKey))
	if err != nil {
		return errors.New("access token is invalid")
	}

	accessibleMap, err := s.accessibleRoles(ctx)
	if err != nil {
		return errors.New("failed to get accessible roles")
	}

	role, ok := accessibleMap[endpointAddress]
	if !ok {
		return nil
	}

	if role == claims.Role {
		return nil
	}

	return errors.New("access denied")
}

// Return map with endpoint address and role that has access to it
func (s *servAccess) accessibleRoles(ctx context.Context) (map[string]int32, error) {
	if accessibleRoles == nil {
		Roles, err := s.accessRepo.GetAccessibleRoles(ctx)
		if err != nil {
			return nil, nil
		}
		accessibleRoles = Roles
	}

	return accessibleRoles, nil
}

func accessToken(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("metadat is not provided")
	}
	authHeader, ok := md["authorization"]

	if !ok || len(authHeader) == 0 {
		return "", errors.New("authorization header is not provided")
	}

	if !strings.HasPrefix(authHeader[0], authPrefix) {
		return "", errors.New("invalid authrization header format")
	}

	accesToken := strings.TrimPrefix(authHeader[0], authPrefix)
	return accesToken, nil
}
