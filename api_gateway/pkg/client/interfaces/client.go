package interfaces

import (
	"X-TENTIONCREW/api_gateway/pkg/pb"
	"X-TENTIONCREW/api_gateway/pkg/utils"
	"context"
)

type AuthClient interface {
	Register(context.Context, utils.SignUpBody) (*pb.RegisterResponse, error)
	GetUser(ctx context.Context, userId uint32) (*pb.GetUserResponse, error)
	UpdateUser(context.Context, utils.UpdateBody) (*pb.UpdateResponse, error)
	DeleteUser(ctx context.Context, userId uint32) (*pb.DeleteResponse, error)
}
