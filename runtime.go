package main

import (
	"github.com/tpkeeper/evm/common"
	"github.com/tpkeeper/evm/common/math"
	"github.com/tpkeeper/evm/params"
	"github.com/tpkeeper/evm/vm"
	"math/big"
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



func CreateVMDefaultConfig() vm.Config {
	return vm.Config{
		Debug:                   true,
		Tracer:                  CreateLogTracer(),
		NoRecursion:             false,
		EnablePreimageRecording: false,
	}

}
func CreateExecuteRuntime(caller common.Address,to *common.Address,data []byte) *vm.EVM {
	msg:=NewTxMessage(caller,to,0,new(big.Int),math.MaxUint64,big.NewInt(1),data,true)

	header:=Header{Number:new(big.Int).SetUint64(1),Difficulty:big.NewInt(6),GasLimit:math.MaxUint64}

	ethChainContext:=EthChainContext{}

	evmContext:=NewEVMContext(msg,&header,&ethChainContext,nil)


	stateDB := MakeNewMockStateDB()
	chainConfig := CreateChainConfig()
	vmConfig := CreateVMDefaultConfig()

	evm := vm.NewEVM(evmContext, stateDB, chainConfig, vmConfig)
	return evm
}





// Message represents a message sent to a contract.
type Message interface {
	From() common.Address
	//FromFrontier() (common.Address, error)
	To() *common.Address

	GasPrice() *big.Int
	Gas() uint64
	Value() *big.Int

	Nonce() uint64
	CheckNonce() bool
	Data() []byte
}





// NewEVMContext creates a new context for use in the EVM.
func NewEVMContext(msg Message, header *Header, chain ChainContext, author *common.Address) vm.Context {
	// If we don't have an explicit author (i.e. not mining), extract from the header
	var beneficiary common.Address
	if author == nil {
		beneficiary, _ = chain.Author(header) // Ignore error, we're past header validation
	} else {
		beneficiary = *author
	}
	return vm.Context{
		CanTransfer: CanTransfer,
		Transfer:    Transfer,
		GetHash:     GetHashFn(header,chain),
		Origin:      msg.From(),
		Coinbase:    beneficiary,
		BlockNumber: new(big.Int).Set(header.Number),
		Time:        new(big.Int).SetUint64(header.Time),
		Difficulty:  new(big.Int).Set(header.Difficulty),
		GasLimit:    header.GasLimit,
		GasPrice:    new(big.Int).Set(msg.GasPrice()),
	}
}

// CanTransfer checks whether there are enough funds in the address' account to make a transfer.
// This does not take the necessary gas in to account to make the transfer valid.
func CanTransfer(db vm.StateDB, addr common.Address, amount *big.Int) bool {
	return db.GetBalance(addr).Cmp(amount) >= 0
}

// Transfer subtracts amount from sender and adds amount to recipient using the given Db
func Transfer(db vm.StateDB, sender, recipient common.Address, amount *big.Int) {
	db.SubBalance(sender, amount)
	db.AddBalance(recipient, amount)
}

// ChainContext supports retrieving headers and consensus parameters from the
// current blockchain to be used during transaction processing.
type ChainContext interface {
	// Author retrieves the Ethereum address of the account that minted the given
	// block, which may be different from the header's coinbase if a consensus
	// engine is based on signatures.
	Author(header *Header) (common.Address, error)
	// GetHeader returns the hash corresponding to their hash.
	GetHeader(common.Hash, uint64) *Header
}



// GetHashFn returns a GetHashFunc which retrieves header hashes by number
func GetHashFn(ref *Header, chain ChainContext) func(n uint64) common.Hash {
	var cache map[uint64]common.Hash

	return func(n uint64) common.Hash {
		// If there's no hash cache yet, make one
		if cache == nil {
			cache = map[uint64]common.Hash{
				ref.Number.Uint64() - 1: ref.ParentHash,
			}
		}
		// Try to fulfill the request from the cache
		if hash, ok := cache[n]; ok {
			return hash
		}
		// Not cached, iterate the blocks and cache the hashes
		for header := chain.GetHeader(ref.ParentHash, ref.Number.Uint64()-1); header != nil; header = chain.GetHeader(header.ParentHash, header.Number.Uint64()-1) {
			cache[header.Number.Uint64()-1] = header.ParentHash
			if n == header.Number.Uint64()-1 {
				return header.ParentHash
			}
		}
		return common.Hash{}
	}
}