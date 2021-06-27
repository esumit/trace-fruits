#!/bin/bash
set -e

apt-get update && apt-get install -y jq uuid-runtime

CC_NAME="trace_fruits_cc"

init_trace_ledger_txn() {
  echo 
  echo "Init : chaincode '$CC_NAME' on '$2' channel" 
  echo  
  local txn="$1"
  local channel="$2"
  peer chaincode invoke -o orderer.example.com:7050  \
              --tls $CORE_PEER_TLS_ENABLED \
              --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem  \
              -C $channel -n $CC_NAME \
              -c "$txn"
}

init_fruitsreceiver() {
    init_trace_ledger_txn '{"function":"TraceEventInitLedger","Args":[]}' "fruitsreceiver"
    sleep 5
}

init_fruitsprovider() {
  init_trace_ledger_txn '{"function":"TraceEventInitLedger","Args":[]}' "fruitsprovider"
  sleep 5
}


echo "==========  TraceEventInitLedger =========="
        ORG_NAME="buttercup"
        CORE_PEER_ADDRESS=peer0.buttercup.example.com:7051
        CORE_PEER_LOCALMSPID=ButtercupMSP
        CORE_PEER_TLS_CERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/buttercup.example.com/peers/peer0.buttercup.example.com/tls/server.crt
        CORE_PEER_TLS_KEY_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/buttercup.example.com/peers/peer0.buttercup.example.com/tls/server.key
        CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/buttercup.example.com/peers/peer0.buttercup.example.com/tls/ca.crt
        CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/buttercup.example.com/users/Admin@buttercup.example.com/msp
        export CORE_PEER_TLS_CLIENTAUTHREQUIRED=true
        export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/$ORG_NAME.example.com/msp/tlscacerts/tlsca.$ORG_NAME.example.com-cert.pem
        export CORE_PEER_TLS_CLIENTKEY_FILE=$(find /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/buttercup.example.com/tlsca/ -type f -not -path *.pem)
        export CORE_PEER_TLS_CLIENTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/$ORG_NAME.example.com/tlsca/tlsca.$ORG_NAME.example.com-cert.pem

        init_fruitsreceiver
        init_fruitsprovider


