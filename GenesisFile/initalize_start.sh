#!/bin/bash
#
# Copyright IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#


UP_DOWN="$1"
CH_NAME="$2"
CLI_TIMEOUT="$3"
IF_COUCHDB="$4"

: ${CLI_TIMEOUT:="10000"}

COMPOSE_FILE1=docker-orderer.yaml
COMPOSE_FILE2=docker-zxm.yaml
COMPOSE_PEER0_ORG1=peer0-baoma.yaml
COMPOSE_PEER0_ORG2=peer0-benchi.yaml
COMPOSE_PEER0_ORG3=peer0-falali.yaml
ORG1=baoma
ORG2=benchi
ORG3=falali
DOMAIN=car


function validateArgs () {
if [ -z "${UP_DOWN}" ]; then
echo "Option up / down / restart not mentioned"
printHelp
exit 1
fi
if [ -z "${CH_NAME}" ]; then
echo "setting to default channel 'mychannel'"
CH_NAME=mychannel
fi
}

function clearContainers () {
CONTAINER_IDS=$(docker ps -aq)
if [ -z "$CONTAINER_IDS" -o "$CONTAINER_IDS" = " " ]; then
echo "---- No containers available for deletion ----"
else
docker rm -f $CONTAINER_IDS
fi
}

function removeUnwantedImages() {
DOCKER_IMAGE_IDS=$(docker images | grep "dev\|none\|test-vp\|peer[0-9]-" | awk '{print $3}')
if [ -z "$DOCKER_IMAGE_IDS" -o "$DOCKER_IMAGE_IDS" = " " ]; then
echo "---- No images available for deletion ----"
else
docker rmi -f $DOCKER_IMAGE_IDS
fi
}



function replacePrivateKey () {
ARCH=`uname -s | grep Darwin`
if [ "$ARCH" == "Darwin" ]; then
OPTS="-it"
else
OPTS="-i"
fi
CURRENT_DIR=$PWD
cd crypto-config/peerOrganizations/${ORG1}.${DOMAIN}.com/ca/
PRIV_KEY=$(ls *_sk)
cd $CURRENT_DIR
sed $OPTS "s/CA1_PRIVATE_KEY/${PRIV_KEY}/g" ../${ORG1}/$COMPOSE_PEER0_ORG1

cd crypto-config/peerOrganizations/${ORG2}.${DOMAIN}.com/ca/
PRIV_KEY=$(ls *_sk)
cd $CURRENT_DIR
sed $OPTS "s/CA1_PRIVATE_KEY/${PRIV_KEY}/g" ../${ORG2}/$COMPOSE_PEER0_ORG2


cd crypto-config/peerOrganizations/${ORG3}.${DOMAIN}.com/ca/
PRIV_KEY=$(ls *_sk)
cd $CURRENT_DIR
sed $OPTS "s/CA1_PRIVATE_KEY/${PRIV_KEY}/g" ../${ORG3}/$COMPOSE_PEER0_ORG3

}


function replacePrivateKey1 () {
ARCH=`uname -s | grep Darwin`
if [ "$ARCH" == "Darwin" ]; then
OPTS="-it"
else
OPTS="-i"
fi
CURRENT_DIR=$PWD
cd crypto-config/peerOrganizations/${ORG1}.${DOMAIN}.com/ca/
PRIV_KEY=$(ls *_sk)
cd $CURRENT_DIR
sed $OPTS "s/${PRIV_KEY}/CA1_PRIVATE_KEY/g" ../${ORG1}/$COMPOSE_PEER0_ORG1

cd crypto-config/peerOrganizations/${ORG2}.${DOMAIN}.com/ca/
PRIV_KEY=$(ls *_sk)
cd $CURRENT_DIR
sed $OPTS "s/${PRIV_KEY}/CA1_PRIVATE_KEY/g" ../${ORG2}/$COMPOSE_PEER0_ORG2


cd crypto-config/peerOrganizations/${ORG3}.${DOMAIN}.com/ca/
PRIV_KEY=$(ls *_sk)
cd $CURRENT_DIR
sed $OPTS "s/${PRIV_KEY}/CA1_PRIVATE_KEY/g" ../${ORG3}/$COMPOSE_PEER0_ORG3

}



function networkUp () {
if [ -d "./crypto-config" ]; then
echo "crypto-config directory already exists."
else
#Generate all the artifacts that includes org certs, orderer genesis block,
# channel configuration transaction
./bin/cryptogen generate --config=./crypto-config.yaml
export FABRIC_CFG_PATH=$PWD
./bin/configtxgen -profile TwoOrgsOrdererGenesis -outputBlock ./channel-artifacts/genesis.block
./bin/configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./channel-artifacts/${CH_NAME}.tx   -channelID $CH_NAME
replacePrivateKey

#docker-compose -f $COMPOSE_FILE1 up -d
#docker-compose -f $COMPOSE_FILE2 up -d

#docker logs -f cli
fi
}

function networkDown () {
replacePrivateKey1
#docker-compose -f $COMPOSE_FILE1 down
#docker-compose -f $COMPOSE_FILE2 down

#Cleanup the chaincode containers
clearContainers

##Cleanup images
removeUnwantedImages

# remove orderer block and other channel configuration transactions and certs
rm -rf channel-artifacts/*.block channel-artifacts/*.tx crypto-config
}

function printHelp () {
    echo "==============================================================="
    echo "=============input up down perform=============================" 
    echo "==============================================================="     
}
validateArgs

#Create the network using docker compose
if [ "${UP_DOWN}" == "up" ]; then
networkUp
elif [ "${UP_DOWN}" == "down" ]; then ## Clear the network
networkDown
# elif [ "${UP_DOWN}" == "restart" ]; then ## Restart the network
# networkDown
# networkUp
else
 printHelp
exit 1
fi

