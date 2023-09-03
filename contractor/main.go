package main

import (
	"context"
	"contractor-services/api"
	"log"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/uptrace/bunrouter"
)

func main() {
	dbUrl := "postgresql://admin:psltest@postgresdb:5432/postgres?sslmode=disable"
	pool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer pool.Close()

	m, err := migrate.New("file://migrations", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Force(1); err != nil && err.Error() != "no change" {
		log.Fatal(err)
	}

	contractorStore := api.NewContractorStore(pool)

	router := bunrouter.New()
	router.POST("/api/contractor/onboard-from-cipa", func(w http.ResponseWriter, req bunrouter.Request) error {
		log.Println("onboarding company from cipa")
		return nil
	})

	router.POST("/api/contractor/company-details", contractorStore.AddContractorDetails)
	router.POST("/api/contractor/company-address", contractorStore.AddContractorAddressDetails)
	router.POST("/api/contractor/contact", contractorStore.AddContractorContacts)
	router.POST("/api/contractor/affiliates", contractorStore.AddContractorAffiliates)
	router.POST("/api/contractor/projects", contractorStore.AddContractorProjects)
	router.POST("/api/contractor/equipment", contractorStore.AddContractorEquipment)
	router.POST("/api/contractor/bank-details", contractorStore.AddContractorBankDetails)
	router.POST("/api/contractor/company-directors", contractorStore.AddContractorDirectors)
	router.POST("/api/contractor/company-shareholders", contractorStore.AddContractorShareHolder)
	router.POST("/api/contractor/company-employees", contractorStore.AddContractorEmployees)
	router.POST("/api/contractor/company-secretary", contractorStore.AddContractorSecretary)
	log.Println("contractor service listening on http://localhost:3002")
	log.Println(http.ListenAndServe(":3002", router))
}
