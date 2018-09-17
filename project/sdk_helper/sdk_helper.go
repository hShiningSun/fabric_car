package sdk_helper

import (
	"fabric_car/project/sdk_const"
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/context"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	// "github.com/hyperledger/fabric-sdk-go/test/integration"
	contextImpl "github.com/hyperledger/fabric-sdk-go/pkg/context"
)

const sdkConfigFile = "../../configfile/configfile.yaml"

// Get1sdk 第一步根据配置 获取sdk
func Get1sdk() *fabsdk.FabricSDK {

	configfile := config.FromFile(sdk_const.ConfigFilePath)

	sdk, err := fabsdk.New(configfile)

	// 检查 sdk 初始化成功没有
	if err != nil {
		fmt.Printf("get sdk have error =" + err.Error())
	}

	return sdk
}

// 第二步骤 根据信息 获取 上下文
func Get2Context(sdk *fabsdk.FabricSDK, username string, orgname string) context.ClientProvider {
	context := sdk.Context(fabsdk.WithUser(username), fabsdk.WithOrg(orgname))
	return context
}

// 第三步 根据上下文 获取先关操作的客户点
// 这里是resm,要获取相应的其他的client 就写get3qita
func Get3resmgmtClint(context context.ClientProvider) *resmgmt.Client {
	remsCliient, err := resmgmt.New(context)

	if err != nil {
		fmt.Println("get remsClient err = " + err.Error())
	}

	return remsCliient
}

// 第三步骤获取 channel 客户端
func Get3channelClient(context1 context.ClientProvider, channel_id string) *channel.Client {

	// oc 和 swift 的闭包写法
	channelProvider := func() (context.Channel, error) {
		return contextImpl.NewChannel(context1, channel_id)
	}

	channelCL, err := channel.New(channelProvider)

	if err != nil {
		fmt.Println("get channelClient err = " + err.Error())
	}

	return channelCL
}
