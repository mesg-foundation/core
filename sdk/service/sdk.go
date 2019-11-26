package servicesdk

import (
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/mesg-foundation/engine/cosmos"
	"github.com/mesg-foundation/engine/hash"
	"github.com/mesg-foundation/engine/protobuf/api"
	accountsdk "github.com/mesg-foundation/engine/sdk/account"
	"github.com/mesg-foundation/engine/service"
)

// SDK is the service sdk.
type SDK struct {
	accountSDK *accountsdk.SDK
	client     *cosmos.Client
}

// New returns the service sdk.
func New(client *cosmos.Client, accountSDK *accountsdk.SDK) *SDK {
	sdk := &SDK{
		accountSDK: accountSDK,
		client:     client,
	}
	return sdk
}

// Create creates a new service from definition.
func (s *SDK) Create(req *api.CreateServiceRequest, accountName, accountPassword string) (*service.Service, error) {
	account, err := s.accountSDK.Get(accountName)
	if err != nil {
		return nil, err
	}
	// TODO: pass account by parameters
	owner, err := cosmostypes.AccAddressFromBech32(account.Address)
	if err != nil {
		return nil, err
	}
	msg := newMsgCreateService(req, owner)
	tx, err := s.client.BuildAndBroadcastMsg(msg, accountName, accountPassword)
	if err != nil {
		return nil, err
	}
	return s.Get(tx.Data)
}

// Get returns the service that matches given hash.
func (s *SDK) Get(hash hash.Hash) (*service.Service, error) {
	var service service.Service
	if err := s.client.Query("custom/"+backendName+"/get/"+hash.String(), nil, &service); err != nil {
		return nil, err
	}
	return &service, nil
}

// List returns all services.
func (s *SDK) List() ([]*service.Service, error) {
	var services []*service.Service
	if err := s.client.Query("custom/"+backendName+"/list", nil, &services); err != nil {
		return nil, err
	}
	return services, nil
}

// Exists returns if a service already exists.
func (s *SDK) Exists(hash hash.Hash) (bool, error) {
	var exists bool
	if err := s.client.Query("custom/"+backendName+"/exists/"+hash.String(), nil, &exists); err != nil {
		return false, err
	}
	return exists, nil
}

// Hash returns the calculate hash of a service.
func (s *SDK) Hash(req *api.CreateServiceRequest) (hash.Hash, error) {
	var h hash.Hash
	if err := s.client.Query("custom/"+backendName+"/hash", req, &h); err != nil {
		return nil, err
	}
	return h, nil
}