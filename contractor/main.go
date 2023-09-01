package main

import (
	"contractor-services/api"
	"log"
	"net/http"

	"github.com/uptrace/bunrouter"
)

func main() {
	router := bunrouter.New()
	router.POST("/api/contractor/company-details", api.ContractorDetails)
	log.Println("contractor service listening on http://localhost:3002")
	log.Println(http.ListenAndServe(":3002", router))
}
