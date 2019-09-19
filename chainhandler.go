package main

import (
	"github.com/tpkeeper/evm/kernel"
)

type ETHChainHandler struct{}

func (ethChainHandler *ETHChainHandler) GetBlockHeaderHash(uint64) kernel.Hash {
	//just return a fake value
	return kernel.HexToHash("this is a demo")
}
