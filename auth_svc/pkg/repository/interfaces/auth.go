package interfaces

import (
	"X-TENTIONCREW/auth_svc/pkg/pb"
	"X-TENTIONCREW/auth_svc/pkg/utils"
	"context"
)

type AuthRepo interface {
	Register(context.Context, *pb.RegisterRequest) (int32, error)
	GetUser(ctx context.Context, id int32) (utils.Response, error)
	UpdateUser(context.Context, *pb.UpdateRequest) (utils.Response, error)
	DeleteUser(ctx context.Context, id int32) error
}
