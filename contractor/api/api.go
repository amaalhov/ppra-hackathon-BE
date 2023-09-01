package api

import (
	"net/http"
	"time"

	"github.com/uptrace/bunrouter"
)

type ContractorDetailsReq struct {
	NameOfCompany        string    `json:"name_of_company"`
	BusinessType         string    `json:"business_type"`
	IsLocalContractor    bool      `json:"is_local_contractor"`
	CipaUin              string    `json:"cipa_uin"`
	NationalIdNumber     string    `json:"omang"`
	IsRegisteredWithCIPA bool      `json:"is_registered_with_cipa"`
	RegistrationDate     time.Time `json:"registration_date"`
	OwnershipType        string    `json:"ownership_type"`
}

func ContractorDetails(w http.ResponseWriter, req http.Request) error {
	return bunrouter.JSON(w, bunrouter.H{
		"message": "contractor service running",
	})
}
