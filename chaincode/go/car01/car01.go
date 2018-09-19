package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"fabric_car/chaincode/go/action"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type CarChincode struct {
}

type car struct {
	CarID  int     `json:"carId"`  // 汽车id
	Name   int     `json:"name"`   // 汽车名字
	Color  int     `json:"color"`  // 颜色
	Amount float64 `json:"amount"` // 汽车金额
	// IDCard       string  `json:"idCard"`       // 身份证
}
type resultInfo struct {
	Code string `json:"code"` //Status 在json中用status替换
	Msg  string `json:"msg"`
	Data car    `json:"data"`
}

func response(isSuccess bool, msg string, data *car) peer.Response {
	var res resultInfo //res := resultInfo{}
	if isSuccess == true {
		res.Code = "0"
	} else {
		res.Code = "1"
	}
	res.Msg = msg

	if data != nil {
		res.Data = *data
	}

	// 把对象转化为json，这里可以得出，sturct 相当于oc的类，
	resJSON, err := json.Marshal(res)

	if err != nil {
		return shim.Error("response -> json error :" + err.Error())
	}

	return shim.Success(resJSON)
}

// Init ...chaincode instantiate and upgrade call
func (c *CarChincode) Init(stub shim.ChaincodeStubInterface) peer.Response {

	fun, args := stub.GetFunctionAndParameters()
	fmt.Println(fun + " is sdk call function name ,there not have use ")

	if len(args) == 0 {
		return response(true, "upgrade successful", nil)
	}

	if len(args) != 4 {
		return response(false, "upgrade error,args num should be 4", nil)
	}

	A := args[0]
	Avalue, err := strconv.Atoi(args[1])
	if err != nil {
		return response(false, "convert args num 1 to int error :"+err.Error(), nil)
	}

	B := args[2]
	Bvalue, err := strconv.Atoi(args[3])
	if err != nil {
		return response(false, "convert args num 3 to int error :"+err.Error(), nil)
	}

	err = stub.PutState(A, []byte(strconv.Itoa(Avalue)))
	if err != nil {
		response(false, err.Error(), nil)
	}

	err = stub.PutState(B, []byte(strconv.Itoa(Bvalue)))
	if err != nil {
		response(false, err.Error(), nil)
	}

	return response(true, "instantiate successful", nil)
}

// Invoke ... call chaincode
func (c *CarChincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fun, args := stub.GetFunctionAndParameters()
	if len(fun) == 0 {
		return response(false, "no have invoke args", nil)
	}

	switch fun {
	case action.Transfer:
		return c.transfer(args, stub)
	case action.Create:
		return c.create(args, stub)
	case action.Update:
		return c.update(args, stub)
	case action.Delete:
		return c.delete(args, stub)
	default:
		return response(false, "传入的函数名称没有找到", nil)
	}

}

func (c *CarChincode) transfer(args []string, stub shim.ChaincodeStubInterface) peer.Response {
	if len(args) != 3 {
		return response(false, "transfer args error not have num 3", nil)
	}

	current := args[0]
	transferValue, err := strconv.Atoi(args[1])
	if err != nil {
		return response(false, "transfer num error :"+err.Error(), nil)
	}
	target := args[2]

	currentValue, err := stub.GetState(current)
	if err != nil {
		return response(false, "transfer get currentValue error :"+err.Error(), nil)
	}
	if currentValue == nil {
		return response(false, "transfer get currentValue nil ", nil)
	}

	targetValue, err := stub.GetState(target)
	if err != nil {
		return response(false, "transfer get targetValue error :"+err.Error(), nil)
	}
	if targetValue == nil {
		return response(false, "transfer get currentValue nil ", nil)
	}

	Aval, err := strconv.Atoi(string(currentValue))
	if err != nil {
		return response(false, err.Error(), nil)
	}
	Bval, err := strconv.Atoi(string(targetValue))
	if err != nil {
		return response(false, err.Error(), nil)
	}

	Aval = Aval - transferValue
	Bval = Bval + transferValue

	err = stub.PutState(current, []byte(strconv.Itoa(Aval)))
	if err != nil {
		return response(false, "put currentVal error :"+err.Error(), nil)
	}

	err = stub.PutState(target, []byte(strconv.Itoa(Bval)))
	if err != nil {
		return response(false, "put targetVal error :"+err.Error(), nil)
	}

	return response(true, "transfer successful", nil)
}
func (c *CarChincode) create(args []string, stub shim.ChaincodeStubInterface) peer.Response {
	if len(args) != 1 {
		return response(true, "create car args error", nil)
	}
	jsonSt := args[0]
	var newCar car
	err := json.Unmarshal([]byte(jsonSt), &newCar)
	if err != nil {
		return response(true, "create car json error:"+err.Error(), nil)
	}

	oldCar, err := stub.GetState(strconv.Itoa(newCar.CarID))
	if err != nil {
		return response(false, "get oldCar with createCar is carId,error="+err.Error(), nil)
	}
	if oldCar != nil {
		return response(false, "createCar  carId is haved", nil)
	}
	carByte, err := json.Marshal(newCar)
	if err != nil {
		return response(false, err.Error(), nil)
	}
	err = stub.PutState(strconv.Itoa(newCar.CarID), carByte)
	if err != nil {
		return response(false, "put carjson error ="+err.Error(), nil)
	}

	return response(true, "create car successful", &newCar)
}
func (c *CarChincode) update(args []string, stub shim.ChaincodeStubInterface) peer.Response {

	return response(true, "", nil)
}
func (c *CarChincode) delete(args []string, stub shim.ChaincodeStubInterface) peer.Response {

	return response(true, "", nil)
}

func main() {
	err := shim.Start(new(CarChincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
