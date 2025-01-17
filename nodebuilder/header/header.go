package header

import (
	"context"

	"github.com/celestiaorg/celestia-node/header"
)

// Module exposes the functionality needed for querying headers from the network.
// Any method signature changed here needs to also be changed in the API struct.
//
//go:generate mockgen -destination=mocks/api.go -package=mocks . Module
type Module interface {
	// GetByHeight returns the ExtendedHeader at the given height, blocking
	// until header has been processed by the store or context deadline is exceeded.
	GetByHeight(context.Context, uint64) (*header.ExtendedHeader, error)
	// Head returns the ExtendedHeader of the chain head.
	Head(context.Context) (*header.ExtendedHeader, error)
	// IsSyncing returns the status of sync
	IsSyncing() bool
}

// API is a wrapper around Module for the RPC.
// TODO(@distractedm1nd): These structs need to be autogenerated.
type API struct {
	GetByHeight func(context.Context, uint64) (*header.ExtendedHeader, error)
	Head        func(context.Context) (*header.ExtendedHeader, error)
	IsSyncing   func() bool
}
