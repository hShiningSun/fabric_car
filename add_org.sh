#!/bin/bash
#
# Copyright IBM Corp All Rights Reserved
#
# SPDX-License-Identifier: Apache-2.0
#

# This script extends the Hyperledger Fabric By Your First Network by
# adding a third organization to the network previously setup in the
# BYFN tutorial.
#

# prepending $PWD/../bin to PATH to ensure we are picking up the correct binaries
# this may be commented out to resolve installed version of tools if desired
#export PATH=${PWD}/../bin:${PWD}:$PATH
export FABRIC_CFG_PATH=${PWD}
ORG_ADD_NAME=chunyuhui
# 1-第一个背书节点 创建新增文件并且 签名，后面需要手动复制 事务文件到第二个节点的容器peer下
# 2-第二个 背书节点签署提交 事务文件
ACTION="$2"

# Print the usage message
function printHelp () {
  echo "Usage: "
  echo "  eyfn.sh up|down|restart|generate [-c <channel name>] [-t <timeout>] [-d <delay>] [-f <docker-compose-file>] [-s <dbtype>]"
  echo "  eyfn.sh -h|--help (print this message)"
  echo "    <mode> - one of 'up', 'down', 'restart' or 'generate'"
  echo "      - 'up' - bring up the network with docker-compose up"
  echo "      - 'down' - clear the network with docker-compose down"
  echo "      - 'restart' - restart the network"
  echo "      - 'generate' - generate required certificates and genesis block"
  echo "    -c <channel name> - channel name to use (defaults to \"mychannel\")"
  echo "    -t <timeout> - CLI timeout duration in seconds (defaults to 10)"
  echo "    -d <delay> - delay duration in seconds (defaults to 3)"
  echo "    -f <docker-compose-file> - specify which docker-compose file use (defaults to docker-compose-cli.yaml)"
  echo "    -s <dbtype> - the database backend to use: goleveldb (default) or couchdb"
  echo "    -l <language> - the chaincode language: golang (default) or node"
  echo "    -i <imagetag> - the tag to be used to launch the network (defaults to \"latest\")"
  echo
  echo "Typically, one would first generate the required certificates and "
  echo "genesis block, then bring up the network. e.g.:"
  echo
  echo "	eyfn.sh generate -c mychannel"
  echo "	eyfn.sh up -c mychannel -s couchdb"
  echo "	eyfn.sh up -l node"
  echo "	eyfn.sh down -c mychannel"
  echo
  echo "Taking all defaults:"
  echo "	eyfn.sh generate"
  echo "	eyfn.sh up"
  echo "	eyfn.sh down"
}

# Ask user for confirmation to proceed
function askProceed () {
  read -p "Continue? [Y/n] " ans
  case "$ans" in
    y|Y|"" )
      echo "proceeding ..."
    ;;
    n|N )
      echo "exiting..."
      exit 1
    ;;
    * )
      echo "invalid response"
      askProceed
    ;;
  esac
}

# Obtain CONTAINER_IDS and remove them
# TODO Might want to make this optional - could clear other containers
function clearContainers () {
  CONTAINER_IDS=$(docker ps -aq)
  if [ -z "$CONTAINER_IDS" -o "$CONTAINER_IDS" == " " ]; then
    echo "---- No containers available for deletion ----"
  else
    docker rm -f $CONTAINER_IDS
  fi
}

# Delete any images that were generated as a part of this setup
# specifically the following images are often left behind:
# TODO list generated image naming patterns
function removeUnwantedImages() {
  DOCKER_IMAGE_IDS=$(docker images | grep "dev\|none\|test-vp\|peer[0-9]-" | awk '{print $3}')
  if [ -z "$DOCKER_IMAGE_IDS" -o "$DOCKER_IMAGE_IDS" == " " ]; then
    echo "---- No images available for deletion ----"
  else
    docker rmi -f $DOCKER_IMAGE_IDS
  fi
}

# Generate the needed certificates, the genesis block and start the network.
# 再需要新增的peer 执行sh
function networkUp () {
      generateCerts
      generateChannelArtifacts
      createConfigTx
}

# Tear down running network
function networkDown () {
  docker-compose -f $COMPOSE_FILE -f $COMPOSE_FILE_ORG3 down --volumes
  docker-compose -f $COMPOSE_FILE -f $COMPOSE_FILE_ORG3 -f $COMPOSE_FILE_COUCH down --volumes
  # Don't remove containers, images, etc if restarting
  if [ "$MODE" != "restart" ]; then
    #Cleanup the chaincode containers
    clearContainers
    #Cleanup images
    removeUnwantedImages
    # remove orderer block and other channel configuration transactions and certs
    rm -rf channel-artifacts/*.block channel-artifacts/*.tx crypto-config ./org3-artifacts/crypto-config/ channel-artifacts/org3.json
    # remove the docker-compose yaml file that was customized to the example
    rm -f docker-compose-e2e.yaml
  fi

  # For some black-magic reason the first docker-compose down does not actually cleanup the volumes
  docker-compose -f $COMPOSE_FILE -f $COMPOSE_FILE_ORG3 down --volumes
  docker-compose -f $COMPOSE_FILE -f $COMPOSE_FILE_ORG3 -f $COMPOSE_FILE_COUCH down --volumes
}

# Use the CLI container to create the configuration transaction needed to add
# Org3 to the network
function createConfigTx () {
  echo
  echo "###############################################################"
  echo "####### Generate and submit config tx to add ${ORG_ADD_NAME} #############"
  echo "####### 生成提交配置上下文，用来新增组织: ${ORG_ADD_NAME} #############"
  echo "####### chenman 执行了，然后lixingxing执行 #############"
  echo "###############################################################"
docker exec cli scripts/step1${ORG_ADD_NAME}.sh $CHANNEL_NAME $CLI_DELAY $LANGUAGE $CLI_TIMEOUT ${ORG_ADD_NAME}
  if [ $? -ne 0 ]; then
    echo "ERROR !!!! Unable to create config tx"
    exit 1
  fi
}

# We use the cryptogen tool to generate the cryptographic material
# (x509 certs) for the new org.  After we run the tool, the certs will
# be parked in the BYFN folder titled ``crypto-config``.

# Generates Org3 certs using cryptogen tool
function generateCerts (){
  which bin1/cryptogen
  if [ "$?" -ne 0 ]; then
    echo "cryptogen tool not found. exiting"
    exit 1
  fi
  echo
  echo "###############################################################"
echo "##### Generate ${ORG_ADD_NAME} certificates using cryptogen tool #########"
  echo "###############################################################"

  (cd ${ORG_ADD_NAME}-artifacts
   set -x
   ../bin1/cryptogen generate --config=./${ORG_ADD_NAME}-crypto.yaml
   res=$?
   set +x
   if [ $res -ne 0 ]; then
     echo "Failed to generate certificates..."
     exit 1
   fi
  )
  echo
}

# Generate channel configuration transaction
function generateChannelArtifacts() {


  which bin1/configtxgen
  if [ "$?" -ne 0 ]; then
    echo "configtxgen tool not found. exiting"
    exit 1
  fi
  echo "##########################################################"
echo "#########  Generating ${ORG_ADD_NAME} config material ###############"
  echo "##########################################################"

#将加密文件复制过来
if [ ! -d "${ORG_ADD_NAME}-artifacts/crypto-config/peerOrganizations/${ORG_ADD_NAME}.example.com" ]; then
cp -Rf crypto-config/peerOrganizations/${ORG_ADD_NAME}.example.com ${ORG_ADD_NAME}-artifacts/crypto-config/peerOrganizations/
fi


  (cd ${ORG_ADD_NAME}-artifacts
   export FABRIC_CFG_PATH=$PWD
   set -x
../bin1/configtxgen -printOrg ${ORG_ADD_NAME}MSP > ../channel-artifacts/${ORG_ADD_NAME}.json
   res=$?
   set +x
   if [ $res -ne 0 ]; then
     echo "Failed to generate ${ORG_ADD_NAME} config material..."
     exit 1
   fi
  )
# 将创世生成的排序节点 文件复制过来有原文件就覆盖了
echo "##########################################################"
echo "#########  替换排序节点文件为创世纪生成的文件 ###############"
echo "##########################################################"
  cp -Rf crypto-config/ordererOrganizations ${ORG_ADD_NAME}-artifacts/crypto-config/
  echo
}


# If BYFN wasn't run abort
if [ ! -d crypto-config ]; then
  echo
  echo "ERROR: Please, run byfn.sh first."
  echo
  exit 1
fi

# Obtain the OS and Architecture string that will be used to select the correct
# native binaries for your platform
OS_ARCH=$(echo "$(uname -s|tr '[:upper:]' '[:lower:]'|sed 's/mingw64_nt.*/windows/')-$(uname -m | sed 's/x86_64/amd64/g')" | awk '{print tolower($0)}')
# timeout duration - the duration the CLI should wait for a response from
# another container before giving up
CLI_TIMEOUT=10
#default for delay
CLI_DELAY=3
# channel name defaults to "mychannel"
CHANNEL_NAME="mychannel"
# use this as the default docker-compose yaml definition
COMPOSE_FILE=docker-compose-cli.yaml
#
COMPOSE_FILE_COUCH=docker-compose-couch.yaml
# use this as the default docker-compose yaml definition
COMPOSE_FILE_ORG3=docker-compose-org3.yaml
#
COMPOSE_FILE_COUCH_ORG3=docker-compose-couch-org3.yaml
# use golang as the default language for chaincode
LANGUAGE=golang
# default image tag
IMAGETAG="latest"

# Parse commandline args
if [ "$1" = "-m" ];then	# supports old usage, muscle memory is powerful!
    shift
fi
MODE=$1;shift
# Determine whether starting, stopping, restarting or generating for announce
if [ "$MODE" == "up" ]; then
  EXPMODE="Starting"
elif [ "$MODE" == "down" ]; then
  EXPMODE="Stopping"
elif [ "$MODE" == "restart" ]; then
  EXPMODE="Restarting"
elif [ "$MODE" == "generate" ]; then
  EXPMODE="Generating certs and genesis block for"
else
  printHelp
  exit 1
fi
while getopts "h?c:t:d:f:s:l:i:" opt; do
  case "$opt" in
    h|\?)
      printHelp
      exit 0
    ;;
    c)  CHANNEL_NAME=$OPTARG
    ;;
    t)  CLI_TIMEOUT=$OPTARG
    ;;
    d)  CLI_DELAY=$OPTARG
    ;;
    f)  COMPOSE_FILE=$OPTARG
    ;;
    s)  IF_COUCHDB=$OPTARG
    ;;
    l)  LANGUAGE=$OPTARG
    ;;
    i)  IMAGETAG=$OPTARG
    ;;
  esac
done

# Announce what was requested

  if [ "${IF_COUCHDB}" == "couchdb" ]; then
        echo
        echo "${EXPMODE} with channel '${CHANNEL_NAME}' and CLI timeout of '${CLI_TIMEOUT}' seconds and CLI delay of '${CLI_DELAY}' seconds and using database '${IF_COUCHDB}'"
  else
        echo "${EXPMODE} with channel '${CHANNEL_NAME}' and CLI timeout of '${CLI_TIMEOUT}' seconds and CLI delay of '${CLI_DELAY}' seconds"
  fi
# ask for confirmation to proceed
askProceed

#Create the network using docker compose
if [ "${MODE}" == "up" ]; then
  networkUp
elif [ "${MODE}" == "down" ]; then ## Clear the network
  networkDown
elif [ "${MODE}" == "generate" ]; then ## Generate Artifacts
  generateCerts
  generateChannelArtifacts
  createConfigTx
elif [ "${MODE}" == "restart" ]; then ## Restart the network
  networkDown
  networkUp
else
  printHelp
  exit 1
fi
