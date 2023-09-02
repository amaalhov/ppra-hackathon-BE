package api

import (
	"contractor-services/model"
	"encoding/json"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/uptrace/bunrouter"
)

type ContractorStore struct {
	db *pgxpool.Pool
}

func NewContractorStore(db *pgxpool.Pool) *ContractorStore {
	return &ContractorStore{db: db}
}

func (c *ContractorStore) AddContractorDetails(w http.ResponseWriter, req bunrouter.Request) error {
	var reqBody model.ContractorDetailsReq

	json.NewDecoder(req.Body).Decode(&reqBody)

	return bunrouter.JSON(w, reqBody)
}

const insertAffiliates = `
	INSERT INTO contractor_affiliate (full_name, address, attachment_url)
	VALUES ($1, $2, $3)
	RETURNING *;
`

func (c *ContractorStore) AddAffiliates(w http.ResponseWriter, req bunrouter.Request) error {
	var reqBody model.ContractorAffiliateReq
	json.NewDecoder(req.Body).Decode(&reqBody)

	conn, err := c.db.Acquire(req.Context())
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return bunrouter.JSON(w, bunrouter.H{
			"message": "something went wrong",
			"success": false,
		})
	}

	var copyData [][]interface{}

	for _, data := range reqBody.ContractorAffiliates {
		copyData = append(copyData, []interface{}{data.Fullname, data.Address, data.AttachmentUlr})
	}

	copyCount, err := conn.CopyFrom(
		req.Context(),
		pgx.Identifier{"contractor_affiliates"},
		[]string{"full_name", "address", "attachment_url"},
		pgx.CopyFromRows(copyData),
	)

	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return bunrouter.JSON(w, bunrouter.H{
			"message": "failed to add affiliates",
			"success": false,
		})
	}

	w.WriteHeader(http.StatusCreated)
	return bunrouter.JSON(w, bunrouter.H{
		"message":                    "successfully added affiliates",
		"number of affiliates added": copyCount,
	})
}

const insertAddress = `
	INSERT INTO contractor_address (country, district_name, town, plot_number, street)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING *;
`

func (c *ContractorStore) AddContractorAddressDetails(w http.ResponseWriter, req bunrouter.Request) error {
	var reqBody model.ContractorAddressReq
	json.NewDecoder(req.Body).Decode(&reqBody)

	conn, err := c.db.Acquire(req.Context())
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return bunrouter.JSON(w, bunrouter.H{
			"message": "something went wrong",
			"success": false,
		})
	}

	row := conn.QueryRow(
		req.Context(), insertAddress, reqBody.Country, reqBody.DistrictName, reqBody.Town, reqBody.PlotNumber, reqBody.Street,
	)

	var addressId int64

	var result model.ContractorAddressReq
	err = row.Scan(&addressId, &result.Country, &result.DistrictName, &result.Town, &result.PlotNumber, &result.Street)

	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return bunrouter.JSON(w, bunrouter.H{
			"message": "failed to add address",
			"success": false,
		})
	}

	w.WriteHeader(http.StatusCreated)
	return bunrouter.JSON(w, result)
}

const insertContacts = `
	INSERT INTO contractor_contact
	(first_name, middle_name, last_name, date_of_birth, email, cellphone, telephone, business_number, contact_type)
	VALUES
	($1, $2, $3, $4, $5, $6, $7, $8)
	($9, $11, $12, $13, $14, $15, $16, $17)
	RETURNING id;
`

const insertContactsAddress = `
	INSERT INTO address
	(first_name, middle_name, last_name, date_of_birth, email, cellphone, telephone, business_number, contact_type)
	VALUES
	($1, $2, $3, $4, $5, $6, $7, $8)
	($9, $11, $12, $13, $14, $15, $16, $17)
	RETURNING id;
`

func (c *ContractorStore) AddContractorContacts(w http.ResponseWriter, req bunrouter.Request) error {
	var reqBody model.ContractorContactReq

	json.NewDecoder(req.Body).Decode(&reqBody)

	return bunrouter.JSON(w, reqBody)
}

func (c *ContractorStore) AddContractorContactDetails(w http.ResponseWriter, req bunrouter.Request) error {

	var reqBody model.ContractorContactReq

	json.NewDecoder(req.Body).Decode(&reqBody)

	return bunrouter.JSON(w, reqBody)
}

func (c *ContractorStore) AddContractorPostalAddress(w http.ResponseWriter, req bunrouter.Request) error {

	var reqBody model.ContractorAddressReq

	json.NewDecoder(req.Body).Decode(&reqBody)

	return bunrouter.JSON(w, reqBody)
}

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
		w.WriteHeader(http.StatusInternalServerError)
		return bunrouter.JSON(w, bunrouter.H{
			"message": "something went wrong",
			"success": false,
		})
	}

	row := conn.QueryRow(
		req.Context(), insertBankDetails, reqBody.BankName, reqBody.Branch, reqBody.BranchAddress, reqBody.AccountNumber, reqBody.AccountType,
	)

	var bankId int64

	var result model.ContractorBankDetailsReq
	err = row.Scan(&bankId, &result.BankName, &result.Branch, &result.BranchAddress, &result.AccountNumber, &result.AccountType)

	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return bunrouter.JSON(w, bunrouter.H{
			"message": "failed to add bank details",
			"success": false,
		})
	}

	w.WriteHeader(http.StatusCreated)
	return bunrouter.JSON(w, result)
}

func (c *ContractorStore) AddContractorDirector(w http.ResponseWriter, req bunrouter.Request) error {
	var reqBody model.ContractorDirectorReq

	json.NewDecoder(req.Body).Decode(&reqBody)

	return bunrouter.JSON(w, reqBody)
}

func (c *ContractorStore) AddContractorShareHolder(w http.ResponseWriter, req bunrouter.Request) error {
	var reqBody model.ContractorDirectorReq

	json.NewDecoder(req.Body).Decode(&reqBody)

	return bunrouter.JSON(w, reqBody)
}

func (c *ContractorStore) AddContractorEmployees(w http.ResponseWriter, req bunrouter.Request) error {
	var reqBody model.ContractorBankDetailsReq

	json.NewDecoder(req.Body).Decode(&reqBody)

	return bunrouter.JSON(w, reqBody)
}

const insertSecretary = `
	INSERT INTO contractor_contact (full_name, nationality, gov_id, gov_id_type, telephone, business_number, cellphone, email, contact_type)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	RETURNING *;
`

func (c *ContractorStore) AddContractorSecretary(w http.ResponseWriter, req bunrouter.Request) error {
	var reqBody model.ContractorContactReq
	json.NewDecoder(req.Body).Decode(&reqBody)

	conn, err := c.db.Acquire(req.Context())
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return bunrouter.JSON(w, bunrouter.H{
			"message": "something went wrong",
			"success": false,
		})
	}

	row := conn.QueryRow(
		req.Context(), insertSecretary, reqBody.Fullname, reqBody.Nationality, reqBody.GovernmentIdentification, reqBody.GovernmentIdentificationType, reqBody.Telephone, reqBody.BusinessPhoneNumber, reqBody.CellNumber, reqBody.Email, "SECRETARY")

	var bankId int64

	var result model.ContractorContactReq
	err = row.Scan(&bankId, &result.Fullname, &result.Nationality, &result.GovernmentIdentification, &result.GovernmentIdentificationType, &result.Telephone, &result.BusinessPhoneNumber, &result.CellNumber, &result.Email)

	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return bunrouter.JSON(w, bunrouter.H{
			"message": "failed to add bank details",
			"success": false,
		})
	}

	w.WriteHeader(http.StatusCreated)
	return bunrouter.JSON(w, result)
}
