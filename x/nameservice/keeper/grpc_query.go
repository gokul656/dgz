package keeper

import (
	"github.com/calibraint/digizanchain/x/nameservice/types"
)

var _ types.QueryServer = Keeper{}
