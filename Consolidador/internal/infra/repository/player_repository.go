package repository

import (
	"database/sql"
	"github.com/rutanioka/Projeto-Cartola-2022/Consolidador/internal/infra/db"
)
type PlayerRepository struct {
	Repository
}

func NewPlayerRepository (dbConn *sql.DB) *PlayerRepository {
	return &PlayerRepository{
		Repository: Repository{
			dbConn: dbConn,
			Queries: db.New(dbConn),
		},
	}
}