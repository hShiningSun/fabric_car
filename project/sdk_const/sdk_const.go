package sdk_const

const (
	// ConfigFilePath 初始化sdk的时候要读取配置文件
	ConfigFilePath = "/root/hycGo/src/fabric_car/project/configfile/configfile.yaml"
	// Config1FilePathClient                 = "../../configfile/config-client.yaml"
	// Config2FilePathEntityMatchers         = "../../configfile/config-entityMatchers.yaml"
	// Config3FilePathChannel                = "../../configfile/config-channel.yaml"
	// Config4FilePathOrderer                = "../../configfile/config-orderers.yaml"
	// Config5FilePathPeers                  = "../../configfile/config-peers.yaml"
	// Config6FilePathOrganizations          = "../../configfile/config-organizations.yaml"
	// Config7FilePathCertificateAuthorities = "../../configfile/config-certificateAuthorities.yaml"
)

// chaincode 相关
const (
	// ChainCodePath chaincode 存在的路径
	ChainCodePath = "fabric_car/chaincode/go/chaincode_example02"

	// ChainCodeVersion chaincode 版本
	ChainCodeVersion = "0.0.3"

	// ChainCodeName chaincode 名字
	ChainCodeName = "cars"

	// ChainCodePolicy 链码安装背书策略
	ChainCodePolicy = "OR ('baomaMSP.member','benchiMSP.member')"

	// ChainCodeInstantiateArgs 安装链码的 参数
	ChainCodeInstantiateArgs = "init,baoma,10,benchi,10"

	// ChainCodeInvokeArgs 链码调用invoke的参数
	ChainCodeInvokeArgs = "baoma,benchi,2"

	// ChainCodeQueryArgs 查询链码的参数
	ChainCodeQueryArgs = "baoma"

	// ChainCodeUpdrageArgs 升级链码的参数
	ChainCodeUpdrageArgs = "init"

	// PeerName 安装 实例化chaincode 的安装节点
	PeerName = "peer0.benchi.car.com" //"peer0.chenman.car.com" "192.168.56.101"
)

// 定义调用sdk的固定值ORG
const (
	// UserName WithUser
	UserName = "Admin"
	// OrgName WithOrg
	OrgName = "benchi"
)

const (
	// ChannelName channelName
	ChannelName = "mychannel"
	// ChannelConfigPath 频道的配置文件，一开始创世纪生成的
	ChannelConfigPath = "../../../GenesisFile/channel-artifacts/mychannel.tx"
)

const (
	// OrderURL 配置文件填写的 order url映射地址
	OrderURL = "grpc://39.105.34.144:7050"
	// OrdererEndpoint 排序节点
	OrdererEndpoint = "car0.car.com"
)
