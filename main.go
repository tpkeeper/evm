package main

import (
	"fmt"
	"github.com/tpkeeper/evm/common"
	"github.com/tpkeeper/evm/vm"
	"math/big"
)

func main() {
	HexTestCode := "6060604052600a8060106000396000f360606040526008565b00"
	//HexTestCode := "6060604052361561006c5760e060020a600035046308551a53811461007457806335a063b4146100865780633fa4f245146100a6578063590e1ae3146100af5780637150d8ae146100cf57806373fac6f0146100e1578063c19d93fb146100fe578063d696069714610112575b610131610002565b610133600154600160a060020a031681565b610131600154600160a060020a0390811633919091161461015057610002565b61014660005481565b610131600154600160a060020a039081163391909116146102d557610002565b610133600254600160a060020a031681565b610131600254600160a060020a0333811691161461023757610002565b61014660025460ff60a060020a9091041681565b61013160025460009060ff60a060020a9091041681146101cc57610002565b005b600160a060020a03166060908152602090f35b6060908152602090f35b60025460009060a060020a900460ff16811461016b57610002565b600154600160a060020a03908116908290301631606082818181858883f150506002805460a060020a60ff02191660a160020a179055506040517f72c874aeff0b183a56e2b79c71b46e1aed4dee5e09862134b8821ba2fddbf8bf9250a150565b80546002023414806101dd57610002565b6002805460a060020a60ff021973ffffffffffffffffffffffffffffffffffffffff1990911633171660a060020a1790557fd5d55c8a68912e9a110618df8d5e2e83b8d83211c57a8ddd1203df92885dc881826060a15050565b60025460019060a060020a900460ff16811461025257610002565b60025460008054600160a060020a0390921691606082818181858883f150508354604051600160a060020a0391821694503090911631915082818181858883f150506002805460a060020a60ff02191660a160020a179055506040517fe89152acd703c9d8c7d28829d443260b411454d45394e7995815140c8cbcbcf79250a150565b60025460019060a060020a900460ff1681146102f057610002565b6002805460008054600160a060020a0390921692909102606082818181858883f150508354604051600160a060020a0391821694503090911631915082818181858883f150506002805460a060020a60ff02191660a160020a179055506040517f8616bbbbad963e4e65b1366f1d75dfb63f9e9704bbbf91fb01bec70849906cf79250a15056"
	//TestInput := []byte{214, 150, 6, 151}
	TestInput := []byte("Contract")
	TestCallerAddress := []byte("TestAddress")
	TestContractAddress := []byte("TestContract")
	calleraddress := common.BytesToAddress(TestCallerAddress)
	contractaddress := common.BytesToAddress(TestContractAddress)
	evm := CreateExecuteRuntime(calleraddress)
	evm.StateDB.CreateAccount(contractaddress)

	evm.StateDB.SetCode(contractaddress, common.Hex2Bytes(HexTestCode))
	caller := vm.AccountRef(evm.Origin)
	ret, _, err := evm.Call(
		caller,
		contractaddress,
		TestInput,
		evm.GasLimit,
		new(big.Int))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ret)
	}
}
