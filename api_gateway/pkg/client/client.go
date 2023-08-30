package client

import (
	"X-TENTIONCREW/api_gateway/pkg/client/interfaces"
	"X-TENTIONCREW/api_gateway/pkg/pb"
	"X-TENTIONCREW/api_gateway/pkg/service"
	"X-TENTIONCREW/api_gateway/pkg/utils"
	"context"
)

type authClient struct {
	Client pb.AuthServiceClient
}

func NewauthClient(service service.Clients) interfaces.AuthClient {
	return &authClient{
		Client: service.Authcli,
	}
}

func (cr *authClient) Register(ctx context.Context, body utils.SignUpBody) (*pb.RegisterResponse, error) {
	res, err := cr.Client.Register(ctx, &pb.RegisterRequest{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Phone:     body.Phone,
	})
	if err != nil {
		return res, err
	}
	return res, nil
}
