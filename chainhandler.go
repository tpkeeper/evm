package main

import (
	"github.com/tpkeeper/evm/common"
)

type ETHChainHandler struct{}

func (ethChainHandler *ETHChainHandler) GetBlockHeaderHash(uint64) common.Hash {
	//just return a fake value
	return common.HexToHash("this is a demo")
}
