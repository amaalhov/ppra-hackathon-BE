package api

import (
	"contractor-services/model"
	"encoding/json"
	"log"
	"net/http"

	"github.com/uptrace/bunrouter"
)

const insertSecretary = `
	INSERT INTO contractor_secretary (full_name, nationality, physical_address, box_address, appointment_date)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, full_name, nationality, physical_address, box_address, appointment_date;
`

func (c *ContractorStore) AddContractorSecretary(w http.ResponseWriter, req bunrouter.Request) error {
	var reqBody model.ContractorSecretaryReq
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
		req.Context(), insertSecretary, reqBody.Fullname, reqBody.Nationality, reqBody.PhysicalAddress, reqBody.BoxAddress, reqBody.AppointmentDate)

	var bankId int64

	var result model.ContractorContactReq
	err = row.Scan(&bankId, &result.Fullname, &result.Nationality, &result.PhysicalAddress, &result.BoxAddress, &result.AppointmentDate)

	if err != nil {
		log.Println(err.Error())
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return bunrouter.JSON(w, bunrouter.H{
			"message": "failed to add secretary details",
			"success": false,
		})
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	return bunrouter.JSON(w, result)
}
