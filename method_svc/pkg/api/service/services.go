package service

import (
	"X-TENTIONCREW/method_svc/pkg/client"
	"X-TENTIONCREW/method_svc/pkg/pb"
	"context"
	"errors"
	"net/http"
	"sync"
	"time"
)

type MethodService struct {
	AuthClient client.AuthClient
	pb.UnimplementedMethodServiceServer
	method1Mutex sync.Mutex
}

func NewMethodService(AuthClient client.AuthClient) pb.MethodServiceServer {
	return &MethodService{
		AuthClient: AuthClient,
	}
}
func (c *MethodService) Method(ctx context.Context, req *pb.MethodRequest) (*pb.MethodResponse, error) {
	switch req.Method {
	case 1:
		c.method1Mutex.Lock()
		defer c.method1Mutex.Unlock()

		response, err := c.AuthClient.GetAll(1)
		if err != nil {
			return nil, err
		}
		time.Sleep(time.Duration(req.WaitTime) * time.Second)

		// Convert the user names to []*pb.MethodUsers
		var methodUsers []*pb.MethodUsers
		for _, user := range response.Users {
			methodUser := &pb.MethodUsers{
				FirstName: user.FirstName,
			}
			methodUsers = append(methodUsers, methodUser)
		}

		return &pb.MethodResponse{
			Status:   http.StatusOK,
			Response: "Success",
			Count:    response.Count,
			Users:    methodUsers,
		}, nil

	case 2:
		response, err := c.AuthClient.GetAll(1)
		if err != nil {
			return nil, err
		}
		time.Sleep(time.Duration(req.WaitTime) * time.Second)

		// Convert the user names to []*pb.MethodUsers
		var methodUsers []*pb.MethodUsers
		for _, user := range response.Users {
			methodUser := &pb.MethodUsers{
				FirstName: user.FirstName,
			}
			methodUsers = append(methodUsers, methodUser)
		}

		return &pb.MethodResponse{
			Status:   http.StatusOK,
			Response: "Success",
			Count:    response.Count,
			Users:    methodUsers,
		}, nil

	default:
		return &pb.MethodResponse{}, errors.New("wrong method")
	}
}
