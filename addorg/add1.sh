
# import env
. addorg/env.sh

function getOriginConfigJson () {
  echo "====== fetch channel new config and save the config_block.pb ======"
  if [ -z "$CORE_PEER_TLS_ENABLED" -o "$CORE_PEER_TLS_ENABLED" = "false" ]; then
      set -x
      peer channel fetch config config_block.pb -o $ORDERER -c $CHANNEL --cafile $ORDERER_CA
      set +x
    else
      set -x
      peer channel fetch config config_block.pb -o $ORDERER -c $CHANNEL --tls --cafile $ORDERER_CA
      set +x
  fi

  echo "====== config_block.pb ----> config_block.json ======"
  configtxlator proto_decode --input config_block.pb --type common.Block --output config_block.json

  echo "====== jq + config_block.json = config.json  ======"
  jq .data.data[0].payload.data.config config_block.json > $ORIGINAL_CONFIG_JSON
}

function createModifyConfigJson () {
  # config.json + genesis-${ORG_NAME}.json = modified_config.json
  echo "===== write new channel config json modified_config.json ======"
  set -x
  jq -s '.[0] * {"channel_group":{"groups":{"Application":{"groups": {"'$ORG_MSPID'":.[1]}}}}}' $ORIGINAL_CONFIG_JSON ./channel-artifacts/${ORG_NAME}.json > $MODIFYED_CONFIG_JSON
  set +x
}

function createConfigUpdate () {
  set -x
  # config.json ----> original_config.pb
  configtxlator proto_encode --input "${ORIGINAL_CONFIG_JSON}" --type common.Config > original_config.pb

  # modified_config.json ----> modified_config.pb
  configtxlator proto_encode --input "${MODIFYED_CONFIG_JSON}" --type common.Config > modified_config.pb

  # modified_config.pb - original_config.pb = config_update.pb
  configtxlator compute_update --channel_id "${CHANNEL}" --original original_config.pb --updated modified_config.pb > config_update.pb

  # config_update.pb ----> config_update.json
  configtxlator proto_decode --input config_update.pb  --type common.ConfigUpdate > config_update.json

  # payloadJSON + config_update.json = config_update_in_envelope.json
  echo '{"payload":{"header":{"channel_header":{"channel_id":"'$CHANNEL'", "type":2}},"data":{"config_update":'$(cat config_update.json)'}}}' | jq . > config_update_in_envelope.json

  # config_update_in_envelope.json ----> ${ORG_NAME}_update_in_envelope.pb
  configtxlator proto_encode --input config_update_in_envelope.json --type common.Envelope > "${UPDATE_ENVELOPE_PB}"
  set +x
}


function signConfigtx () {
  peer channel signconfigtx -f $UPDATE_ENVELOPE_PB
}



unction validateArgs () {
if [ -z "${ORDERER}" ]; then
echo "====== env.sh please set ORDERER value ======"
exit 1
fi

if [ -z "${CHANNEL}" ]; then
echo "====== env.sh please set CHANNEL value ======"
exit 1
fi

if [ -z "${ORDERER_CA}" ]; then
echo "====== env.sh please set ORDERER_CA value ======"
exit 1
fi

if [ -z "${ORG_MSPID}" ]; then
echo "====== env.sh please set ORG_MSPID value ======"
exit 1
fi

if [ -z "${ORG_NAME}" ]; then
echo "====== env.sh please set ORG_NAME value ======"
exit 1
fi

}



function printHelp () {
    echo "==============================================================="
    echo "== copy ${UPDATE_ENVELOPE_PB} other MSP_Peer sign and update =="
    echo "==============================================================="
    echo "==============================================================="
    echo "==== peer channel signconfigtx -f ${UPDATE_ENVELOPE_PB} ======="  
    echo "==============================================================="
    echo "==== peer channel update -f ${UPDATE_ENVELOPE_PB} -c ${CHANNEL} -o ${ORDERER}  --cafile ${ORDERER_CA} ====="
    echo "==============================================================="  
}

validateArgs

# 得到通道 初始配置json
getOriginConfigJson

# 创建通道修改后的 配置json
createModifyConfigJson

# 创建要更新的配置事务
createConfigUpdate

# 本节点对这个更新事务 进行签名
signConfigtx

