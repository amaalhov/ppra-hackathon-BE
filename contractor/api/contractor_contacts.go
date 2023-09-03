package api

import (
	"contractor-services/model"
	"encoding/json"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/uptrace/bunrouter"
)

func (c *ContractorStore) AddContractorContacts(w http.ResponseWriter, req bunrouter.Request) error {
	var reqBody []model.ContractorContactReq
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

	var copyData [][]interface{}

	for _, data := range reqBody {
		copyData = append(copyData, []interface{}{
			id, data.FirstName, data.MiddleName, data.LastName, data.DateOfBirth, data.Email, data.CellNumber, data.Telephone,
			data.BusinessPhoneNumber, data.Country, data.DistrictName, data.Town, data.Street, data.BoxAddress,
			data.PlotNumber, data.PhysicalAddress, data.ContactType,
		})
	}

	copyCount, err := conn.CopyFrom(
		req.Context(),
		pgx.Identifier{"contractor_contact"},
		[]string{"company_uuid", "first_name", "middle_name", "last_name", "date_of_birth", "email", "cellphone", "telephone",
			"business_number", "country", "district", "town", "street", "box_address", "plot_number",
			"physical_address", "contact_type",
		},
		pgx.CopyFromRows(copyData),
	)

	if err != nil {
		log.Println(err.Error())
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return bunrouter.JSON(w, bunrouter.H{
			"message": "failed to add affiliates",
			"success": false,
		})
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	return bunrouter.JSON(w, bunrouter.H{
		"message":      "successfully added affiliates",
		"number_added": copyCount,
	})
}
