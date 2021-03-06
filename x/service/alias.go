package service

import (
	"github.com/mesg-foundation/engine/x/service/internal/keeper"
	"github.com/mesg-foundation/engine/x/service/internal/types"
)

// const aliases
const (
	ModuleName   = types.ModuleName
	RouterKey    = types.RouterKey
	StoreKey     = types.StoreKey
	QuerierRoute = types.QuerierRoute
)

// functions and variable aliases
var (
	NewKeeper           = keeper.NewKeeper
	NewQuerier          = keeper.NewQuerier
	RegisterCodec       = types.RegisterCodec
	NewGenesisState     = types.NewGenesisState
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis

	ModuleCdc = types.ModuleCdc

	QueryGet   = types.QueryGet
	QueryList  = types.QueryList
	QueryExist = types.QueryExist

	EventType              = types.EventType
	AttributeKeyHash       = types.AttributeKeyHash
	AttributeKeyAddress    = types.AttributeKeyAddress
	AttributeActionCreated = types.AttributeActionCreated
)

// module types
type (
	Keeper       = keeper.Keeper
	GenesisState = types.GenesisState

	MsgCreate = types.MsgCreate
)
