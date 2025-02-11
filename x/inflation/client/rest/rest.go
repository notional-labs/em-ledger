// This software is Copyright (c) 2019-2020 e-Money A/S. It is not offered under an open source license.
//
// Please contact partners@e-money.com for licensing related questions.

package rest

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gorilla/mux"
)

// RegisterRoutes registers minting module REST handlers on the provided router.
func RegisterRoutes(cliCtx client.Context, r *mux.Router) {
	r.HandleFunc(
		"/inflation/current", queryInflationHandlerFn(cliCtx),
	).Methods("GET")
}
