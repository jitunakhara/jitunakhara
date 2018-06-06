package main

import (
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type SimpleChaincode struct {
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Example Init")
	
	return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Example Invoke")
	function, args := stub.GetFunctionAndParameters()
	if function == "addData" {
		Aval, err := strconv.Atoi(args[1])
		if err != nil {
			return shim.Error("Invalid transaction amount, expecting a integer value")
		}
		err = stub.PutState(args[0], []byte(strconv.Itoa(Aval)))
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success(nil)
	} else if function == "readData" {

		Avalbytes, err := stub.GetState(args[0])
		if err != nil {
			jsonResp := "{\"Error\":\"Failed to get state for " + args[0] + "\"}"
			return shim.Error(jsonResp)
		}
		jsonResp := "{\"Name\":\"" + args[0] + "\",\"Amount\":\"" + string(Avalbytes) + "\"}"
		fmt.Printf("Query Response:%s\n", jsonResp)
		return shim.Success([]byte(jsonResp))
	}

	return shim.Error("Invalid invoke function name. Expecting \"invoke\" \"delete\" \"query\"")
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
