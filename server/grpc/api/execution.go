package api

import (
	"context"

	"github.com/mesg-foundation/engine/execution"
	"github.com/mesg-foundation/engine/protobuf/acknowledgement"
	"github.com/mesg-foundation/engine/protobuf/api"
	"github.com/mesg-foundation/engine/sdk"
)

// ExecutionServer serve execution functions.
type ExecutionServer struct {
	sdk *sdk.SDK
}

// NewExecutionServer creates a new ExecutionServer.
func NewExecutionServer(sdk *sdk.SDK) *ExecutionServer {
	return &ExecutionServer{sdk: sdk}
}

// Create creates an execution.
func (s *ExecutionServer) Create(ctx context.Context, req *api.CreateExecutionRequest) (*api.CreateExecutionResponse, error) {
	credUsername, credPassphrase, err := GetCredentialFromContext(ctx)
	if err != nil {
		return nil, err
	}
	exec, err := s.sdk.Execution.Create(req, credUsername, credPassphrase)
	if err != nil {
		return nil, err
	}

	return &api.CreateExecutionResponse{
		Hash: exec.Hash,
	}, nil
}

// Get returns execution from given hash.
func (s *ExecutionServer) Get(ctx context.Context, req *api.GetExecutionRequest) (*execution.Execution, error) {
	return s.sdk.Execution.Get(req.Hash)
}

// Stream returns stream of executions.
func (s *ExecutionServer) Stream(req *api.StreamExecutionRequest, resp api.Execution_StreamServer) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stream, errC, err := s.sdk.Execution.Stream(ctx, req)
	if err != nil {
		return err
	}

	if err := acknowledgement.SetStreamReady(resp); err != nil {
		return err
	}

	for {
		select {
		case exec := <-stream:
			if err := resp.Send(exec); err != nil {
				return err
			}
		case err := <-errC:
			return err
		case <-resp.Context().Done():
			return resp.Context().Err()
		}
	}
}

// Update updates execution from given hash.
func (s *ExecutionServer) Update(ctx context.Context, req *api.UpdateExecutionRequest) (*api.UpdateExecutionResponse, error) {
	credUsername, credPassphrase, err := GetCredentialFromContext(ctx)
	if err != nil {
		return nil, err
	}
	if _, err := s.sdk.Execution.Update(req, credUsername, credPassphrase); err != nil {
		return nil, err
	}
	return &api.UpdateExecutionResponse{}, nil
}

// List returns all executions.
func (s *ExecutionServer) List(ctx context.Context, req *api.ListExecutionRequest) (*api.ListExecutionResponse, error) {
	executions, err := s.sdk.Execution.List()
	if err != nil {
		return nil, err
	}
	return &api.ListExecutionResponse{Executions: executions}, nil
}