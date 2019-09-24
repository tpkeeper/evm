package main

import (
	"github.com/tpkeeper/evm/common"
	"math/big"
)

type EthChainContext struct {
	author common.Address
	header *Header
}

func (eth *EthChainContext)Author(header *Header) (common.Address, error) {
	return eth.author,nil
}

func (eth *EthChainContext)GetHeader(hash common.Hash,uint642 uint64) *Header {
	return eth.header
}


type TxMessage struct {
	to         *common.Address
	from       common.Address
	nonce      uint64
	amount     *big.Int
	gasLimit   uint64
	gasPrice   *big.Int
	data       []byte
	checkNonce bool
}

func NewTxMessage(from common.Address, to *common.Address, nonce uint64, amount *big.Int, gasLimit uint64,
	gasPrice *big.Int, data []byte, checkNonce bool) TxMessage {
	return TxMessage{
		from:       from,
		to:         to,
		nonce:      nonce,
		amount:     amount,
		gasLimit:   gasLimit,
		gasPrice:   gasPrice,
		data:       data,
		checkNonce: checkNonce,
	}
}

func (m TxMessage) From() common.Address { return m.from }
func (m TxMessage) To() *common.Address  { return m.to }
func (m TxMessage) GasPrice() *big.Int   { return m.gasPrice }
func (m TxMessage) Value() *big.Int      { return m.amount }
func (m TxMessage) Gas() uint64          { return m.gasLimit }
func (m TxMessage) Nonce() uint64        { return m.nonce }
func (m TxMessage) Data() []byte         { return m.data }
func (m TxMessage) CheckNonce() bool     { return m.checkNonce }

type Header struct {
	ParentHash  common.Hash    `json:"parentHash"       gencodec:"required"`
	UncleHash   common.Hash    `json:"sha3Uncles"       gencodec:"required"`
	Coinbase    common.Address `json:"miner"            gencodec:"required"`
	Root        common.Hash    `json:"stateRoot"        gencodec:"required"`
	TxHash      common.Hash    `json:"transactionsRoot" gencodec:"required"`
	ReceiptHash common.Hash    `json:"receiptsRoot"     gencodec:"required"`
	Difficulty  *big.Int       `json:"difficulty"       gencodec:"required"`
	Number      *big.Int       `json:"number"           gencodec:"required"`
	GasLimit    uint64         `json:"gasLimit"         gencodec:"required"`
	GasUsed     uint64         `json:"gasUsed"          gencodec:"required"`
	Time        uint64         `json:"timestamp"        gencodec:"required"`
	Extra       []byte         `json:"extraData"        gencodec:"required"`
	MixDigest   common.Hash    `json:"mixHash"`
	Nonce       BlockNonce     `json:"nonce"`
}

type BlockNonce [8]byte