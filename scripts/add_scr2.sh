CHANNEL_NAME="$1"
DELAY="$2"
LANGUAGE="$3"
TIMEOUT="$4"
ORG_ADD_NAME="$5"

COUNTER=1
MAX_RETRY=5

: ${CHANNEL_NAME:="mychannel"}
: ${DELAY:="3"}
: ${LANGUAGE:="golang"}
: ${TIMEOUT:="10"}


. scripts/utils.sh

ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/car.com/orderers/car0.car.com/msp/tlscacerts/tlsca.car.com-cert.pem
CC_SRC_PATH="github.com/chaincode/go/chaincode_example02"


echo "=== 生成更新事务 ==="
createConfigUpdate ${CHANNEL_NAME} config.json modified_config.json ${ORG_ADD_NAME}_update_in_envelope.pb


echo "=== 签名事务 ${ORG_ADD_NAME}_update_in_envelope.pb ==="
signConfigtxAsPeerOrg 1 ${ORG_ADD_NAME}_update_in_envelope.pb


echo "将事务${ORG_ADD_NAME}_update_in_envelope.pb 复制 出来"
echo "=== docker cp hostname:/opt/gopath/src/github.com/hyperledger/fabric/peer/${ORG_ADD_NAME}_update_in_envelope.pb /home/car/fabric_kafka/${ORG_ADD_NAME}-artifacts/${ORG_ADD_NAME}_update_in_envelope.pb ==="
echo "接下来 该第二个背书节点 来签署了"
echo "=== 将事务${ORG_ADD_NAME}_update_in_envelope.pb 复制 进去 ==="
echo "=== docker cp /home/car/fabric_kafka/{ORG_ADD_NAME}-artifacts/{ORG_ADD_NAME}_update_in_envelope.pb hostname:/opt/gopath/src/github.com/hyperledger/fabric/peer/{ORG_ADD_NAME}_update_in_envelope.pb ==="

echo "=== 复制好了 就下一步骤 在第二个背书节点  执行 add1.sh up3 "