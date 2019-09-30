package servicesdk

import (
	"errors"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"github.com/mesg-foundation/engine/cosmos"
	"github.com/mesg-foundation/engine/database"
	"github.com/mesg-foundation/engine/database/store"
	"github.com/mesg-foundation/engine/hash"
	"github.com/mesg-foundation/engine/protobuf/api"
	ownershipsdk "github.com/mesg-foundation/engine/sdk/ownership"
	"github.com/mesg-foundation/engine/service"
	"github.com/mesg-foundation/engine/service/validator"
	abci "github.com/tendermint/tendermint/abci/types"
)

const backendName = "service"

// Backend is the service backend.
type Backend struct {
	cdc        *codec.Codec
	storeKey   *cosmostypes.KVStoreKey
	ownerships *ownershipsdk.Backend
}

// NewBackend returns the backend of the service sdk.
func NewBackend(appFactory *cosmos.AppFactory, ownerships *ownershipsdk.Backend) *Backend {
	backend := &Backend{
		cdc:        appFactory.Cdc(),
		storeKey:   cosmostypes.NewKVStoreKey(backendName),
		ownerships: ownerships,
	}
	appBackendBasic := cosmos.NewAppModuleBasic(backendName)
	appBackend := cosmos.NewAppModule(appBackendBasic, backend.cdc, backend.handler, backend.querier)
	appFactory.RegisterModule(appBackend)
	appFactory.RegisterStoreKey(backend.storeKey)

	backend.cdc.RegisterConcrete(msgCreateService{}, "service/create", nil)
	backend.cdc.RegisterConcrete(msgDeleteService{}, "service/delete", nil)

	return backend
}

func (s *Backend) db(request cosmostypes.Request) *database.ServiceDB {
	return database.NewServiceDB(store.NewCosmosStore(request.KVStore(s.storeKey)), s.cdc)
}

func (s *Backend) handler(request cosmostypes.Request, msg cosmostypes.Msg) cosmostypes.Result {
	switch msg := msg.(type) {
	case msgCreateService:
		srv, err := s.Create(request, &msg)
		if err != nil {
			return cosmostypes.ErrInternal(err.Error()).Result()
		}
		return cosmostypes.Result{
			Data: srv.Hash,
		}
	case msgDeleteService:
		err := s.Delete(request, msg.Hash)
		if err != nil {
			return cosmostypes.ErrInternal(err.Error()).Result()
		}
		return cosmostypes.Result{}
	default:
		errmsg := fmt.Sprintf("Unrecognized service Msg type: %v", msg.Type())
		return cosmostypes.ErrUnknownRequest(errmsg).Result()
	}
}

func (s *Backend) querier(request cosmostypes.Request, path []string, req abci.RequestQuery) (interface{}, error) {
	switch path[0] {
	case "get":
		hash, err := hash.Decode(path[1])
		if err != nil {
			return nil, err
		}
		return s.Get(request, hash)
	case "list":
		return s.List(request)
	case "exists":
		var createServiceRequest api.CreateServiceRequest
		if err := proto.Unmarshal(req.Data, &createServiceRequest); err != nil {
			return nil, err
		}
		return s.Exists(request, &createServiceRequest)
	default:
		return nil, errors.New("unknown service query endpoint" + path[0])
	}
}

// Create creates a new service from definition.
func (s *Backend) Create(request cosmostypes.Request, msg *msgCreateService) (*service.Service, error) {
	return create(s.db(request), msg.Request, msg.Owner, s.ownerships, request)
}

// Delete deletes the service by hash.
func (s *Backend) Delete(request cosmostypes.Request, hash hash.Hash) error {
	return s.db(request).Delete(hash)
}

// Get returns the service that matches given hash.
func (s *Backend) Get(request cosmostypes.Request, hash hash.Hash) (*service.Service, error) {
	return s.db(request).Get(hash)
}

// Exists returns true if a specific set of data exists in the database, false otherwise
func (s *Backend) Exists(request cosmostypes.Request, req *api.CreateServiceRequest) (*api.ExistsServiceResponse, error) {
	h := initializeService(req).Hash
	exists, err := s.db(request).Exist(h)
	if err != nil {
		return nil, err
	}
	response := &api.ExistsServiceResponse{
		Exists: exists,
	}
	if response.Exists {
		response.Hash = h
	}
	return response, nil
}

// List returns all services.
func (s *Backend) List(request cosmostypes.Request) ([]*service.Service, error) {
	return s.db(request).All()
}

func create(db *database.ServiceDB, req *api.CreateServiceRequest, owner cosmostypes.AccAddress, ownerships *ownershipsdk.Backend, request cosmostypes.Request) (*service.Service, error) {
	// create service
	srv := initializeService(req)

	// check if service already exists.
	if _, err := db.Get(srv.Hash); err == nil {
		return nil, &AlreadyExistsError{Hash: srv.Hash}
	}

	// TODO: the following test should be moved in New function
	if srv.Sid == "" {
		// make sure that sid doesn't have the same length with id.
		srv.Sid = "_" + srv.Hash.String()
	}

	if err := validator.ValidateService(srv); err != nil {
		return nil, err
	}

	if _, err := ownerships.CreateServiceOwnership(request, srv.Hash, owner); err != nil {
		return nil, err
	}

	if err := db.Save(srv); err != nil {
		return nil, err
	}
	return srv, nil
}

func initializeService(req *api.CreateServiceRequest) *service.Service {
	// create service
	srv := &service.Service{
		Sid:           req.Sid,
		Name:          req.Name,
		Description:   req.Description,
		Configuration: req.Configuration,
		Tasks:         req.Tasks,
		Events:        req.Events,
		Dependencies:  req.Dependencies,
		Repository:    req.Repository,
		Source:        req.Source,
	}

	// calculate and apply hash to service.
	srv.Hash = hash.Dump(srv)
	return srv
}
