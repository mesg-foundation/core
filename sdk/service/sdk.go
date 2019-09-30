package servicesdk

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"github.com/mesg-foundation/engine/cosmos"
	"github.com/mesg-foundation/engine/hash"
	"github.com/mesg-foundation/engine/protobuf/api"
	accountsdk "github.com/mesg-foundation/engine/sdk/account"
	"github.com/mesg-foundation/engine/service"
	cmn "github.com/tendermint/tendermint/libs/common"
)

// SDK is the service sdk.
type SDK struct {
	cdc        *codec.Codec
	accountSDK *accountsdk.SDK
	client     *cosmos.Client
}

// NewSDK returns the service sdk.
func New(cdc *codec.Codec, client *cosmos.Client, accountSDK *accountsdk.SDK) Service {
	sdk := &SDK{
		cdc:        cdc,
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
	accNumber, accSeq := uint64(0), uint64(0)
	owner, err := cosmostypes.AccAddressFromBech32(account.Address)
	if err != nil {
		return nil, err
	}
	msg := newMsgCreateService(s.cdc, req, owner)
	tx, err := s.client.BuildAndBroadcastMsg(msg, accountName, accountPassword, accNumber, accSeq)
	if err != nil {
		return nil, err
	}
	return s.Get(tx.Data)
}

// Delete deletes the service by hash.
func (s *SDK) Delete(hash hash.Hash, accountName, accountPassword string) error {
	// TODO: pass account by parameters
	accNumber, accSeq := uint64(0), uint64(0)
	msg := newMsgDeleteService(s.cdc, hash)
	_, err := s.client.BuildAndBroadcastMsg(msg, accountName, accountPassword, accNumber, accSeq)
	return err
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
func (s *SDK) Exists(req *api.CreateServiceRequest) (*api.ExistsServiceResponse, error) {
	var response api.ExistsServiceResponse
	b, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}
	if err := s.client.Query("custom/"+backendName+"/exists", cmn.HexBytes(b), &response); err != nil {
		return nil, err
	}
	return &response, nil
}
