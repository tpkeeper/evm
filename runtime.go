package main

import (
	"github.com/tpkeeper/evm/common"
	"github.com/tpkeeper/evm/common/math"
	"github.com/tpkeeper/evm/params"
	"github.com/tpkeeper/evm/vm"
	"math/big"
	"time"
)

func CreateLogTracer() *vm.StructLogger {
	logConf := vm.LogConfig{
		DisableMemory:  false,
		DisableStack:   false,
		DisableStorage: false,
		Debug:          false,
		Limit:          0,
	}
	return vm.NewStructLogger(&logConf)

}
func CreateChainConfig() *params.ChainConfig {
	chainCfg := params.ChainConfig{
		ChainID:        big.NewInt(1),
		HomesteadBlock: new(big.Int),
		DAOForkBlock:   new(big.Int),
		DAOForkSupport: false,
		EIP150Block:    new(big.Int),
		EIP155Block:    new(big.Int),
		EIP158Block:    new(big.Int),
	}
	return &chainCfg
}
func CreateExecuteContext(caller common.Address) vm.Context {
	context := vm.Context{
		Origin:      caller,
		GasPrice:    new(big.Int),
		Coinbase:    common.BytesToAddress([]byte("coinbase")),
		GasLimit:    math.MaxUint64,
		BlockNumber: new(big.Int),
		Time:        big.NewInt(time.Now().Unix()),
		Difficulty:  new(big.Int),
		CanTransfer: func(sd vm.StateDB, add common.Address, b *big.Int) bool { return true },
		Transfer:    func(sd vm.StateDB, add common.Address, addr common.Address, b *big.Int){},
	}
	return context
}
func CreateVMDefaultConfig() vm.Config {
	return vm.Config{
		Debug:                   true,
		Tracer:                  CreateLogTracer(),
		NoRecursion:             false,
		EnablePreimageRecording: false,
	}

}
func CreateExecuteRuntime(caller common.Address) *vm.EVM {
	context := CreateExecuteContext(caller)
	stateDB := MakeNewMockStateDB()
	chainConfig := CreateChainConfig()
	vmConfig := CreateVMDefaultConfig()

	evm := vm.NewEVM(context, stateDB, chainConfig, vmConfig)
	return evm
}
