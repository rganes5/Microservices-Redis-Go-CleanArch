package service

import (
	"api_gateway/pkg/config"
	"api_gateway/pkg/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Clients struct {
	Authcli   pb.AuthServiceClient
	MethodCli pb.MethodServiceClient
}

func InitClient(c *config.Config) (Clients, error) {
	authcc, autherr := grpc.Dial(c.AuthSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if autherr != nil {
		return Clients{}, autherr
	}
	methodcc, methoderr := grpc.Dial(c.MethSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if methoderr != nil {
		return Clients{}, methoderr
	}
	authclient := pb.NewAuthServiceClient(authcc)
	methodclient := pb.NewMethodServiceClient(methodcc)
	return Clients{
		Authcli:   authclient,
		MethodCli: methodclient,
	}, nil
}
