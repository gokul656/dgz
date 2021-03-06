package nameservice_test

import (
	"testing"

	keepertest "github.com/calibraint/digizanchain/testutil/keeper"
	"github.com/calibraint/digizanchain/testutil/nullify"
	"github.com/calibraint/digizanchain/x/nameservice"
	"github.com/calibraint/digizanchain/x/nameservice/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.NameserviceKeeper(t)
	nameservice.InitGenesis(ctx, *k, genesisState)
	got := nameservice.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
