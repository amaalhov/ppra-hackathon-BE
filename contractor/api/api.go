package api

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type ContractorStore struct {
	db *pgxpool.Pool
}

func NewContractorStore(db *pgxpool.Pool) *ContractorStore {
	return &ContractorStore{db: db}
}
