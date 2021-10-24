#!/bin/bash

source scripts/utils.sh
FABRIC_CFG_PATH=$PWD/../config/

. scripts/envVar.sh

CHANNEL_NAME="trackingchannel"
CC_NAME="emptrackcc"

successInvokeTx() {
  setGlobals 1
  parsePeerConnectionParameters 1 2
  res=$?
  verifyResult $res "Invoke transaction failed on channel '$CHANNEL_NAME' due to uneven number of peer and org parameters "
  #set -x
  
  fcn_call='{"function":"AddOrganisation","Args":["{\"orgID\":\"string\",\"name\":\"string\",\"code\":\"string\",\"location\":\"string\"}"]}'

  echo -e ""
  echo -e "Scenario 1 : ${C_YELLOW}Valid Invoke Transactions${C_RESET}"
  echo ""
  echo -e "         ${C_BLUE}Function${C_RESET} : 'StoreFileMetadata'"
  echo ""
  echo -e '         args     : ["{\"orgID\":\"string\",\"name\":\"string\",\"code\":\"string\",\"location\":\"string\"}"]'
  echo ""
  echo -e "         ${C_BLUE}Command${C_RESET} : "
set -x
  peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile $ORDERER_CA -C $CHANNEL_NAME -n ${CC_NAME} $PEER_CONN_PARMS -c ${fcn_call} >&log.txt
  { set +x; } 2>/dev/null
    echo ""
    echo ""
  echo -e "         ${C_BLUE}Output${C_RESET}   : ${C_GREEN}$(cat log.txt)${C_RESET}"


fcn_call='{"function":"AddOrganisation","Args":["{\"orgID\":\"string1\",\"name\":\"string1\",\"code\":\"string\",\"location\":\"string\"}"]}'

  echo -e ""
  echo -e "Scenario 1 : ${C_YELLOW}Valid Invoke Transactions${C_RESET}"
  echo ""
  echo -e "         ${C_BLUE}Function${C_RESET} : 'StoreFileMetadata'"
  echo ""
  echo -e '         args     : ["{\"orgID\":\"string\",\"name\":\"string\",\"code\":\"string\",\"location\":\"string\"}"]'
  echo ""
  echo -e "         ${C_BLUE}Command${C_RESET} : "
set -x
  peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile $ORDERER_CA -C $CHANNEL_NAME -n ${CC_NAME} $PEER_CONN_PARMS -c ${fcn_call} >&log.txt
  { set +x; } 2>/dev/null
    echo ""
    echo ""
  echo -e "         ${C_BLUE}Output${C_RESET}   : ${C_GREEN}$(cat log.txt)${C_RESET}"

}



successChaincodeQuery() {
  echo -e ""
  echo -e "Scenario 2 : ${C_YELLOW}Query to get File Metadata Transactions${C_RESET}"
  echo ""
  echo -e "         ${C_BLUE}Function${C_RESET} : 'GetFileMetadata'"
  echo ""
  echo -e "         ${C_BLUE}args${C_RESET}     : ['69dde88229fdcb24c05a10e2be2c1e54fb6ed9b36dab733de997d36c63576c3f']"
  echo ""
  echo -e "         ${C_BLUE}Command${C_RESET} : "
  set -x
    peer chaincode query -C $CHANNEL_NAME -n ${CC_NAME} -c '{"Args":["ListAllOrganisation"," "]}' >&log.txt
   { set +x; } 2>/dev/null
  echo ""
  echo ""
  echo -e "         ${C_BLUE}Output${C_RESET}   : ${C_GREEN}$(cat log.txt)${C_RESET}"
 
}

EmpInvokeTx() {
#   setGlobals 1
#   parsePeerConnectionParameters 1 2
  #set -x
  fcn_call='{"function":"AddEmployee","Args":["{\"employeeID\":\"estring\",\"name\":\"string\",\"email\":\"string\",\"designation\":\"string\",\"organisation\":\"string\",\"experience\":\"string\",\"skills\":[\"string\"]}"]}'
  echo -e ""
  echo -e "Scenario 3 : ${C_YELLOW}Duplicate invoke Transactions -- should result in error${C_RESET}"
  echo ""
  echo -e "         ${C_BLUE}Function${C_RESET} : 'StoreFileMetadata'"
  echo ""
  echo -e '         args     : ["{\"employeeID\":\"estring\",\"name\":\"string\",\"email\":\"string\",\"designation\":\"string\",\"organisation\":\"string\",\"experience\":\"string\",\"skills\":[\"string\"]}"]'
  echo ""
  echo -e "         ${C_BLUE}Command${C_RESET} : "
  set -x
  peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile $ORDERER_CA -C $CHANNEL_NAME -n ${CC_NAME} $PEER_CONN_PARMS -c ${fcn_call} >&log.txt
  { set +x; } 2>/dev/null
  echo ""
  echo ""
  echo -e "         ${C_BLUE}Output${C_RESET}   : ${C_GREEN}$(cat log.txt)${C_RESET}"

}


EmpTranferInvokeTx() {
#   setGlobals 1
#   parsePeerConnectionParameters 1 2
  #set -x
  fcn_call='{"function":"TransferEmployee","Args":["{\"empID\":\"estring\",\"organisation\":\"string1\"}"]}'
  echo -e ""
  echo -e "Scenario 3 : ${C_YELLOW}Duplicate invoke Transactions -- should result in error${C_RESET}"
  echo ""
  echo -e "         ${C_BLUE}Function${C_RESET} : 'StoreFileMetadata'"
  echo ""
  echo -e '         args     : ["{\"empID\":\"estring\",\"organisation\":\"string1\"}"]}"]'
  echo ""
  echo -e "         ${C_BLUE}Command${C_RESET} : "
  set -x
  peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile $ORDERER_CA -C $CHANNEL_NAME -n ${CC_NAME} $PEER_CONN_PARMS -c ${fcn_call} >&log.txt
  { set +x; } 2>/dev/null
  echo ""
  echo ""
  echo -e "         ${C_BLUE}Output${C_RESET}   : ${C_GREEN}$(cat log.txt)${C_RESET}"

}

EmpChaincodeQuery() {
  echo -e ""
  echo -e "Scenario 4 : ${C_YELLOW}Invalid Hash Query -- should result in error${C_RESET}"
  echo ""
  echo -e "         ${C_BLUE}Function${C_RESET} : 'ListAllEmployees'"
  echo ""
  echo -e "         ${C_BLUE}args${C_RESET}     : [' ']"
  echo ""
  echo -e "         ${C_BLUE}Command${C_RESET} : "
  set -x
    peer chaincode query -C $CHANNEL_NAME -n ${CC_NAME} -c '{"Args":["ListAllEmployees"," "]}' >&log.txt
   { set +x; } 2>/dev/null
  echo ""
  echo ""
  echo -e "         ${C_BLUE}Output${C_RESET}   : ${C_GREEN}$(cat log.txt)${C_RESET}"
 
}


EmpHistoryChaincodeQuery() {
  echo -e ""
  echo -e "Scenario 4 : ${C_YELLOW}Invalid Hash Query -- should result in error${C_RESET}"
  echo ""
  echo -e "         ${C_BLUE}Function${C_RESET} : 'GetEmployeeTransferHistory'"
  echo ""
  echo -e "         ${C_BLUE}args${C_RESET}     : [' ']"
  echo ""
  echo -e "         ${C_BLUE}Command${C_RESET} : "
  set -x
    peer chaincode query -C $CHANNEL_NAME -n ${CC_NAME} -c '{"Args":["GetEmployeeTransferHistory","estring"]}' >&log.txt
   { set +x; } 2>/dev/null
  echo ""
  echo ""
  echo -e "         ${C_BLUE}Output${C_RESET}   : ${C_GREEN}$(cat log.txt)${C_RESET}"
 
}

echo "-------------------------------------------------------------"
echo "--------------- Blockchain Transactions ---------------------"
echo "-------------------------------------------------------------"

successInvokeTx

# EmpInvokeTx

EmpTranferInvokeTx

# echo "-------------------------------------------------------------"
# echo "-------------------------------------------------------------"
echo ""
echo ""
echo "------------------------------------------------------------------------------"
sleep 2
# successChaincodeQuery
# EmpChaincodeQuery

EmpHistoryChaincodeQuery


echo ""
echo ""
echo "------------------------------------------------------------------------------"
sleep 2
# DuplicateInvokeTx
echo ""
echo ""
echo "------------------------------------------------------------------------------"
sleep 2 
# failedChaincodeQuery
echo ""
echo ""
echo "------------------------------ END --------------------------------------------"

# echo "-------------------------------------------------------------"
# echo "----- Successfull Chaincode Query to get Has Metadata -------"
# echo "-------------------------------------------------------------"

# successChaincodeQuery

# echo "-------------------------------------------------------------"
# echo "-------------------------------------------------------------"
# echo ""
# echo ""
# sleep 1

# echo "-------------------------------------------------------------"
# echo "--------- Duplicate Hash Transaction invokation -------------"
# echo "-------------------------------------------------------------"

# DuplicateInvokeTx

# echo "-------------------------------------------------------------"
# echo "-------------------------------------------------------------"

# sleep 2

# echo "-------------------------------------------------------------"
# echo "--------- Failed Chaincode Query  ----------------"
# echo "-------------------------------------------------------------"
# failedChaincodeQuery

# echo "-------------------------------------------------------------"
# echo "-------------------  END ------------------------------------"
