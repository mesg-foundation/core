package grpc

import (
	"context"
	"errors"

	protobuf_api "github.com/mesg-foundation/core/protobuf/api"
	"github.com/mesg-foundation/core/protobuf/definition"
	"github.com/mesg-foundation/core/sdk"
	"github.com/mesg-foundation/core/server/grpc/api"
	"github.com/mr-tron/base58"
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
func (s *ServiceServer) Create(ctx context.Context, request *protobuf_api.CreateServiceRequest) (*protobuf_api.CreateServiceResponse, error) {
	srv, err := api.FromProtoService(request.Definition)
	if err != nil {
		return nil, err
	}

	if err := s.sdk.Service.Create(srv); err != nil {
		return nil, err
	}
	return &protobuf_api.CreateServiceResponse{Sid: srv.Sid, Hash: base58.Encode(srv.Hash)}, nil
}

// Delete deletes service by hash or sid.
func (s *ServiceServer) Delete(ctx context.Context, request *protobuf_api.DeleteServiceRequest) (*protobuf_api.DeleteServiceResponse, error) {
	hash, err := base58.Decode(request.Hash)
	if err != nil {
		return nil, err
	}

	srv, err := s.sdk.GetService(hash)
	if err != nil {
		return nil, err
	}
	// first, check if service has any running instances.
	instances, err := s.sdk.Instance.GetAllByService(srv.Hash)
	if err != nil {
		return nil, err
	}
	if len(instances) > 0 {
		return nil, errors.New("service has running instances. in order to delete the service, stop its instances first")
	}
	return &protobuf_api.DeleteServiceResponse{}, s.sdk.Service.Delete(hash)
}

// Get returns service from given hash.
func (s *ServiceServer) Get(ctx context.Context, req *protobuf_api.GetServiceRequest) (*definition.Service, error) {
	hash, err := base58.Decode(req.Hash)
	if err != nil {
		return nil, err
	}

	exec, err := s.sdk.Service.Get(hash)
	if err != nil {
		return nil, err
	}
	return api.ToProtoService(exec), nil
}
