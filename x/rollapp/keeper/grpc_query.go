package keeper

import (
	"github.com/gridfoundation/gridiron/x/rollapp/types"
)

var _ types.QueryServer = Keeper{}
