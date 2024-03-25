package access

import (
	"context"

	desc "github.com/Sysleec/auth/pkg/access_v1"
	"github.com/golang/protobuf/ptypes/empty"
)

func (i *Implementation) GetName(ctx context.Context, _ *empty.Empty) (*desc.GetNameResponse, error) {
	name, err := i.accessService.GetName(ctx)
	if err != nil {
		return nil, err
	}
	return &desc.GetNameResponse{
		Name: name,
	}, nil
}
