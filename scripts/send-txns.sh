#!/bin/bash
set -e

apt-get update && apt-get install -y jq uuid-runtime

# Fruit Type

declare -a FruitType=('deciduous','citrus','subtropical')

# Deciduous Fruits

declare -a DeciduousFruitName=('Apricots' 'peaches' 'nectarines' 'plums' 'grapes' 'pears' 'apples')

CC_NAME="trace_fruits_cc"
RANDOM_NUMBER=600

send_txn() {
	echo
	echo "Channel :  '$2' , Chaincode  : '$CC_NAME'"
	echo
	echo ">>> Supply chain event :"
	echo
	echo "'$1'"
	echo
	local txn="$1"
	local channel="$2"

	echo "<<< Response:"
	echo
	peer chaincode invoke -o orderer.example.com:7050 \
		--tls $CORE_PEER_TLS_ENABLED \
		--cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem \
		-C $channel -n $CC_NAME \
		-c "$txn"
}

sellingnconsumtion_to_fruitsreceiver() {
	supply_event='{"Args":["EventSellingConsumptionRegister","{\"header\":{\"trace-id\":\"'$1'\",\"event-id\":\"'$(uuidgen)'\",\"gln\":\"9501101530911\",\"gln-location-type\":\"store\",\"gtin\":\"09501101530003\",\"gtin-trade-id-type\":\"pallet-trade-id\",\"batch-lot\":\"AB-123\"},\"attributes\":{\"metaInfo\":313,\"fruit-type\":\"deciduous\",\"fruit-name\":\"'${DeciduousFruitName[$((RANDOM % 7))]}'\",\"weight\":212,\"why\":\"selling-consumption-traceability\",\"when\":\"2021-01-14 18:47:22.151667 +1000 AEST\",\"expiration-datetime\":\"2021-07-15 16:28:22.1515 +1000 AEST\"}}"]}'
	send_txn "$supply_event" "fruitsreceiver"
	sleep 5
}

receiving_to_fruitsreceiver() {
	supply_event='{"Args":["EventReceivingRegister","{\"header\":{\"trace-id\":\"'$1'\",\"event-id\":\"'$(uuidgen)'\",\"gln\":\"9501101530911\",\"gln-location-type\":\"inbound\",\"gtin\":\"09501101530003\",\"gtin-trade-id-type\":\"pallet-trade-id\",\"batch-lot\":\"AB-123\"},\"attributes\":{\"metaInfo\":313,\"fruit-type\":\"deciduous\",\"fruit-name\":\"'${DeciduousFruitName[$((RANDOM % 7))]}'\",\"weight\":212,\"why\":\"receiving-traceability\",\"when\":\"2021-02-16 18:47:22.151667 +1000 AEST\",\"expiration-datetime\":\"2021-08-22 17:21:22.1515 +1000 AEST\"}}"]}'
	send_txn "$supply_event" "fruitsreceiver"
	sleep 5
}

transporting_to_fruitsreceiver() {
	supply_event='{"Args":["EventTransportingRegister","{\"header\":{\"trace-id\":\"'$1'\",\"event-id\":\"'$(uuidgen)'\",\"gln\":\"9501101530911\",\"gln-location-type\":\"outbound\",\"gtin\":\"09501101530003\",\"gtin-trade-id-type\":\"pallet-trade-id\",\"batch-lot\":\"AB-123\"},\"attributes\":{\"metaInfo\":313,\"fruit-type\":\"deciduous\",\"fruit-name\":\"'${DeciduousFruitName[$((RANDOM % 7))]}'\",\"weight\":212,\"why\":\"transporting-traceability\",\"when\":\"2021-03-18 18:47:22.151667 +1000 AEST\",\"expiration-datetime\":\"2021-09-23 13:27:22.1515 +1000 AEST\"}}"]}'
	send_txn "$supply_event" "fruitsreceiver"
	sleep 5
}

shipping_to_fruitsprovider() {
	supply_event='{"Args":["EventShippingRegister","{\"header\":{\"trace-id\":\"'$1'\",\"event-id\":\"'$(uuidgen)'\",\"gln\":\"9501101530911\",\"gln-location-type\":\"loading\",\"gtin\":\"09501101530003\",\"gtin-trade-id-type\":\"pallet-trade-id\",\"batch-lot\":\"AB-123\"},\"attributes\":{\"metaInfo\":313,\"fruit-type\":\"deciduous\",\"fruit-name\":\"'${DeciduousFruitName[$((RANDOM % 7))]}'\",\"weight\":212,\"why\":\"shipping-traceability\",\"when\":\"2021-04-20 18:47:22.151667 +1000 AEST\",\"expiration-datetime\":\"2021-10-18 12:17:11.1515 +1000 AEST\"}}"]}'
	send_txn "$supply_event" "fruitsprovider"
	sleep 5
}

packing_to_fruitsprovider() {
	supply_event='{"Args":["EventPackingRegister","{\"header\":{\"trace-id\":\"'$1'\",\"event-id\":\"'$(uuidgen)'\",\"gln\":\"9501101530911\",\"gln-location-type\":\"packing\",\"gtin\":\"09501101530003\",\"gtin-trade-id-type\":\"case-trade-id\",\"batch-lot\":\"AB-123\"},\"attributes\":{\"metaInfo\":313,\"fruit-type\":\"deciduous\",\"fruit-name\":\"'${DeciduousFruitName[$((RANDOM % 7))]}'\",\"weight\":212,\"why\":\"packing-repacking-traceability\",\"when\":\"2021-05-22 18:47:22.151667 +1000 AEST\",\"expiration-datetime\":\"2021-11-21 09:23:11.1515 +1000 AEST\"}}"]}'
	send_txn "$supply_event" "fruitsprovider"
	sleep 5
}

harvesting_to_fruitsprovider() {
	supply_event='{"Args":["EventHarvestingRegister","{\"header\":{\"trace-id\":\"'$1'\",\"event-id\":\"'$(uuidgen)'\",\"gln\":\"9501101530911\",\"gln-location-type\":\"planting\",\"gtin\":\"09501101530003\",\"gtin-trade-id-type\":\"crate-trade-id\",\"batch-lot\":\"AB-123\"},\"attributes\":{\"metaInfo\":313,\"fruit-type\":\"deciduous\",\"fruit-name\":\"'${DeciduousFruitName[$((RANDOM % 7))]}'\",\"weight\":212,\"why\":\"harvesting-traceability\",\"when\":\"2021-06-24 18:47:22.151667 +1000 AEST\",\"expiration-datetime\":\"2021-12-24 11:18:09.1515 +1000 AEST\"}}"]}'
	send_txn "$supply_event" "fruitsprovider"
	sleep 5
}

for ((i = 0; i < $RANDOM_NUMBER; ++i)); do
	trace_id=$(uuidgen)
	echo
	echo
	echo "Invoking chaincode to put supply chain events : $i "
	echo
	echo "Press [CTRL+C] to stop.."
	if [ $((RANDOM % 2)) == 1 ]; then
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
		harvesting_to_fruitsprovider "$trace_id"
		packing_to_fruitsprovider "$trace_id"
		shipping_to_fruitsprovider "$trace_id"
		transporting_to_fruitsreceiver "$trace_id"
		receiving_to_fruitsreceiver "$trace_id"
		sellingnconsumtion_to_fruitsreceiver "$trace_id"
	else
		ORG_NAME="popstar"
		CORE_PEER_ADDRESS=peer0.popstar.example.com:7051
		CORE_PEER_LOCALMSPID=PopstarMSP
		CORE_PEER_TLS_CERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/popstar.example.com/peers/peer0.popstar.example.com/tls/server.crt
		CORE_PEER_TLS_KEY_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/popstar.example.com/peers/peer0.popstar.example.com/tls/server.key
		CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/popstar.example.com/peers/peer0.popstar.example.com/tls/ca.crt
		CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/popstar.example.com/users/Admin@popstar.example.com/msp
		export CORE_PEER_TLS_CLIENTAUTHREQUIRED=true
		export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/$ORG_NAME.example.com/msp/tlscacerts/tlsca.$ORG_NAME.example.com-cert.pem
		export CORE_PEER_TLS_CLIENTKEY_FILE=$(find /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/$ORG_NAME.example.com/tlsca/ -type f -not -path *.pem)
		export CORE_PEER_TLS_CLIENTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/$ORG_NAME.example.com/tlsca/tlsca.$ORG_NAME.example.com-cert.pem
		harvesting_to_fruitsprovider "$trace_id"
		packing_to_fruitsprovider "$trace_id"
		shipping_to_fruitsprovider "$trace_id"
		transporting_to_fruitsreceiver "$trace_id"
		receiving_to_fruitsreceiver "$trace_id"
		sellingnconsumtion_to_fruitsreceiver "$trace_id"
	fi
done
