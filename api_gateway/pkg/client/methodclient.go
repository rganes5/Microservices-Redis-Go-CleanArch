package client

import (
	"X-TENTIONCREW/api_gateway/pkg/client/interfaces"
	"X-TENTIONCREW/api_gateway/pkg/pb"
	"X-TENTIONCREW/api_gateway/pkg/service"
	"X-TENTIONCREW/api_gateway/pkg/utils"
	"context"
)

type MethodClient struct {
	Client pb.MethodServiceClient
}

func NewMethodClient(service service.Clients) interfaces.MethodClient {
	return &MethodClient{
		Client: service.MethodCli,
	}
}

func (cr *MethodClient) MethodService(ctx context.Context, body utils.MethodsRequest) (*pb.MethodResponse, error) {
	res, err := cr.Client.Method(ctx, &pb.MethodRequest{
		Method:   body.Method,
		WaitTime: body.WaitTime,
	})
	if err != nil {
		return res, err
	}
	return res, nil
}
