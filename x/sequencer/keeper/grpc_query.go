package keeper

import (
	"github.com/gridfoundation/gridiron/x/sequencer/types"
)

var _ types.QueryServer = Keeper{}
