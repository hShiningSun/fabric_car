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

ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/car.com/orderers/car0.car.com/msp/tlscacerts/tlsca.car.com-cert.pem
CC_SRC_PATH="github.com/chaincode/go/chaincode_example02"

echo "==== 先安装 jq 工具 ===="
apt-get -y update && apt-get -y install jq

echo "导入 工具方法类 utils"
. scripts/utils.sh


echo "==== 获取现在的channel的配置文件 ===="
fetchChannelConfig ${CHANNEL_NAME} config.json

if [ ! -d "config.json" ]; then
    echo "===获取channel 当前配置成功=="
else 
echo "=== 获取channel配置 config.json 失败，自己去检查"
exit 1
fi


echo "=== 恭喜你进入最复杂的一步骤,请将config.json复制出来，再加入创世纪新增的json ==="
echo "=== 将创世生成的${ORG_ADD_NAME}.json 里面org信息截取出来 === "
echo "=== 添加到 config.json 并且取上一个新名字 modified_config.json ==="
echo "=== 将 modified_config.json 复制回 docke 容器的 peer 下==="

echo "=== 弄好之后--恭喜你，朋友，可以执行下一个脚本add_scr2.sh 生成 更新事务了 ===="

