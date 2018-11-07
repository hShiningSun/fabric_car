. env.sh
../GenesisFile/bin/cryptogen generate --config=./crypto-config.yaml #./${ORG_NAME}-crypto.yaml
cp -Rf crypto-config/peerOrganizations/ ../GenesisFile/crypto-config/peerOrganizations
export FABRIC_CFG_PATH=$PWD
../GenesisFile/bin/configtxgen -printOrg ${ORG_NAME}MSP > ../GenesisFile/channel-artifacts/${ORG_NAME}.json