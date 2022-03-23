package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/calibraint/digizanchain/testutil/keeper"
	"github.com/calibraint/digizanchain/x/nameservice/keeper"
	"github.com/calibraint/digizanchain/x/nameservice/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.NameserviceKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
