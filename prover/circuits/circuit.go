package age

import (
	"github.com/brevis-network/brevis-sdk/sdk"
)

type AppCircuit struct{}

func (c *AppCircuit) Allocate() (maxReceipts, maxStorage, maxTransactions int) {
	// Our app is only ever going to use one storage data at a time so
	// we can simply limit the max number of data for storage to 1 and
	// 0 for all others
	return 0, 0, 1
}

func (c *AppCircuit) Define(api *sdk.CircuitAPI, in sdk.DataInput) error {
	txs := sdk.NewDataStream(api, in.Transactions)

	tx := sdk.GetUnderlying(txs, 0)
	// This is our main check logic
	//api.Uint248.AssertIsEqual(tx.Nonce, sdk.ConstUint248(0))
	api.Uint248.AssertIsLessOrEqual(sdk.ConstUint248(40077593), tx.BlockNum)

	// Output variables can be later accessed in our app contract
	api.OutputAddress(tx.From)
	api.OutputUint(64, tx.BlockNum)

	return nil
}
