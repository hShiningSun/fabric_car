package sdk_order

import (
	"errors"
	"fabric_car/project/sdk_helper"
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
)

func GetOrders() ([]fab.Orderer, error) {
	orderers := []fab.Orderer{}

	sdk := sdk_helper.Get1sdk()
	// context1 := sdk_helper.Get2Context(context1)

	ctx, err := sdk.Context()()
	if err != nil {
		return nil, errors.New("get ctx for GetOrderer error: " + err.Error())
	}

	endp := ctx.EndpointConfig()
	orderconfigs := endp.OrderersConfig()
	inp := ctx.InfraProvider()

	for _, config := range orderconfigs {

		orderer, err := inp.CreateOrdererFromConfig(&config)
		if err != nil {
			fmt.Printf("get order failed:", err.Error())
			return nil, err
		}

		orderers = append(orderers, orderer)
	}

	return orderers, nil

}

func GetOrderWithName(orderName string) (fab.Orderer, error) {
	var orderer fab.Orderer

	sdk := sdk_helper.Get1sdk()
	// context1 := sdk_helper.Get2Context(context1)

	ctx, err := sdk.Context()()
	if err != nil {
		return nil, errors.New("get ctx for GetOrderer error: " + err.Error())
	}

	endp := ctx.EndpointConfig()
	orderconfig, ok := endp.OrdererConfig(orderName)
	if !ok {
		fmt.Printf("get orderconfig failed:")
		return nil, errors.New("get orderconfig err for GetOrderWithName ")
	}
	inp := ctx.InfraProvider()

	orderer, err = inp.CreateOrdererFromConfig(orderconfig)

	if err != nil {
		fmt.Printf("get orderer failed:", err.Error())
		return nil, err
	}

	return orderer, nil

}
