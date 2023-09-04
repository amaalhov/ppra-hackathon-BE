package api

import (
	"contractor-services/model"
	"encoding/json"
	"log"
	"net/http"

	"github.com/uptrace/bunrouter"
)

const insertAddress = `
	INSERT INTO contractor_address (company_uuid, country, district_name, town, plot_number, street)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id, country, district_name, town, plot_number, street;
`

func (c *ContractorStore) AddContractorAddressDetails(w http.ResponseWriter, req bunrouter.Request) error {
	var reqBody model.ContractorAddressReq
	json.NewDecoder(req.Body).Decode(&reqBody)

	id := req.URL.Query().Get("company_uuid")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return bunrouter.JSON(w, bunrouter.H{
			"message": "company_uuid is required",
			"status":  false,
		})
	}

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
		req.Context(), insertAddress, id, reqBody.Country, reqBody.DistrictName, reqBody.Town, reqBody.PlotNumber, reqBody.Street,
	)

	var addressId int64

	var result model.ContractorAddressReq
	err = row.Scan(&addressId, &result.Country, &result.DistrictName, &result.Town, &result.PlotNumber, &result.Street)

	if err != nil {
		log.Println(err.Error())
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return bunrouter.JSON(w, bunrouter.H{
			"message": "failed to add address",
			"success": false,
		})
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	return bunrouter.JSON(w, result)
}
