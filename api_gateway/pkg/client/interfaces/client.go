package interfaces

import (
	"X-TENTIONCREW/api_gateway/pkg/pb"
	"X-TENTIONCREW/api_gateway/pkg/utils"
	"context"
)

type AuthClient interface {
	Register(context.Context, utils.SignUpBody) (*pb.RegisterResponse, error)
}
