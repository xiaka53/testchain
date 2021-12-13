package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "githup.com/xiaka53/testchain/testutil/keeper"
	"githup.com/xiaka53/testchain/x/testchain/keeper"
	"githup.com/xiaka53/testchain/x/testchain/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.TestchainKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
