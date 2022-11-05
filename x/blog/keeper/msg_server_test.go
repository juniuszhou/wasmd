package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/CosmWasm/wasmd/x/blog/types"
    "github.com/CosmWasm/wasmd/x/blog/keeper"
    keepertest "github.com/CosmWasm/wasmd/testutil/keeper"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.BlogKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
