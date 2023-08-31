package interfaces

import "method_svc/pkg/pb"

type AuthClient interface {
	GetAll(flag int32) (*pb.MethodResponse, error)
}
