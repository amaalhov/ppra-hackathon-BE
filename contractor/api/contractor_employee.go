package api

import (
	"contractor-services/model"
	"encoding/json"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/uptrace/bunrouter"
)

const insertEmployeeStats = `
	INSERT INTO contractor_employee_stats (company_uuid, total_number_of_citizens, total_number_of_non_citizens, total_employees)
	VALUES ($1, $2, $3, $4)
	RETURNING id;
`

func (c *ContractorStore) AddContractorEmployees(w http.ResponseWriter, req bunrouter.Request) error {
	var reqBody model.ContractorEmployeeReq
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
		req.Context(), insertEmployeeStats, id, reqBody.TotalNumberOfBotswanaCitizens, reqBody.TotalNumberOfNonBotswanaCitizens, reqBody.TotalEmployees)

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

	for _, data := range reqBody.Employees {
		copyData = append(copyData, []interface{}{data.Firstname, data.Middlename, data.Lastname, data.DateOfBirth, data.Gender, statsId})
	}

	copyCount, err := conn.CopyFrom(
		req.Context(),
		pgx.Identifier{"contractor_employee"},
		[]string{"first_name", "middle_name", "last_name", "date_of_birth", "gender", "employee_stats_id"},
		pgx.CopyFromRows(copyData),
	)

	if err != nil {
		log.Println(err.Error())
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return bunrouter.JSON(w, bunrouter.H{
			"message": "failed to add employee data",
			"success": false,
		})
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	return bunrouter.JSON(w, bunrouter.H{
		"message":      "successfully added employee data",
		"number_added": copyCount,
	})

}
