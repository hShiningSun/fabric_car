package sdk_channel

import (
	"errors"
	"fabric_car/project/sdk_const"
	"fabric_car/project/sdk_peer"
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	// "github.com/hyperledger/fabric/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"

	"fabric_car/project/sdk_helper"

	mspclient "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
)

// CreateChannel 创建channel
// orgUser 创建者的组织用户 一般默认Admin
// orgName 创建者的组织名字 我这里例子是chenman
// 组织的名字 同时也是申请属于这一个组织的sdk 用途
func CreateChannel(channelID string, orgUser string, orgName string) error {

	// 第一步 根据配置文件 获取sdk，1.获取配置，2.根据配置生成sdk
	sdk := sdk_helper.Get1sdk()
	context1 := sdk_helper.Get2Context(sdk, orgUser, orgName)
	resmgmtClient := sdk_helper.Get3resmgmtClint(context1)

	// 有了客户端就创建channel

	// 创建 需要 本地 签名授权，网络才知道 是谁在创建，有没有这个权限
	// 先获会员系统 客户端
	mspClient, err := mspclient.New(sdk.Context(), mspclient.WithOrg(orgName))
	if err != nil {
		fmt.Println("mspClient error =" + err.Error())
	}
	// 会员客户端 通过 会员的名字，参数是id  来返回一个该会员的签名
	adminIdentity, err := mspClient.GetSigningIdentity(orgUser)
	if err != nil {
		fmt.Println("adminIdentity error =" + err.Error())
	}

	// 正式创建channel  传入 频道名字 ，频道配置文件的位置，还有创建者 的签名
	req := resmgmt.SaveChannelRequest{ChannelID: channelID,
		ChannelConfigPath: sdk_const.ChannelConfigPath,
		SigningIdentities: []msp.SigningIdentity{adminIdentity}}

	// 之前这里一直报错排序节点连接不上，后来发现配置文件出错，配置一下映射，再orderer 再设置一下映射就好了
	// 有这个频道了，再创建就会报错
	resp, err := resmgmtClient.SaveChannel(req, resmgmt.WithOrdererEndpoint(sdk_const.OrdererEndpoint), resmgmt.WithRetry(retry.DefaultResMgmtOpts))

	if err != nil {

		fmt.Println("SaveChannel error =" + err.Error())
		return err
	}
	if resp.TransactionID == "" {
		fmt.Println("创建channel 失败")
		return errors.New("TransactionID = null")
	}

	return nil
}

// JoinChannel 加入频道 要讲组织所有peer加入的话 url=""就可以
func JoinChannel(peerURL string, orgUser string, orgName string) error {

	sdk := sdk_helper.Get1sdk()
	peerContext := sdk_helper.Get2Context(sdk, orgUser, orgName)
	rmsclient := sdk_helper.Get3resmgmtClint(peerContext)

	// if err != nil {
	// 	fmt.Printf("rmsclient create err = " + err.Error())
	// }

	// 要加入频道的节点
	ctx, err := sdk.Context()()
	if err != nil {
		fmt.Printf("context creation failed: %s", err)
	}

	endc := ctx.EndpointConfig()

	orderc, ok := endc.OrdererConfig(sdk_const.OrderURL)
	if !ok {
		fmt.Println("get orderconfig error")
	}

	order, err := ctx.InfraProvider().CreateOrdererFromConfig(orderc)
	if err != nil {
		fmt.Println("get order err = " + err.Error())
	}

	// 加入频道
	// 不填写节点，默认将所有对等节点加入
	if peerURL == "" {
		err = rmsclient.JoinChannel(sdk_const.ChannelName, resmgmt.WithOrderer(order))
	} else {
		peer, _ := sdk_peer.GetPeerWithNameOrURL(peerURL)
		err = rmsclient.JoinChannel(sdk_const.ChannelName, resmgmt.WithOrderer(order), resmgmt.WithTargets(peer))
	}

	if err != nil {
		fmt.Printf("join channel err = " + err.Error())
		return err
	} else {
		fmt.Printf("Join Channel Successful")
	}

	return nil
}
