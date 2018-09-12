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




setGlobals 0 2
set -x
peer channel update -f ${ORG_ADD_NAME}_update_in_envelope.pb -c ${CHANNEL_NAME} -o car0.car.com:7050  --cafile ${ORDERER_CA}
set +x

echo "=== 第二个背书节点签署事务成功 ===""