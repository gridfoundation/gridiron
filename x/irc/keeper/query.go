package keeper

import (
	"github.com/gridfoundation/gridiron/x/irc/types"
)

var _ types.QueryServer = Keeper{}
