package testchain_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/xiaka53/testchain/testutil/keeper"
	"github.com/xiaka53/testchain/x/testchain"
	"github.com/xiaka53/testchain/x/testchain/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TestchainKeeper(t)
	testchain.InitGenesis(ctx, *k, genesisState)
	got := testchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
