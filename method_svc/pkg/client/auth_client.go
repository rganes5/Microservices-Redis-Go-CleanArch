package client

import (
	"context"
	"method_svc/pkg/client/interfaces"
	"method_svc/pkg/config"
	"method_svc/pkg/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthClient struct {
	Server pb.AuthServiceClient
}

func NewAuthClient(server pb.AuthServiceClient) interfaces.AuthClient {
	return &AuthClient{
		Server: server,
	}
}

func InitAuthClient(c *config.Config) (pb.AuthServiceClient, error) {
	authcc, autherr := grpc.Dial(c.AuthSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if autherr != nil {
		return nil, autherr
	}
	return pb.NewAuthServiceClient(authcc), nil
}

func (c *AuthClient) GetAll(flag int32) (*pb.MethodResponse, error) {
	res, err := c.Server.GetAll(context.Background(), &pb.GetAllRequest{
		Flag: flag,
	})
	if err != nil {
		return nil, err
	}
	// Convert []*pb.MethodUser to []*pb.MethodUsers
	var methodUsers []*pb.MethodUsers
	for _, user := range res.Users {
		methodUser := &pb.MethodUsers{
			FirstName: user.FirstName,
		}
		methodUsers = append(methodUsers, methodUser)
	}

	methodRes := &pb.MethodResponse{
		Status:   res.Status,
		Response: res.Response,
		Error:    res.Error,
		Count:    res.Count,
		Users:    methodUsers,
	}

	return methodRes, nil
}
