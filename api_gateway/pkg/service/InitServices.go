package service

import (
	"X-TENTIONCREW/api_gateway/pkg/config"
	"X-TENTIONCREW/api_gateway/pkg/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Clients struct {
	Authcli pb.AuthServiceClient
}

func InitClient(c *config.Config) (Clients, error) {
	authcc, autherr := grpc.Dial(c.AuthSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if autherr != nil {
		return Clients{}, autherr
	}
	authclient := pb.NewAuthServiceClient(authcc)
	return Clients{
		Authcli: authclient,
	}, nil
}
