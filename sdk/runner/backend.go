package runnersdk

import (
	"errors"
	"fmt"

	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/mesg-foundation/engine/codec"
	"github.com/mesg-foundation/engine/cosmos"
	"github.com/mesg-foundation/engine/hash"
	"github.com/mesg-foundation/engine/runner"
	instancesdk "github.com/mesg-foundation/engine/sdk/instance"
	abci "github.com/tendermint/tendermint/abci/types"
)

const backendName = "runner"

// Backend is the runner backend.
type Backend struct {
	storeKey     *cosmostypes.KVStoreKey
	instanceBack *instancesdk.Backend
}

// NewBackend returns the backend of the runner sdk.
func NewBackend(appFactory *cosmos.AppFactory, instanceBack *instancesdk.Backend) *Backend {
	backend := &Backend{
		storeKey:     cosmostypes.NewKVStoreKey(backendName),
		instanceBack: instanceBack,
	}
	appBackendBasic := cosmos.NewAppModuleBasic(backendName)
	appBackend := cosmos.NewAppModule(appBackendBasic, backend.handler, backend.querier)
	appFactory.RegisterModule(appBackend)
	appFactory.RegisterStoreKey(backend.storeKey)

	return backend
}

func (s *Backend) handler(request cosmostypes.Request, msg cosmostypes.Msg) (hash.Hash, cosmostypes.Error) {
	switch msg := msg.(type) {
	case msgCreateRunner:
		run, err := s.Create(request, &msg)
		if err != nil {
			return run.Hash, cosmostypes.ErrInternal(err.Error())
		}
		return run.Hash, nil
	case msgDeleteRunner:
		if err := s.Delete(request, &msg); err != nil {
			return nil, cosmostypes.ErrInternal(err.Error())
		}
		return nil, nil
	default:
		errmsg := fmt.Sprintf("Unrecognized runner Msg type: %v", msg.Type())
		return nil, cosmostypes.ErrUnknownRequest(errmsg)
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
	default:
		return nil, errors.New("unknown runner query endpoint" + path[0])
	}
}

// Create creates a new runner.
func (s *Backend) Create(request cosmostypes.Request, msg *msgCreateRunner) (*runner.Runner, error) {
	store := request.KVStore(s.storeKey)
	inst, err := s.instanceBack.FetchOrCreate(request, msg.ServiceHash, msg.EnvHash)
	if err != nil {
		return nil, err
	}
	run := &runner.Runner{
		Address:      msg.Address.String(),
		InstanceHash: inst.Hash,
	}
	run.Hash = hash.Dump(run)
	if store.Has(run.Hash) {
		return nil, fmt.Errorf("runner %q already exists", run.Hash)
	}
	value, err := codec.MarshalBinaryBare(run)
	if err != nil {
		return nil, err
	}
	store.Set(run.Hash, value)
	return run, nil
}

// Delete deletes a runner.
func (s *Backend) Delete(request cosmostypes.Request, msg *msgDeleteRunner) error {
	store := request.KVStore(s.storeKey)
	run := runner.Runner{}
	value := store.Get(msg.RunnerHash)
	if err := codec.UnmarshalBinaryBare(value, &run); err != nil {
		return err
	}
	if run.Address != msg.Address.String() {
		return errors.New("only the runner owner can remove itself")
	}
	store.Delete(msg.RunnerHash)
	return nil
}

// Get returns the runner that matches given hash.
func (s *Backend) Get(request cosmostypes.Request, hash hash.Hash) (*runner.Runner, error) {
	store := request.KVStore(s.storeKey)
	if !store.Has(hash) {
		return nil, fmt.Errorf("runner %q not found", hash)
	}
	value := store.Get(hash)
	var run *runner.Runner
	return run, codec.UnmarshalBinaryBare(value, &run)
}

// List returns all runners.
func (s *Backend) List(request cosmostypes.Request) ([]*runner.Runner, error) {
	var (
		runners []*runner.Runner
		iter    = request.KVStore(s.storeKey).Iterator(nil, nil)
	)
	for iter.Valid() {
		var run *runner.Runner
		value := iter.Value()
		if err := codec.UnmarshalBinaryBare(value, &run); err != nil {
			return nil, err
		}
		runners = append(runners, run)
		iter.Next()
	}
	iter.Close()
	return runners, nil
}
