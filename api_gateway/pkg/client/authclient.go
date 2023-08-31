package client

import (
	"api_gateway/pkg/client/interfaces"
	"api_gateway/pkg/pb"
	"api_gateway/pkg/service"
	"api_gateway/pkg/utils"
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

func (cr *authClient) GetUser(ctx context.Context, userId uint32) (*pb.GetUserResponse, error) {
	res, err := cr.Client.GetUser(ctx, &pb.GetUserRequest{
		ID: int32(userId),
	})
	if err != nil {
		return res, err
	}
	return res, nil
}

func (cr *authClient) UpdateUser(ctx context.Context, body utils.UpdateBody) (*pb.UpdateResponse, error) {
	res, err := cr.Client.Update(ctx, &pb.UpdateRequest{
		ID:        int32(body.Id),
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

func (cr *authClient) DeleteUser(ctx context.Context, userId uint32) (*pb.DeleteResponse, error) {
	res, err := cr.Client.Delete(ctx, &pb.DeleteRequest{
		ID: int32(userId),
	})
	if err != nil {
		return res, err
	}
	return res, nil
}
