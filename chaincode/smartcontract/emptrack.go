package emptrack

import (
	"encoding/json"
	"fmt"
	"time"

	"chaincode/smartcontract/errorcode"
	"chaincode/smartcontract/util"

	"github.com/golang/protobuf/ptypes"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	log "github.com/sirupsen/logrus"
)

const (
	DOCTYPE_ORG string = "org"
	DOCTYPE_EMP string = "employee"
)

// SmartContract of this fabric sample
type EmployeeTracking struct {
	contractapi.Contract
}

type OrganisationModel struct {
	Doctype  string `json:"docType"`
	ID       string `json:"orgID"`
	Name     string `json:"name"`
	Code     string `json:"code"`
	Location string `json:"location"`
}

type EmployeeModel struct {
	Doctype      string   `json:"docType"`
	ID           string   `json:"employeeID"`
	Name         string   `json:"name"`
	Email        string   `json:"email"`
	Designation  string   `json:"designation"`
	Organisation string   `json:"organisation"`
	Experience   string   `json:"experience"`
	Skills       []string `json:"skills"`
}

// StoreFileMetadata : Store the file metadata with hash as the kay
func (e *EmployeeTracking) AddOrganisation(ctx contractapi.TransactionContextInterface, orgString string) error {
	log.Debugf("%s()", util.FunctionName())

	var newOrg OrganisationModel

	err := json.Unmarshal([]byte(orgString), &newOrg)
	if err != nil {
		return errorcode.Internal.WithMessage("failed to parse New org information, %v", err).LogReturn()
	}

	// Query data from ledger to verify this is new hash
	txDataJSONasBytes, err := ctx.GetStub().GetState(newOrg.ID)
	if err != nil {
		return errorcode.Internal.WithMessage("failed to get data from ledger, %v", err).LogReturn()
	} else if txDataJSONasBytes != nil {
		return errorcode.DuplicateHash.WithMessage("Organisation '%s' Already registered", newOrg.ID).LogReturn()
	}
	newOrg.Doctype = DOCTYPE_ORG

	// Convert struct to Bytes to store in Ledger
	txDataJSONasBytes, err = json.Marshal(newOrg)
	if err != nil {
		return errorcode.Internal.WithMessage("failed to Marshal Organisation data , %v", err).LogReturn()
	}

	// Store data to ledger
	err = ctx.GetStub().PutState(newOrg.ID, txDataJSONasBytes)
	if err != nil {
		return errorcode.Internal.WithMessage("failed to store data, %v", err).LogReturn()
	}

	return nil
}

// GetFileMetadata : Returns the file metadata for a given hash
func (e *EmployeeTracking) ListAllOrganisation(ctx contractapi.TransactionContextInterface, orgName string) ([]*OrganisationModel, error) {
	log.Debugf("%s()", util.FunctionName())

	queryString := fmt.Sprintf(`{"selector":{"docType":"%s"}}`, DOCTYPE_ORG)

	log.Debugf("Query String : %v", queryString)

	return getQueryResultForQueryString(ctx, queryString)

}

// StoreFileMetadata : Store the file metadata with hash as the kay
func (e *EmployeeTracking) AddEmployee(ctx contractapi.TransactionContextInterface, orgString string) error {
	log.Debugf("%s()", util.FunctionName())

	var newEmployee EmployeeModel

	err := json.Unmarshal([]byte(orgString), &newEmployee)
	if err != nil {
		return errorcode.Internal.WithMessage("failed to parse New Employee information, %v", err).LogReturn()
	}

	// Query data from ledger to verify this is new hash
	txDataJSONasBytes, err := ctx.GetStub().GetState(newEmployee.ID)
	if err != nil {
		return errorcode.Internal.WithMessage("failed to get data from ledger, %v", err).LogReturn()
	} else if txDataJSONasBytes != nil {
		return errorcode.DuplicateHash.WithMessage("Employee '%s' Already registered", newEmployee.ID).LogReturn()
	}
	newEmployee.Doctype = DOCTYPE_EMP

	// Convert struct to Bytes to store in Ledger
	txDataJSONasBytes, err = json.Marshal(newEmployee)
	if err != nil {
		return errorcode.Internal.WithMessage("failed to Marshal Employee data , %v", err).LogReturn()
	}

	// Store data to ledger
	err = ctx.GetStub().PutState(newEmployee.ID, txDataJSONasBytes)
	if err != nil {
		return errorcode.Internal.WithMessage("failed to store data, %v", err).LogReturn()
	}

	return nil
}

type TransferEmpModel struct {
	EmpID        string `json:"empID"`
	Organisation string `json:"organisation"`
}

// StoreFileMetadata : Store the file metadata with hash as the kay
func (e *EmployeeTracking) TransferEmployee(ctx contractapi.TransactionContextInterface, transferDetails string) error {
	log.Debugf("%s()", util.FunctionName())

	var empTranfer TransferEmpModel

	err := json.Unmarshal([]byte(transferDetails), &empTranfer)
	if err != nil {
		return errorcode.Internal.WithMessage("failed to parse New Employee information, %v", err).LogReturn()
	}

	// Query data from ledger to verify this is new hash
	txDataJSONasBytes, err := ctx.GetStub().GetState(empTranfer.EmpID)
	if err != nil {
		return errorcode.Internal.WithMessage("failed to get data from ledger, %v", err).LogReturn()
	} else if txDataJSONasBytes == nil {
		return errorcode.DuplicateHash.WithMessage("Employee  with ID '%s' not found in the network", empTranfer.EmpID).LogReturn()
	}

	var emp EmployeeModel

	err = json.Unmarshal([]byte(txDataJSONasBytes), &emp)
	if err != nil {
		return errorcode.Internal.WithMessage("failed to parse Queried Employee information, %v", err).LogReturn()
	}

	if emp.Organisation == empTranfer.Organisation {
		return errorcode.Internal.WithMessage("failed to Tranfer Employee as destination organisation is same as current organisation").LogReturn()
	}

	queryString := fmt.Sprintf(`{"selector":{"docType":"%s"}}`, DOCTYPE_ORG)

	orgs, err := getQueryResultForQueryString(ctx, queryString)
	if err != nil {
		return errorcode.DuplicateHash.WithMessage("Failed to get Organisation details ").LogReturn()
	}

	validOrgName := false
	for _, v := range orgs {
		if v.Name == empTranfer.Organisation {
			validOrgName = true
		}
	}

	if !validOrgName {
		return errorcode.Internal.WithMessage("Organisation name '%v' is not registered in the network", empTranfer.Organisation).LogReturn()
	}

	emp.Organisation = empTranfer.Organisation

	// Convert struct to Bytes to store in Ledger
	txDataJSONasBytes, err = json.Marshal(emp)
	if err != nil {
		return errorcode.Internal.WithMessage("failed to Marshal Employee data , %v", err).LogReturn()
	}

	// Store data to ledger
	err = ctx.GetStub().PutState(emp.ID, txDataJSONasBytes)
	if err != nil {
		return errorcode.Internal.WithMessage("failed to store data, %v", err).LogReturn()
	}

	return nil
}

// GetFileMetadata : Returns the file metadata for a given hash
func (e *EmployeeTracking) ListAllEmployees(ctx contractapi.TransactionContextInterface, orgName string) ([]EmployeeModel, error) {
	log.Debugf("%s()", util.FunctionName())

	queryString := fmt.Sprintf(`{"selector":{"docType":"%s"}}`, DOCTYPE_EMP)

	log.Debugf("Employee search string %s", queryString)
	return getQueryResultForEmployeeString(ctx, queryString)

}

type HistoryQueryResult struct {
	Record    *EmployeeModel `json:"record"`
	TxId      string         `json:"txId"`
	Timestamp time.Time      `json:"timestamp"`
	IsDelete  bool           `json:"isDelete"`
}

type EmployeeTransferHistoryModel struct {
	From      string    `json:"fromCompany"`
	To        string    `json:"toCompany`
	Timestamp time.Time `json:"timestamp"`
	TxId      string    `json:"BlockchainID"`
	IsDelete  bool      `json:"isDeleted"`
}

// GetAssetHistory returns the chain of custody for an asset since issuance.
func (e *EmployeeTracking) GetEmployeeTransferHistory(ctx contractapi.TransactionContextInterface, assetID string) ([]EmployeeTransferHistoryModel, error) {
	log.Printf("GetAssetHistory: ID %v", assetID)

	resultsIterator, err := ctx.GetStub().GetHistoryForKey(assetID)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var records []HistoryQueryResult
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset EmployeeModel
		if len(response.Value) > 0 {
			err = json.Unmarshal(response.Value, &asset)
			if err != nil {
				return nil, err
			}
		} else {
			asset = EmployeeModel{
				ID: assetID,
			}
		}

		timestamp, err := ptypes.Timestamp(response.Timestamp)
		if err != nil {
			return nil, err
		}

		record := HistoryQueryResult{
			TxId:      response.TxId,
			Timestamp: timestamp,
			Record:    &asset,
			IsDelete:  response.IsDelete,
		}
		records = append(records, record)
	}

	var empMovement []EmployeeTransferHistoryModel

	for i, v := range records {
		newMv := EmployeeTransferHistoryModel{
			To:        v.Record.Organisation,
			From:      "",
			Timestamp: v.Timestamp,
			TxId:      v.TxId,
			IsDelete:  v.IsDelete,
		}
		if i == (len(records) - 1) {
			newMv.From = ""
		} else {
			newMv.From = records[i+1].Record.Organisation
		}
		empMovement = append(empMovement, newMv)
	}

	return empMovement, nil
}

// getQueryResultForQueryString executes the passed in query string.
// The result set is built and returned as a byte array containing the JSON results.
func getQueryResultForQueryString(ctx contractapi.TransactionContextInterface, queryString string) ([]*OrganisationModel, error) {
	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	return constructQueryResponseFromIterator(resultsIterator)
}

// constructQueryResponseFromIterator constructs a slice of assets from the resultsIterator
func constructQueryResponseFromIterator(resultsIterator shim.StateQueryIteratorInterface) ([]*OrganisationModel, error) {
	var assets []*OrganisationModel
	for resultsIterator.HasNext() {
		queryResult, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var asset OrganisationModel
		err = json.Unmarshal(queryResult.Value, &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}

	log.Debug("Get Organisation : %+v", assets)

	return assets, nil
}

// getQueryResultForQueryString executes the passed in query string.
// The result set is built and returned as a byte array containing the JSON results.
func getQueryResultForEmployeeString(ctx contractapi.TransactionContextInterface, queryString string) ([]EmployeeModel, error) {
	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	return constructQueryResponseFromEmployeeIterator(resultsIterator)
}

// constructQueryResponseFromIterator constructs a slice of assets from the resultsIterator
func constructQueryResponseFromEmployeeIterator(resultsIterator shim.StateQueryIteratorInterface) ([]EmployeeModel, error) {
	var assets []EmployeeModel
	for resultsIterator.HasNext() {
		queryResult, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var asset EmployeeModel
		err = json.Unmarshal(queryResult.Value, &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, asset)
	}

	return assets, nil
}
