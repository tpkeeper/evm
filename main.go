package main

import (
	"fmt"
	"github.com/tpkeeper/evm/common"
	"github.com/tpkeeper/evm/vm"
	"math/big"
)

func main() {
	HexTestCode := "6060604052600a8060106000396000f360606040526008565b00"

	TestInput := []byte("Contract")
	TestCallerAddress := []byte("TestAddress")
	TestContractAddress := []byte("TestContract")

	callerAddress := common.BytesToAddress(TestCallerAddress)
	contractAddress := common.BytesToAddress(TestContractAddress)

	//prepare runtime
	evm := CreateExecuteRuntime(callerAddress,&contractAddress,common.Hex2Bytes(HexTestCode))

	//create account
	evm.StateDB.CreateAccount(contractAddress)

	//create contract
	evm.StateDB.SetCode(contractAddress, common.Hex2Bytes(HexTestCode))

	//execute contract
	ret, _, err := evm.Call(
		vm.AccountRef(evm.Origin),
		contractAddress,
		TestInput,
		evm.GasLimit,
		new(big.Int))

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(ret)

}
