package api

import (
	"contractor-services/model"
	"encoding/json"
	"log"
	"net/http"

	"github.com/uptrace/bunrouter"
)

const insertBankDetails = `
	INSERT INTO contractor_bank_detail (bank_name, branch, branch_address, account_number, account_type)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING *;
`

func (c *ContractorStore) AddContractorBankDetails(w http.ResponseWriter, req bunrouter.Request) error {

	var reqBody model.ContractorBankDetailsReq
	json.NewDecoder(req.Body).Decode(&reqBody)

	conn, err := c.db.Acquire(req.Context())
	if err != nil {
		log.Println(err.Error())
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return bunrouter.JSON(w, bunrouter.H{
			"message": "something went wrong",
			"success": false,
		})
	}
	defer conn.Release()

	row := conn.QueryRow(
		req.Context(), insertBankDetails, reqBody.BankName, reqBody.Branch, reqBody.BranchAddress, reqBody.AccountNumber, reqBody.AccountType,
	)

	var bankId int64

	var result model.ContractorBankDetailsReq
	err = row.Scan(&bankId, &result.BankName, &result.Branch, &result.BranchAddress, &result.AccountNumber, &result.AccountType)

	if err != nil {
		log.Println(err.Error())
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return bunrouter.JSON(w, bunrouter.H{
			"message": "failed to add bank details",
			"success": false,
		})
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	return bunrouter.JSON(w, result)
}
