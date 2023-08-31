package interfaces

import (
	"api_gateway/pkg/pb"
	"api_gateway/pkg/utils"
	"context"
)

type MethodClient interface {
	MethodService(context.Context, utils.MethodsRequest) (*pb.MethodResponse, error)
}
