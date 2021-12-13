package keeper

import (
	"github.com/xiaka53/testchain/x/testchain/types"
)

var _ types.QueryServer = Keeper{}
