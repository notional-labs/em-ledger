// This software is Copyright (c) 2019-2020 e-Money A/S. It is not offered under an open source license.
//
// Please contact partners@e-money.com for licensing related questions.

package issuer

import (
	"github.com/e-money/em-ledger/x/issuer/keeper"
	types "github.com/e-money/em-ledger/x/issuer/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func initGenesis(ctx sdk.Context, k keeper.Keeper, state types.GenesisState) {
	for _, issuer := range state.Issuers {
		k.AddIssuer(ctx, issuer)
	}
}

func defaultGenesisState() *types.GenesisState {
	return &types.GenesisState{}
}
