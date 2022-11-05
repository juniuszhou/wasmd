package keeper

import (
	"github.com/CosmWasm/wasmd/x/blog/types"
)

var _ types.QueryServer = Keeper{}
