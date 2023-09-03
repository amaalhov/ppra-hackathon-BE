package api

import (
	"contractor-services/model"
	"encoding/json"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/uptrace/bunrouter"
)

func (c *ContractorStore) AddContractorShareHolder(w http.ResponseWriter, req bunrouter.Request) error {
	var reqBody []model.ContractorShareHolderReq

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

	var copyData [][]interface{}

	for _, data := range reqBody {
		copyData = append(copyData, []interface{}{data.Fullname, data.Nationality, data.BoxAddress, data.PhysicalAddress, data.AppointmentDate})
	}

	copyCount, err := conn.CopyFrom(
		req.Context(),
		pgx.Identifier{"contractor_shareholder"},
		[]string{"full_name", "nationality", "postal_address", "physical_address", "appointment_date"},
		pgx.CopyFromRows(copyData),
	)

	if err != nil {
		log.Println(err.Error())
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return bunrouter.JSON(w, bunrouter.H{
			"message": "failed to add shareholders",
			"success": false,
		})
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	return bunrouter.JSON(w, bunrouter.H{
		"message":      "successfully added shareholders",
		"number_added": copyCount,
	})
}
