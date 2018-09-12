#!/bin/bash
#
# Copyright IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

# This script is designed to be run in the ${ORG_ADD_NAME}cli container as the
# first step of the EYFN tutorial.  It creates and submits a
# configuration transaction to add ${ORG_ADD_NAME} to the network previously
# setup in the BYFN tutorial.
#

CHANNEL_NAME=mychannel
# DELAY="$2"
# LANGUAGE="$3"
# TIMEOUT="$4"
ORG_ADD_NAME=chunyuhui
# : ${CHANNEL_NAME:="mychannel"}
# : ${DELAY:="3"}
# : ${LANGUAGE:="golang"}
# : ${TIMEOUT:="10"}
# LANGUAGE=`echo "$LANGUAGE" | tr [:upper:] [:lower:]`
# COUNTER=1
# MAX_RETRY=5
ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/car.com/orderers/car0.car.com/msp/tlscacerts/tlsca.car.com-cert.pem

CC_SRC_PATH="github.com/chaincode/go/chaincode_example02"


# import utils
. scripts/utils.sh
#
#echo
#echo "========= Creating config transaction to add ${ORG_ADD_NAME} to network =========== "
#echo "========= 创建 事务配置 来 新增 ${ORG_ADD_NAME} 的网络 =========== "
#echo

#echo "Installing jq"
#apt-get -y update && apt-get -y install jq

## Fetch the config for the channel, writing it to config.json
#fetchChannelConfig ${CHANNEL_NAME} config.json
#
## Modify the configuration to append the new org
#set -x
#jq -s '.[0] * {"channel_group":{"groups":{"Application":{"groups": {"${ORG_ADD_NAME}MSP":.[1]}}}}}' config.json ./channel-artifacts/${ORG_ADD_NAME}.json > modified_config.json
#set +x
#
## Compute a config update, based on the differences between config.json and modified_config.json, write it as a transaction to ${ORG_ADD_NAME}_update_in_envelope.pb
#createConfigUpdate ${CHANNEL_NAME} config.json modified_config.json ${ORG_ADD_NAME}_update_in_envelope.pb
#
#echo
#echo "========= Config transaction to add ${ORG_ADD_NAME} to network created ===== "
#echo
#
#echo "Signing config transaction 签名 这是新增配置事务"
#echo
#signConfigtxAsPeerOrg 1 ${ORG_ADD_NAME}_update_in_envelope.pb
#
#
#echo "将{ORG_ADD_NAME}_update_in_envelope.pb 复制出去，让第二台机子能够访问到"
#docker cp hostname:/opt/gopath/src/github.com/hyperledger/fabric/peer/{ORG_ADD_NAME}_update_in_envelope.pb /home/car/fabric_kafka/{ORG_ADD_NAME}-artifacts/{ORG_ADD_NAME}_update_in_envelope.pb
#
#echo
#echo "========= Submitting transaction from a different peer (peer0.org2) which also signs it ========= "
#echo "========= peer0.org2-benchi 也签署该事务 ========= "
#echo

#echo "将{ORG_ADD_NAME}_update_in_envelope.pb 复制过来，让第二台机子能够访问到"
#docker cp /home/car/fabric_kafka/chunyuhui-artifacts/chunyuhui_update_in_envelope.pb hostname:/opt/gopath/src/github.com/hyperledger/fabric/peer/chunyuhui_update_in_envelope.pb

setGlobals 0 2
set -x
peer channel update -f ${ORG_ADD_NAME}_update_in_envelope.pb -c ${CHANNEL_NAME} -o car0.car.com:7050  --cafile ${ORDERER_CA}
set +x

echo
echo "========= Config transaction to add ${ORG_ADD_NAME} to network submitted! =========== "
echo "========= 配置事务添加${ORG_ADD_NAME} 提交到网络! =========== "
echo

echo "==== 执行 docker exec cli scripts/step2${ORG_ADD_NAME}.sh === 加入频道 "
exit 0
