package api

import (
	"context"

	protobuf_api "github.com/mesg-foundation/engine/protobuf/api"
	"github.com/mesg-foundation/engine/sdk"
	"github.com/mesg-foundation/engine/service"
)

// ServiceServer is the type to aggregate all Service APIs.
type ServiceServer struct {
	sdk *sdk.SDK
}

// NewServiceServer creates a new ServiceServer.
func NewServiceServer(sdk *sdk.SDK) *ServiceServer {
	return &ServiceServer{sdk: sdk}
}

// Create creates a new service from definition.
func (s *ServiceServer) Create(ctx context.Context, req *protobuf_api.CreateServiceRequest) (*protobuf_api.CreateServiceResponse, error) {
	credUsername, credPassphrase, err := GetCredentialFromContext(ctx)
	if err != nil {
		return nil, err
	}
	srv, err := s.sdk.Service.Create(req, credUsername, credPassphrase)
	if err != nil {
		return nil, err
	}
	return &protobuf_api.CreateServiceResponse{Hash: srv.Hash}, nil
}

// Get returns service from given hash.
func (s *ServiceServer) Get(ctx context.Context, req *protobuf_api.GetServiceRequest) (*service.Service, error) {
	return s.sdk.Service.Get(req.Hash)
}

// List returns all services.
func (s *ServiceServer) List(ctx context.Context, req *protobuf_api.ListServiceRequest) (*protobuf_api.ListServiceResponse, error) {
	services, err := s.sdk.Service.List()
	if err != nil {
		return nil, err
	}

	return &protobuf_api.ListServiceResponse{Services: services}, nil
}

// Exists returns if a service already exists.
func (s *ServiceServer) Exists(ctx context.Context, req *protobuf_api.ExistsServiceRequest) (*protobuf_api.ExistsServiceResponse, error) {
	exists, err := s.sdk.Service.Exists(req.Hash)
	if err != nil {
		return nil, err
	}
	return &protobuf_api.ExistsServiceResponse{
		Exists: exists,
	}, nil
}

// Hash returns the calculated hash of a service request.
func (s *ServiceServer) Hash(ctx context.Context, req *protobuf_api.CreateServiceRequest) (*protobuf_api.HashServiceResponse, error) {
	h, err := s.sdk.Service.Hash(req)
	if err != nil {
		return nil, err
	}
	return &protobuf_api.HashServiceResponse{
		Hash: h,
	}, nil
}