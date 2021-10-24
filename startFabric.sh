#!/bin/bash
#
# Exit on first error
set -e

# don't rewrite paths for Windows Git Bash users
export MSYS_NO_PATHCONV=1
starttime=$(date +%s)
CC_SRC_LANGUAGE="go"

CC_SRC_PATH="../chaincode/"



# launch network; create channel and join peer to channel
pushd ./hlf-network
./network.sh down
./network.sh up createChannel -ca -c trackingchannel -i 2.2.2 -s couchdb
./network.sh deployCC -c trackingchannel -ccn emptrackcc -ccv 1 -ccl ${CC_SRC_LANGUAGE} -ccp ${CC_SRC_PATH}
cp -rf ./organizations/peerOrganizations/org1.example.com/connection-org1.json ../api/
popd