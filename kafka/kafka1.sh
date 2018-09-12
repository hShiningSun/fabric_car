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

COMPOSE_FILE=docker-kafka1.yml



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

#启动
function networkUp () {
    networkDown
    docker-compose -f $COMPOSE_FILE up
}

#关闭
function networkDown () {
    docker-compose -f $COMPOSE_FILE down
    clearContainers
    #Cleanup images
    removeUnwantedImages
}

validateArgs

#Create the network using docker compose
if [ "${UP_DOWN}" == "up" ]; then
networkUp
elif [ "${UP_DOWN}" == "down" ]; then ## Clear the network
networkDown
elif [ "${UP_DOWN}" == "restart" ]; then ## Restart the network
networkDown
networkUp
else
printHelp
exit 1
fi

