package main

import (
	"fabric_car/project/sdk_chaincode"
	"fabric_car/project/sdk_channel"
	"fmt"
	// "fabric_car/project/sdk_const"
	// "fabric_car/project/sdk_helper"
	// "fabric_car/project/sdk_order"
	// "fabric_car/project/sdk_peer"
	// "fmt"
	// "strings"
	// "github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	// "github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	// // "github.com/hyperledger/fabric-sdk-go/test/integration"
	// "github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/common/cauthdsl"
)

type kk struct {
	st1 string
	st2 string
}

func main() {

	num := 6

	switch num {
	case 0:
		testCreateChannel() // 创建频道
	case 1:
		testJoinChannel() // 加入频道
	case 2:
		testInstanllChaincode() //安装链码
	case 3:
		testIntantiateChaincode() //实例化链码
	case 4:
		testInvokeChincode() //调用链码
	case 5:
		testQueryChainCode() //查询链码
	case 6:
		testUpdragChainCode() //升级链码
	default:

	}
}
func testCreateChannel() {
	err := sdk_channel.CreateChannel("mychannel", "Admin", "chenman")
	if err != nil {
		fmt.Println("create channel failed err = " + err.Error())
	} else {
		fmt.Println("~~create channel successful")
	}
}

func testJoinChannel() {
	err := sdk_channel.JoinChannel("192.168.56.115", "Admin", "chenman")
	if err != nil {
		fmt.Println("join channel failed err = " + err.Error())
	} else {
		fmt.Println("~~join channel successful")
	}
}

func testInstanllChaincode() {
	sdk_chaincode.InstallChaincode()
}

func testIntantiateChaincode() {
	sdk_chaincode.InstantiateChaincode()
}

func testInvokeChincode() {
	sdk_chaincode.InvokeChainCode()
}

func testQueryChainCode() {
	sdk_chaincode.QueryChaincode()
}

func testUpdragChainCode() {
	// 先安装新的链码 升级才能找到新的包
	sdk_chaincode.InstallChaincode()
	sdk_chaincode.UpgradeChancode()
}

// 首先 根据配置文件 获取sdk
// 要做什么操作，先实例化一个 对应的 客户端 来进行操作
// 实例化客户端 需要相应的 上下文
// 所以第二步骤就是 实例化 你需要操作 的客户端 所需要的上下文喽
// 用客户端 进行 相关操作就可以了

// 下面就按照我 搭建 环境测试的例子步骤 来 使用sdk
// 创建chanell channel_sdk_example
// peer 加入此 频道
// 安装 chaincode
// 实例化 chaincode
// 查询 chaincode

// func createChannel(channelName string) bool {

// 	// 第一步 根据配置文件 获取sdk，1.获取配置，2.根据配置生成sdk
// 	sdkconfig := config.FromFile(sdk_const.ConfigFilePath)
// 	sdk, err := fabsdk.New(sdkconfig)

// 	// 检查 sdk 初始化成功没有
// 	if err != nil {
// 		fmt.Printf("get sdk have error =" + err.Error())
// 	}

// 	// 构建 创建频道需要的上下文,fabsdk 就是用来创建上下文的
// 	// 回一下 创建频道 必要条件
// 	// 1. 根据peer.yaml启动peer 容器，peer 跟orderer 连接起来
// 	// 2. 进入容器，根据channel.tx
// 	// 3. 创建频道
// 	// 看fabric 的上下文api有几个分别是，withOrg withUser WithIdentfiler
// 	// 相关的组织 ，相关的用户 相关的ident
// 	// 相关的组织，我这个节点先暂时定义 chenman ,
// 	createChannelContext := sdk.Context(fabsdk.WithOrg(sdk_const.OrgName), fabsdk.WithUser(sdk_const.UserName))

// 	// 创建channel 需要resm 客户端
// 	resmgmtClient, err := resmgmt.New(createChannelContext)
// 	if err != nil {
// 		fmt.Printf("resmgmtClient get error =" + err.Error())
// 	}

// 	// 有了客户端就创建channel

// 	mspClient, err := mspclient.New(sdk.Context(), mspclient.WithOrg(sdk_const.OrgName))
// 	if err != nil {
// 		fmt.Println("mspClient error =" + err.Error())
// 	}
// 	adminIdentity, err := mspClient.GetSigningIdentity(sdk_const.UserName)
// 	if err != nil {
// 		fmt.Println("adminIdentity error =" + err.Error())
// 	}

// 	req := resmgmt.SaveChannelRequest{ChannelID: sdk_const.ChannelName,
// 		ChannelConfigPath: sdk_const.ChannelConfigPath,
// 		SigningIdentities: []msp.SigningIdentity{adminIdentity}}

// 	// 之前这里一直报错排序节点连接不上，后来发现配置文件出错，配置一下映射，再orderer 再设置一下映射就好了
// 	// 有这个频道了，再创建就会报错
// 	resp, err := resmgmtClient.SaveChannel(req, resmgmt.WithOrdererEndpoint("hyc0.car.com"), resmgmt.WithRetry(retry.DefaultResMgmtOpts))
// 	is := true
// 	if err != nil {
// 		is = false
// 		fmt.Println("SaveChannel error =" + err.Error())
// 	}
// 	if resp.TransactionID == "" {
// 		fmt.Println("创建channel 失败")
// 	}

// 	return is
// }
