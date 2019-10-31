package evm

import (
	"github.com/tpkeeper/evm/common"
	"github.com/tpkeeper/evm/vm"
	"github.com/tpkeeper/evm/params"
	"github.com/tpkeeper/evm/rawdb"
	"github.com/tpkeeper/evm/state"
	"github.com/tpkeeper/evm/types"
	"math/big"
	"testing"
)

type chainContextTest struct{
}


func (c chainContextTest)GetHeader(common.Hash, uint64) *types.Header  {
	return &types.Header{}
}


func TestApplyTransaction(t *testing.T) {
	db,err:=rawdb.NewLevelDBDatabase("testdata",5,5,"world")
	if err!=nil{
		t.Error(err)
	}
	//root :=common.Hash{222 ,136 ,59 ,116, 50 ,27, 206, 138, 8, 57, 53, 52, 85, 228, 131, 215, 182, 208, 127, 142, 223, 26 ,103 ,19, 163, 185, 115, 235, 232, 26, 166 ,122}
	stateDb, err := state.New(common.HexToHash("de883b74321bce8a0839353455e483d7b6d07f8edf1a6713a3b973ebe81aa67a"), state.NewDatabaseWithCache(db, 0))


	balance:=stateDb.GetBalance(common.BytesToAddress([]byte("contract")))
	t.Log("balance",balance)

	stateDb.SetBalance(common.BytesToAddress([]byte("contract")),big.NewInt(5000000))



	chainConfig := params.ChainConfig{
		EIP150Block:    new(big.Int),
		EIP155Block:    new(big.Int),
		EIP158Block:    new(big.Int),
	}

	author:=common.BytesToAddress([]byte("author"))
	header:=types.Header{Number:big.NewInt(5),Difficulty:big.NewInt(5),GasLimit:500000}

	tx:=types.NewTransaction(
		0,
		common.BytesToAddress([]byte("contract")),
		common.HexToAddress("095e7baea6a6c7c4c2dfeb977efac326af552d87"),
		big.NewInt(0),
		uint64(100000),
		big.NewInt(1),
		nil,
	)

	gasPool := new(GasPool).AddGas(header.GasLimit)
	vmConfig:=vm.Config{}
	if err!=nil{
		t.Error(err)
	}

	chainContext:= chainContextTest{}

	usedGas:=uint64(50)
	receipt,err:=ApplyTransaction(&chainConfig,chainContext,&author,gasPool,stateDb,&header,tx,&usedGas,vmConfig)
	if err!=nil{
		t.Fatal(err)
	}
	root:=stateDb.IntermediateRoot(false)
	t.Log("root",root.String())
	stateDb.Commit(false)
	stateDb.Database().TrieDB().Commit(root,true)
	t.Log(receipt)


}
