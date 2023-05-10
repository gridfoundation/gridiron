package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/gridfoundation/gridiron/testutil/keeper"
	"github.com/gridfoundation/gridiron/x/sequencer/keeper"
	"github.com/gridfoundation/gridiron/x/sequencer/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.SequencerKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
