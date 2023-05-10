package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/gridfoundation/gridiron/testutil/keeper"
	"github.com/gridfoundation/gridiron/x/irc/keeper"
	"github.com/gridfoundation/gridiron/x/irc/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, _, ctx := keepertest.IRCKeeper(t)
	return keeper.NewMsgServerImpl(k), sdk.WrapSDKContext(ctx)
}
