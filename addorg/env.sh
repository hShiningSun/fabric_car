CHANNEL=mychannel

#例如 ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/car.com/orderers/car0.car.com/msp/tlscacerts/tlsca.car.com-cert.pem
ORDERER_CA=  

#例如 ORDERER=orderer.example.com:7050
ORDERER=   

#例如 ORG_NAME=baoma
ORG_NAME= 

#例如 ORG_MSPID=baomaMSP
ORG_MSPID=



#====================== 下面配置一般固定的不需要更改 ===========================

# 获取初始最新通道的配置json名字
ORIGINAL_CONFIG_JSON=config.json

# 修改后的 通道配置json名字
MODIFYED_CONFIG_JSON=modified_config.json

# 最终第一个节点要签署的事务文件名字
UPDATE_ENVELOPE_PB=${ORG_NAME}_update_in_envelope.pb