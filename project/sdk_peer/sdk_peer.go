package sdk_peer

import (
	"errors"
	"fabric_car/project/sdk_helper"
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	// "github.com/hyperledger/fabric-sdk-go/test/integration"
)

// GetAllPeerInAllNetwork 获取全部网络全部节点
func GetAllPeerInAllNetwork() ([]fab.Peer, error) {
	allPerr := []fab.Peer{}
	// 第一步获取sdk
	sdk := sdk_helper.Get1sdk()

	// 要加入频道的节点
	ctx, err := sdk.Context()()
	if err != nil {
		fmt.Printf("context creation failed: %s", err)
		return allPerr, err
	}

	endp := ctx.EndpointConfig()
	networkPeers := endp.NetworkPeers()
	inP := ctx.InfraProvider()

	for _, networkPeer := range networkPeers {

		peer1, err := inP.CreatePeerFromConfig(&networkPeer)
		if err != nil {
			fmt.Println("~~~get peer for networkPeer error = " + err.Error())
			return allPerr, err
		}
		allPerr = append(allPerr, peer1)

	}

	return allPerr, nil
}

// GetPeerWithNameOrURL 通过url来获取peer
func GetPeerWithNameOrURL(NameOrURL string) (fab.Peer, error) {

	// 第一步获取sdk
	sdk := sdk_helper.Get1sdk()

	// 要加入频道的节点
	ctx, err := sdk.Context()()
	if err != nil {
		fmt.Printf("context creation failed: %s", err)
		return nil, err
	}

	endp := ctx.EndpointConfig()

	peerconfig, ok := endp.PeerConfig(NameOrURL)
	if !ok {
		return nil, errors.New("get peer error")
	}
	peer, err := ctx.InfraProvider().CreatePeerFromConfig(&fab.NetworkPeer{
		PeerConfig: *peerconfig,
	})

	if err != nil {
		return nil, errors.New("get peer error")
	}

	if peer == nil {
		fmt.Println("~~do not have peer url :" + NameOrURL)
		return nil, errors.New("not have peer url" + NameOrURL)
	}

	return peer, nil

}

// GetPeersWithChannel 通过频道名字获取所有节点
func GetPeersWithChannel(channelName string) ([]fab.Peer, error) {

	channelPeers := []fab.Peer{}

	sdk := sdk_helper.Get1sdk()

	// 要加入频道的节点
	ctx, err := sdk.Context()()
	if err != nil {
		fmt.Printf("ctx create failed: %s", err)
		return channelPeers, err
	}

	endp := ctx.EndpointConfig()
	inp := ctx.InfraProvider()
	fabchannelPeers, ok := endp.ChannelPeers(channelName)

	if !ok {
		fmt.Println("get fab.ChannelPeers err")
		return channelPeers, errors.New("get fab.ChannelPeers err")
	}

	for _, value := range fabchannelPeers {
		networkpeer := value.NetworkPeer
		peer, err := inp.CreatePeerFromConfig(&networkpeer)
		if err != nil {
			return nil, errors.New("get peer for fabchannelPeers error:" + err.Error())
		}
		channelPeers = append(channelPeers, peer)
	}

	if len(channelPeers) == 0 {
		return nil, errors.New("get peer for fabchannelPeers nums = 0")
	}

	return channelPeers, nil

}

// GetPeersWithMSPID 获取一个MSPID 下所有的节点
func GetPeersWithMSPID(mspid string) ([]fab.Peer, error) {

	peersMSP := []fab.Peer{}

	sdk := sdk_helper.Get1sdk()

	ctx, err := sdk.Context()()
	if err != nil {
		return nil, errors.New("get ctx for GetPeersWithMSPID error: " + err.Error())
	}

	localDiscoveryProvider := ctx.LocalDiscoveryProvider()
	discoveryService, err := localDiscoveryProvider.CreateLocalDiscoveryService(mspid)
	if err != nil {
		return nil, errors.New("get discoveryService for GetPeersWithMSPID error: " + err.Error())
	}

	peersMSP, err = discoveryService.GetPeers()
	if err != nil {
		return nil, errors.New("get peersMSP for GetPeersWithMSPID error: " + err.Error())
	}

	return peersMSP, nil

}
