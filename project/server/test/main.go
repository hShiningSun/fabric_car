package main

import (
	"encoding/json"
	"fabric_car/project/sdk_chaincode"
	"fabric_car/project/sdk_channel"
	"fabric_car/project/sdk_const"
	"fmt"
	"net/http"
	"strings"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type resultInfo struct {
	Code string `json:"code"` //Status 在json中用status替换
	Msg  string `json:"msg"`
	Data car    `json:"data"`
}

type car struct {
	CarID  int     `json:"carId"`  // 汽车id
	Name   string  `json:"name"`   // 汽车名字
	Color  string  `json:"color"`  // 颜色
	Amount float64 `json:"amount"` // 汽车金额
	// IDCard       string  `json:"idCard"`       // 身份证
}

var routes = gin.Default()

func main() {

	// routes.POST("car/invoke", testInvokeChincode)
	routes.GET("channel/create", createChannel)       //routes.GET("channel/create:channelid/:orguser/:orgname", createChannel)
	routes.GET("channel/join", joinChannel)           //routes.POST("channel/join/:orguser/:orgname/:peer", joinChannel)
	routes.GET("chaincode/install", installChaincode) //routes.POST("chaincode/install/:orguser/:orgname/:peer/:chaincodePath", installChaincode)
	routes.GET("chaincode/instantiated", instantiatedChaincode)
	routes.GET("chaincode/upgrade", upgradeChaincode)

	// 调用
	routes.POST("car/create", createCar)
	routes.POST("car/query", queryCar)

	routes.Run(":8888")
	// num := 3

	// switch num {
	// case 0:
	// 	testCreateChannel() // 创建频道
	// case 1:
	// 	testJoinChannel() // 加入频道
	// case 2:
	// 	testInstanllChaincode() //安装链码
	// case 3:
	// 	testIntantiateChaincode() //实例化链码
	// case 4:
	// 	testInvokeChincode() //调用链码
	// case 5:
	// 	testQueryChainCode() //查询链码
	// case 6:
	// 	testUpdragChainCode() //升级链码
	// default:

	// }
}

func createChannel(c *gin.Context) {
	channelid := c.DefaultQuery("channelid", sdk_const.ChannelName) //可设置默认值
	orguser := c.Query("orguser")
	orgname := c.Query("orgname")
	fmt.Println(channelid)
	fmt.Println(orguser)
	fmt.Println(orgname)
	err := sdk_channel.CreateChannel(channelid, orguser, orgname)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "创建channel成功",
		})
	}
}
func joinChannel(c *gin.Context) {
	orguser := c.Query("orguser")
	orgname := c.Query("orgname")
	peerURL := c.Query("peer")
	err := sdk_channel.JoinChannel(peerURL, orguser, orgname)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "加入channel成功",
		})
	}

}
func installChaincode(c *gin.Context) {

	request := sdk_chaincode.InstallChainCodeRequest{
		OrgUser:          c.Query("orguser"),
		OrgName:          c.Query("orgname"),
		PeerURL:          c.Query("peer"),
		ChaincodeName:    c.Query("chaincodeName"),
		ChainCodePath:    c.Query("chaincodePath"),
		ChaincodeVersion: c.Query("chaincodeVersion"),
	}

	err := sdk_chaincode.InstallChaincode(request)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "installChincode成功",
		})
	}

}
func instantiatedChaincode(c *gin.Context) {
	//
	request := sdk_chaincode.InstantiateChaincodeRequest{
		OrgUser:                  c.Query("orguser"),
		OrgName:                  c.Query("orgname"),
		PeerURL:                  c.Query("peer"),
		ChainCodePolicy:          c.Query("chaincodePolicy"),
		ChainCodeInstantiateArgs: c.Query("chaincodeInstantiateArgs"),
		ChaincodeName:            c.Query("chaincodeName"),
		ChainCodePath:            c.Query("chaincodePath"),
		ChaincodeVersion:         c.Query("chaincodeVersion"),
	}

	err := sdk_chaincode.InstantiateChaincode(request)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "instantiatedChincode成功",
		})
	}
}
func upgradeChaincode(c *gin.Context) {
	installedRequest := sdk_chaincode.InstallChainCodeRequest{
		OrgUser:          c.Query("orguser"),
		OrgName:          c.Query("orgname"),
		PeerURL:          c.Query("peer"),
		ChaincodeName:    c.Query("chaincodeName"),
		ChainCodePath:    c.Query("chaincodePath"),
		ChaincodeVersion: c.Query("chaincodeVersion"),
	}

	err := sdk_chaincode.InstallChaincode(installedRequest)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "installChincode成功",
		})
	}

	upgradeRequest := sdk_chaincode.UpgradeChancodeRequest{
		OrgUser:              c.Query("orguser"),
		OrgName:              c.Query("orgname"),
		PeerURL:              c.Query("peer"),
		ChaincodeName:        c.Query("chaincodeName"),
		ChainCodePath:        c.Query("chaincodePath"),
		ChaincodeVersion:     c.Query("chaincodeVersion"),
		ChainCodePolicy:      c.Query("chaincodePolicy"),
		ChainCodeUpdrageArgs: c.Query("chaincodeUpdrageArgs"),
		ChannelName:          c.Query("channelName"),
	}

	err = sdk_chaincode.UpgradeChancode(upgradeRequest)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "升级成功",
		})
	}

}
func createCar(c *gin.Context) {
	jsonSt := getJSONString(c)
	request := &channel.Request{
		ChaincodeID: "cc_car",
		Fcn:         "create",
		Args:        [][]byte{[]byte(jsonSt)},
	}

	invokeRequest := sdk_chaincode.InvokeChainCodeRequest{
		OrgUser:     c.Query("orguser"),
		OrgName:     c.Query("orgname"),
		PeerURL:     c.Query("peer"),
		ChannelName: c.Query("channelName"),
	}

	response := sdk_chaincode.InvokeChainCode(*request, invokeRequest)

	responSt := string(response.Payload)

	var retursInfo resultInfo
	err := json.Unmarshal([]byte(responSt), &retursInfo)

	if err == nil && retursInfo.Data.CarID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    retursInfo.Code,
			"message": retursInfo.Msg,
		})
	} else if err == nil && retursInfo.Data.CarID != 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    retursInfo.Code,
			"message": retursInfo.Msg,
			"data":    retursInfo.Data,
		})
	} else if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    retursInfo.Code,
			"message": retursInfo.Msg,
		})
	}
}

func queryCar(c *gin.Context) {

	//写查询了

	// jsonSt := getJSONString(c)

	args := [][]byte{}
	argSts := strings.Split(c.Query("chaincodeQueryArgs"), ",")

	for _, st := range argSts {
		args = append(args, []byte(st))
	}

	request := &channel.Request{
		ChaincodeID: "cc_car",
		Fcn:         "query",
		Args:        args,
	}

	invokeRequest := sdk_chaincode.InvokeChainCodeRequest{
		OrgUser:     c.Query("orguser"),
		OrgName:     c.Query("orgname"),
		PeerURL:     c.Query("peer"),
		ChannelName: c.Query("channelName"),
	}

	response := sdk_chaincode.InvokeChainCode(*request, invokeRequest)

	responSt := string(response.Payload)

	var retursInfo resultInfo
	err := json.Unmarshal([]byte(responSt), &retursInfo)

	if err == nil && retursInfo.Data.CarID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    retursInfo.Code,
			"message": retursInfo.Msg,
		})
	} else if err == nil && retursInfo.Data.CarID != 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    retursInfo.Code,
			"message": retursInfo.Msg,
			"data":    retursInfo.Data,
		})
	} else if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    retursInfo.Code,
			"message": retursInfo.Msg,
		})
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
	// sdk_chaincode.InstallChaincode()
}

func testIntantiateChaincode() {
	// sdk_chaincode.InstantiateChaincode()
}

// getJSONString 得到服务器传来转化后的json
func getJSONString(c *gin.Context) string { //[][]byte {

	var obj car
	err := c.BindJSON(&obj)
	if err == nil {

	}
	// 转化好了对象 再转化成json字符串
	jsonst, errjson := json.Marshal(obj)
	if errjson != nil {
		logrus.Debug("转化有误，检查getJsonSt")
	}

	fmt.Print(string(jsonst))
	//return [][]byte{[]byte(string(jsonst))}
	return string(jsonst)
}

func testInvokeChincode() {
	sdk_chaincode.InvokeChainCodeTest()
}

func testQueryChainCode() {
	sdk_chaincode.QueryChaincode()
}

func testUpdragChainCode() {
	// 先安装新的链码 升级才能找到新的包
	// sdk_chaincode.InstallChaincode()
	// sdk_chaincode.UpgradeChancode()
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
