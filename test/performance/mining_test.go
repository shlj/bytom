package performance

import (
	"os"
	"testing"

	dbm "github.com/tendermint/tmlibs/db"

	"github.com/bytom/account"
	"github.com/bytom/config"
	"github.com/bytom/mining"
	"github.com/bytom/test"
)

// Function NewBlockTemplate's benchmark - 0.05s
func BenchmarkNewBlockTpl(b *testing.B) {
	testDB := dbm.NewDB("testdb", "leveldb", "temp")
	defer os.RemoveAll("temp")

	chain, err := test.MockChain(testDB)
	if err != nil {
		b.Fatal(err)
	}
	accountManager := account.NewManager(testDB, chain)

	txPool := test.MockTxPool()
	genesisBlock := config.GenerateGenesisBlock()
	chain.SaveBlock(genesisBlock)
	chain.ConnectBlock(genesisBlock)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mining.NewBlockTemplate(chain, txPool, accountManager)
	}
}
