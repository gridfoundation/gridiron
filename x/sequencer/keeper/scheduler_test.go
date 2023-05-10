package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/gridfoundation/gridiron/testutil/keeper"
	"github.com/gridfoundation/gridiron/testutil/nullify"
	"github.com/gridfoundation/gridiron/x/sequencer/keeper"
	"github.com/gridfoundation/gridiron/x/sequencer/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNScheduler(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Scheduler {
	items := make([]types.Scheduler, n)
	for i := range items {
		items[i].SequencerAddress = strconv.Itoa(i)

		keeper.SetScheduler(ctx, items[i])
	}
	return items
}

func TestSchedulerGet(t *testing.T) {
	keeper, ctx := keepertest.SequencerKeeper(t)
	items := createNScheduler(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetScheduler(ctx,
			item.SequencerAddress,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestSchedulerRemove(t *testing.T) {
	keeper, ctx := keepertest.SequencerKeeper(t)
	items := createNScheduler(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveScheduler(ctx,
			item.SequencerAddress,
		)
		_, found := keeper.GetScheduler(ctx,
			item.SequencerAddress,
		)
		require.False(t, found)
	}
}

func TestSchedulerGetAll(t *testing.T) {
	keeper, ctx := keepertest.SequencerKeeper(t)
	items := createNScheduler(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllScheduler(ctx)),
	)
}
