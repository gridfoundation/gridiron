package keeper_test

import (
	"testing"

	testkeeper "github.com/gridfoundation/gridiron/testutil/keeper"
	"github.com/gridfoundation/gridiron/testutil/sample"
	"github.com/gridfoundation/gridiron/x/rollapp/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.RollappKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params.DisputePeriodInBlocks, k.DisputePeriodInBlocks(ctx))
	require.EqualValues(t, len(params.DeployerWhitelist), len(k.DeployerWhitelist(ctx)))
	for i := range params.DeployerWhitelist {
		require.EqualValues(t, params.DeployerWhitelist[i], k.DeployerWhitelist(ctx)[i])
	}
}

func TestGetParamsWithDeployerWhitelist(t *testing.T) {
	k, ctx := testkeeper.RollappKeeper(t)

	params := types.DefaultParams()
	params.DeployerWhitelist = []types.DeployerParams{{Address: sample.AccAddress(), MaxRollapps: 0}, {Address: sample.AccAddress(), MaxRollapps: 50}}

	k.SetParams(ctx, params)

	require.EqualValues(t, params.DisputePeriodInBlocks, k.DisputePeriodInBlocks(ctx))
	require.EqualValues(t, len(params.DeployerWhitelist), len(k.DeployerWhitelist(ctx)))
	for i := range params.DeployerWhitelist {
		require.EqualValues(t, params.DeployerWhitelist[i], k.DeployerWhitelist(ctx)[i])
	}
}
