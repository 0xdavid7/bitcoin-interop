package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/axelarnetwork/axelar-core/x/btc/types"
)

// InitGenesis initializes the state from a genesis file
func (k BaseKeeper) InitGenesis(ctx sdk.Context, state types.GenesisState) {
	// TODO: add btc genesis
}

// ExportGenesis generates a genesis file from the state
func (k BaseKeeper) ExportGenesis(ctx sdk.Context) types.GenesisState {
	return *types.NewGenesisState(k.getParams(ctx))
}

func (k BaseKeeper) getParams(ctx sdk.Context) types.Params {
	// TODO: add btc chains params
	return types.Params{}
}
