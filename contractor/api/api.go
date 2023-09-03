package api

import (
	"contractor-services/model"
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/uptrace/bunrouter"
)

type ContractorStore struct {
	db *pgxpool.Pool
}

func NewContractorStore(db *pgxpool.Pool) *ContractorStore {
	return &ContractorStore{db: db}
}

func (c *ContractorStore) AddContractorEquipment(w http.ResponseWriter, req bunrouter.Request) error {

	var reqBody model.ContractorVehicleReq

	json.NewDecoder(req.Body).Decode(&reqBody)

	return bunrouter.JSON(w, reqBody)
}

func (c *ContractorStore) AddContractorEmployees(w http.ResponseWriter, req bunrouter.Request) error {
	var reqBody model.ContractorEmployeeReq

	json.NewDecoder(req.Body).Decode(&reqBody)

	return bunrouter.JSON(w, reqBody)
}
