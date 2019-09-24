package main

import (
	"github.com/tpkeeper/evm/common"
	"math/big"
)

type MockStateDB struct {
	stateStore map[common.Address][]byte
}

func MakeNewMockStateDB() *MockStateDB {
	mockstatedb := new(MockStateDB)
	mockstatedb.stateStore = make(map[common.Address][]byte)
	return mockstatedb
}

func (MockStateDB) CreateAccount(common.Address)           {}
func (MockStateDB) SubBalance(common.Address, *big.Int)    {}
func (MockStateDB) AddBalance(common.Address, *big.Int)    {}
func (MockStateDB) GetBalance(common.Address) *big.Int     { return big.NewInt(1000) }
func (MockStateDB) GetNonce(common.Address) uint64         { return 0 }
func (MockStateDB) SetNonce(common.Address, uint64)        {}
func (MockStateDB) GetCodeHash(common.Address) common.Hash { return common.Hash{} }
func (mockstatedb MockStateDB) GetCode(address common.Address) []byte {
	_, ok := mockstatedb.stateStore[address]
	if ok {
		return mockstatedb.stateStore[address]
	} else {
		return nil
	}
}
func (mockstatedb MockStateDB) SetCode(address common.Address, data []byte) {
	mockstatedb.stateStore[address] = data
}
func (mockstatedb MockStateDB) GetCodeSize(address common.Address) int {
	_, ok := mockstatedb.stateStore[address]
	if ok {
		return len(mockstatedb.stateStore[address])
	} else {
		return 0
	}
}
func (MockStateDB) AddRefund(uint64)                                  {}
func (MockStateDB) GetRefund() uint64                                 { return 0 }



func (MockStateDB) SubRefund(uint64){}

func (MockStateDB) GetCommittedState(common.Address, common.Hash) common.Hash{return common.Hash{}}
func (MockStateDB) GetState(common.Address, common.Hash) common.Hash  { return common.Hash{} }
func (MockStateDB) SetState(common.Address, common.Hash, common.Hash) {}
func (MockStateDB) Suicide(common.Address) bool                       { return false }
func (MockStateDB) HasSuicided(common.Address) bool                   { return false }
func (MockStateDB) Exist(common.Address) bool {
	return true
}
func (MockStateDB) Empty(common.Address) bool                                          { return false }
func (MockStateDB) RevertToSnapshot(int)                                               {}
func (MockStateDB) Snapshot() int                                                      { return 0 }
func (MockStateDB) AddLog(*common.Log)                                                 {}
func (MockStateDB) AddPreimage(common.Hash, []byte)                                    {}
//func (MockStateDB) ForEachStorage(common.Address, func(common.Hash, common.Hash) bool) {}
func (MockStateDB) HaveSufficientBalance(common.Address, *big.Int) bool {
	return true
}
func (MockStateDB) TransferBalance(common.Address, common.Address, *big.Int) {

}
func (MockStateDB) ForEachStorage(common.Address, func(common.Hash, common.Hash) bool) error{
	return nil
}