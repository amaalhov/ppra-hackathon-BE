package api

import (
	"contractor-services/model"
	"encoding/json"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/uptrace/bunrouter"
)

const insertCipaDetails = `
	INSERT INTO contractor (name, business_type, ownership_type, cipa_uin, is_registered_with_cipa, registration_date)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING uid;
`

const insertCompanyDetails = `
	INSERT INTO contractor (name, business_type, ownership_type, national_id, is_registered_with_cipa, registration_date)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING uid;
`

func (c *ContractorStore) AddContractorDetails(w http.ResponseWriter, req bunrouter.Request) error {
	var reqBody model.ContractorDetailsReq
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

	var row pgx.Row

	if reqBody.IsRegisteredWithCIPA {
		row = conn.QueryRow(
			req.Context(), insertCipaDetails, reqBody.NameOfCompany, reqBody.BusinessType, reqBody.OwnershipType, reqBody.CipaUin, reqBody.IsRegisteredWithCIPA, reqBody.RegistrationDate)
	} else {
		row = conn.QueryRow(
			req.Context(), insertCompanyDetails, reqBody.NameOfCompany, reqBody.BusinessType, reqBody.OwnershipType, reqBody.NationalIdNumber, reqBody.IsRegisteredWithCIPA, reqBody.RegistrationDate)
	}

	var companyUuid string

	//var result model.ContractorDetailsReq
	err = row.Scan(&companyUuid)

	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return bunrouter.JSON(w, bunrouter.H{
			"message": "failed to add company",
			"success": false,
		})
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	return bunrouter.JSON(w, bunrouter.H{"company_uuid": companyUuid})
}
