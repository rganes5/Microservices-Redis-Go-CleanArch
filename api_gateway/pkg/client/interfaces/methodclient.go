package interfaces

import (
	"X-TENTIONCREW/api_gateway/pkg/pb"
	"X-TENTIONCREW/api_gateway/pkg/utils"
	"context"
)

type MethodClient interface {
	MethodService(context.Context, utils.MethodsRequest) (*pb.MethodResponse, error)
}
