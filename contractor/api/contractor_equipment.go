package api

import (
	"contractor-services/model"
	"encoding/json"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/uptrace/bunrouter"
)

const insertEquipmentStats = `
	INSERT INTO contractor_equipment (value_of_assets, value_of_equipment, paid_up_capital)
	VALUES ($1, $2, $3)
	RETURNING id;
`

func (c *ContractorStore) AddContractorEquipment(w http.ResponseWriter, req bunrouter.Request) error {

	var reqBody model.ContractorVehicleReq

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
		req.Context(), insertEquipmentStats, reqBody.ValueOfAssets, reqBody.ValueOfEquipment, reqBody.PaidUpCapital)

	var statsId int64

	err = row.Scan(&statsId)

	if err != nil {
		log.Println(err.Error())
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return bunrouter.JSON(w, bunrouter.H{
			"message": "failed to add address",
			"success": false,
		})
	}

	var copyData [][]interface{}

	for _, data := range reqBody.Vehicles {
		copyData = append(copyData, []interface{}{data.RegisteredOwner, data.Ownership, data.RegistrationNumber, data.DateOfRegistration, data.VehicleModel, statsId})
	}

	copyCount, err := conn.CopyFrom(
		req.Context(),
		pgx.Identifier{"contractor_vehicle"},
		[]string{"registered_owner", "ownership", "registration_number", "date_of_registration", "vehicle_model", "equipment_id"},
		pgx.CopyFromRows(copyData),
	)

	if err != nil {
		log.Println(err.Error())
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return bunrouter.JSON(w, bunrouter.H{
			"message": "failed to add vehicles",
			"success": false,
		})
	}

	copyData = nil
	for _, data := range reqBody.Plants {
		copyData = append(copyData, []interface{}{data.RegisteredOwner, data.Ownership, data.Description, data.RegistrationNumber, data.DateOfPurchase, statsId})
	}

	copyCount, err = conn.CopyFrom(
		req.Context(),
		pgx.Identifier{"contractor_plant"},
		[]string{"registered_owner", "ownership", "description", "registration_number", "date_of_purchase", "equipment_id"},
		pgx.CopyFromRows(copyData),
	)

	if err != nil {
		log.Println(err.Error())
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return bunrouter.JSON(w, bunrouter.H{
			"message": "failed to add plants and equipments",
			"success": false,
		})
	}

	copyData = nil
	for _, data := range reqBody.Properties {
		copyData = append(copyData, []interface{}{data.Ownership, data.PresentValue, data.AttachmentUrl, data.Locality, statsId})
	}

	copyCount, err = conn.CopyFrom(
		req.Context(),
		pgx.Identifier{"contractor_property"},
		[]string{"ownership", "present_value", "attachment_url", "locality", "equipment_id"},
		pgx.CopyFromRows(copyData),
	)

	if err != nil {
		log.Println(err.Error())
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return bunrouter.JSON(w, bunrouter.H{
			"message": "failed to add properties",
			"success": false,
		})
	}

	copyData = nil
	for _, data := range reqBody.OfficeEquipments {
		copyData = append(copyData, []interface{}{data.OfficeEquipment, data.PresentValue, data.AttachmentUrl, statsId})
	}

	copyCount, err = conn.CopyFrom(
		req.Context(),
		pgx.Identifier{"contractor_office_equipment"},
		[]string{"office_equipment", "present_value", "attachment_url", "equipment_id"},
		pgx.CopyFromRows(copyData),
	)

	if err != nil {
		log.Println(err.Error())
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return bunrouter.JSON(w, bunrouter.H{
			"message": "failed to add office equipment",
			"success": false,
		})
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	return bunrouter.JSON(w, bunrouter.H{
		"message":      "successfully added company assets",
		"number_added": copyCount,
	})

}
