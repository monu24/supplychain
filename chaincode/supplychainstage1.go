package main

import (
	"errors"
	"fmt"

	"encoding/json"



	"github.com/hyperledger/fabric/core/chaincode/shim"
)


type SimpleChaincode struct {
}





type MilkContainer struct{

        ContainerID string `json:"containerid"`
        User string        `json:"user"`
}

type SupplyCoin struct{

        CoinID string `json:"coinid"`
        User string        `json:"user"`
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
func(t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
                                                          
}

       err = stub.PutState("hello world",[]byte("welcome ti supply chain management"))  //Just to check the network 
       if err != nil {
		return nil, err
}
/*
      var empty []string
	jsonAsBytes, _ := json.Marshal(empty)                          //marshal an emtpy array of strings to clear the index
	err = stub.PutState(milkIndexStr, jsonAsBytes)                 //Making milk container list as empty - resetting
	if err != nil {
		return nil, err
} 

        err = stub.PutState(coinIndexStr, jsonAsBytes)                 //Making coin list as empty
        if err != nil {
                return nil, err
}
*/
        return nil, nil

}



func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {													//initialize the chaincode state, used as reset
		return t.Init(stub, "init", args)
        }else if function == "Create_milkcontainer" {		//creates a milk container-invoked by supplier   
		return t.Create_milkcontainer(stub, args)      

	}else if function == "Create_coinmarket" {		//creates a coin - invoked by market
		return t.Create_coinmarket(stub, args)	

        }else if function == "Create_coinlogistics"{              //creates a coin - invoked by logistics 
                return t.Create_coinmarket(stub, args)
       } 


     return nil,nil

}


func (t *SimpleChaincode) Create_milkcontainer(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
var err error

// "1x22" "supplier" 
// args[0] args[1] 

id := args[0]
user := args[1]

milkAsBytes, err := stub.GetState(id) 
if err != nil {
		return nil, errors.New("Failed to get details og given id") 
}

res := MilkContainer{} 
json.Unmarshal(milkAsBytes, &res)

if res.ContainerID == id{

        fmt.Println("Container already exixts")
        fmt.Println(res)
        return nil,errors.New("This cpontainer alreadt exists")
}

res.ContainerID = id
res.User = user

milkAsBytes, _ =json.Marshal(res)

stub.PutState(id,milkAsBytes)

return nil,nil

}


func (t *SimpleChaincode) Create_coinmarket(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

//"1x245" "Market"
id := args[0]
user:= args[1]

coinAsBytes , err := stub.GetState(id)
if err != nil{
              return nil, errors.New("Failed to get details of given id")
} 

res :=SupplyCoin{}

json.Unmarshal(coinAsBytes, &res)

if res.CoinID == id{

          fmt.Println("Coin already exists")
          fmt.Println(res)
          return nil,errors.New("This coin already exists")
}

res.CoinID = id
res.User = user

coinAsBytes, _ = json.Marshal(res)
stub.PutState(id,coinAsBytes)
return nil,nil
}
 

func(t *SimpleChaincode) Create_coinlogistics(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {


// "1x226" "Logistics"

id := args[0]
user:= args[1]

coinAsBytes , err := stub.GetState(id)
if err != nil{
              return nil, errors.New("Failed to get details of given id")
}

res :=SupplyCoin{}

json.Unmarshal(coinAsBytes, &res)

if res.CoinID == id{

          fmt.Println("Coin already exists")
          fmt.Println(res)
          return nil,errors.New("This coin already exists")
}

res.CoinID = id
res.User = user

coinAsBytes, _ = json.Marshal(res)
stub.PutState(id,coinAsBytes)
return nil,nil
}

func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface,function string, args []string) ([]byte, error) {

if function == "read" {						//read a variable
		return t.read(stub, args)
	}
	fmt.Println("query did not find func: " + function)						//error

	return nil, errors.New("Received unknown function query")
}


func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var name, jsonResp string
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the variable to query")
	}


	name = args[0]
	valAsbytes, err := stub.GetState(name)				//get the var from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for  \"}"
		return nil, errors.New(jsonResp)
	}

	



return valAsbytes, nil										       //send it onward
}









