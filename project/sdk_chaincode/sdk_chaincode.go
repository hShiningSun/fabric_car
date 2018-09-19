package sdk_chaincode

import (
	"fabric_car/project/sdk_const"
	"fabric_car/project/sdk_helper"
	"fabric_car/project/sdk_peer"
	"fmt"
	"os"
	"strings"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	mspclient "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/fab/ccpackager/gopackager"

	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/context"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"

	// "github.com/hyperledger/fabric-sdk-go/test/integration"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/common/cauthdsl"
	// "github.com/hyperledger/fabric-sdk-go/test/integration"
)

func InstallChaincode() {
	// 第一步获取sdk
	sdk := sdk_helper.Get1sdk()

	// 第二步获取上下文
	context := sdk_helper.Get2Context(sdk, sdk_const.UserName, sdk_const.OrgName)

	// 第三步获取需要的客户端
	remsClient := sdk_helper.Get3resmgmtClint(context)

	// 获取需要安装链码的peer
	peerchenman, _ := sdk_peer.GetPeerWithNameOrURL(sdk_const.PeerName)

	// 组装安装请求
	gopath := os.Getenv("GOPATH")
	ccpkg, err := gopackager.NewCCPackage(sdk_const.ChainCodePath, gopath)
	if err != nil {
		fmt.Println("get ccpkg err = " + err.Error())
	}
	ccreq := resmgmt.InstallCCRequest{
		Name:    sdk_const.ChainCodeName,
		Path:    sdk_const.ChainCodePath, //不能填写背书节点的路径，填写本地sdk能读到的路径
		Version: sdk_const.ChainCodeVersion,
		Package: ccpkg,
	}

	// 开始安装 ,填写一个请求，和节点  不填写节点 就默认所有
	ccreponse, err := remsClient.InstallCC(ccreq, resmgmt.WithTargets(peerchenman))

	if err != nil {
		fmt.Println("cc install err = " + err.Error())
	}

	if len(ccreponse) <= 0 {
		fmt.Println("cc install err")
	} else {
		fmt.Println("cc install success")
	}

}

// 实例化chaincode
func InstantiateChaincode() {
	sdk := sdk_helper.Get1sdk()
	context1 := sdk_helper.Get2Context(sdk, sdk_const.UserName, sdk_const.OrgName)
	resmClient := sdk_helper.Get3resmgmtClint(context1)

	mspClient, err := mspclient.New(sdk.Context(), mspclient.WithOrg(sdk_const.OrgName))
	user, err := mspClient.GetSigningIdentity(sdk_const.UserName)

	key := user.Identifier().MSPID + "_" + user.Identifier().ID
	sessions := make(map[string]context.ClientProvider)
	var session context.ClientProvider
	session = sessions[key]
	if session == nil {
		session = sdk.Context(fabsdk.WithIdentity(user))
		fmt.Println("[%s]------[%s]", user.Identifier().ID, user.Identifier().MSPID)
		// cliconfig.Config().Logger().Debugf("Created session for user [%s] in org [%s]", user.Identifier().ID, user.Identifier().MSPID)
		sessions[key] = session
	}

	resmClient, _ = resmgmt.New(session)

	// 实例化链码策略

	// 这个报错ccPolicy := cauthdsl.SignedByAnyMember(members)
	ccPolicy, err := cauthdsl.FromString(sdk_const.ChainCodePolicy)
	if err != nil {
		fmt.Println("get ccpolicy  err = " + err.Error())
	}

	Args1 := [][]byte{}
	argSts := strings.Split(sdk_const.ChainCodeInstantiateArgs, ",")

	for _, st := range argSts {
		Args1 = append(Args1, []byte(st))
	}

	reqcc := resmgmt.InstantiateCCRequest{
		Name:    sdk_const.ChainCodeName,
		Path:    sdk_const.ChainCodePath, //不能填写背书节点的路径，填写本地sdk能读到的路径
		Version: sdk_const.ChainCodeVersion,
		Policy:  ccPolicy,
		Args:    Args1, //[][]byte{[]byte("init"), []byte("chenman"), []byte("10"), []byte("lixingxing"), []byte("10")},
	}

	// 给chenman 实例化
	peerchenman, _ := sdk_peer.GetPeerWithNameOrURL(sdk_const.PeerName)
	fmt.Println("peer is url = " + peerchenman.URL())

	response, err := resmClient.InstantiateCC(sdk_const.ChannelName, reqcc, resmgmt.WithTargets(peerchenman))
	if err != nil {
		fmt.Println("instantiate cc  err = " + err.Error())
		return
	}

	if response.TransactionID == "" {
		fmt.Println("instantiate cc err = " + err.Error())
	} else {
		fmt.Println("instantiate cc Successful")
	}

}

//查询
func QueryChaincode() {
	sdk := sdk_helper.Get1sdk()
	context1 := sdk_helper.Get2Context(sdk, sdk_const.UserName, sdk_const.OrgName)

	client := sdk_helper.Get3channelClient(context1, sdk_const.ChannelName)

	Args1 := [][]byte{}
	argSts := strings.Split(sdk_const.ChainCodeQueryArgs, ",")

	for _, st := range argSts {
		Args1 = append(Args1, []byte(st))
	}

	req := channel.Request{
		ChaincodeID: sdk_const.ChainCodeName,
		Fcn:         "query",
		Args:        Args1, //[][]byte{[]byte("chenman")},
	}

	peerchenman, _ := sdk_peer.GetPeerWithNameOrURL(sdk_const.PeerName)
	response, err := client.Query(req, channel.WithTargets(peerchenman))

	if err != nil {
		fmt.Println("request chaincode query err = " + err.Error())

		return
	}

	if len(response.Payload) == 0 {
		fmt.Println("request chaincode query err = " + err.Error())

		return
	} else {
		fmt.Println("query success " + string(response.Payload))
	}

}

// InvokeChainCode 正式调用的方法
func InvokeChainCode(request channel.Request) channel.Response {
	sdk := sdk_helper.Get1sdk()
	context := sdk_helper.Get2Context(sdk, sdk_const.UserName, sdk_const.OrgName)
	channelClient := sdk_helper.Get3channelClient(context, sdk_const.ChannelName)

	peer1, err := sdk_peer.GetPeerWithNameOrURL(sdk_const.PeerName)
	resp, err := channelClient.Execute(request, channel.WithTargets(peer1))
	if err != nil {
		fmt.Println("invoke chaincode err = " + err.Error())
	}
	return resp
}

// iterations 并发的数量
func InvokeChainCodeTest() {
	sdk := sdk_helper.Get1sdk()
	context := sdk_helper.Get2Context(sdk, sdk_const.UserName, sdk_const.OrgName)
	channelClient := sdk_helper.Get3channelClient(context, sdk_const.ChannelName)

	Args1 := [][]byte{}
	argSts := strings.Split(sdk_const.ChainCodeInvokeArgs, ",")

	for _, st := range argSts {
		Args1 = append(Args1, []byte(st))
	}

	peer1, err := sdk_peer.GetPeerWithNameOrURL(sdk_const.PeerName)
	resp, err := channelClient.Execute(channel.Request{
		ChaincodeID: sdk_const.ChainCodeName,
		Fcn:         "invoke",
		Args:        Args1, //[][]byte{[]byte("lixingxing"), []byte("chenman"), []byte("2")},
	}, channel.WithTargets(peer1))

	if err != nil {
		fmt.Println("invoke chaincode err = " + err.Error())
	}

	if resp.ChaincodeStatus != 200 {
		fmt.Println("invoke chaincode err")
	} else {
		fmt.Println("~~hyc~~invoke chaincode success ")
	}

}

func UpgradeChancode() {

	sdk := sdk_helper.Get1sdk()
	context1 := sdk_helper.Get2Context(sdk, sdk_const.UserName, sdk_const.OrgName)
	client1 := sdk_helper.Get3resmgmtClint(context1)

	ccPolicy, err := cauthdsl.FromString(sdk_const.ChainCodePolicy)
	if err != nil {
		fmt.Println("get ccpolicy  err = " + err.Error())
	}

	//Args:    [][]byte{[]byte("init"), []byte("chenman"), []byte("10"), []byte("lixingxing"), []byte("10")},
	Args1 := [][]byte{}
	argSts := strings.Split(sdk_const.ChainCodeUpdrageArgs, ",")

	for _, st := range argSts {
		Args1 = append(Args1, []byte(st))
	}

	reqcc := resmgmt.UpgradeCCRequest{
		Name:    sdk_const.ChainCodeName,
		Path:    sdk_const.ChainCodePath, //不能填写背书节点的路径，填写本地sdk能读到的路径
		Version: sdk_const.ChainCodeVersion,
		Args:    Args1,
		Policy:  ccPolicy,
	}

	peer1, err := sdk_peer.GetPeerWithNameOrURL(sdk_const.PeerName)
	resp, err := client1.UpgradeCC(sdk_const.ChannelName, reqcc, resmgmt.WithTargets(peer1))

	if err != nil {
		fmt.Println("chaincode upgrade err = " + err.Error())

		return
	}

	if resp.TransactionID == "" {
		fmt.Println("chaincode upgrade err = " + err.Error())

		return
	} else {
		fmt.Println("chaincode upgrade successful ")
	}

}
