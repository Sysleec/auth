package access

import (
	"context"

	desc "github.com/Sysleec/auth/pkg/access_v1"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Check(ctx context.Context, req *desc.CheckRequest) (*empty.Empty, error) {
	err := i.accessService.Check(ctx, req.GetEndpointAddress())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
