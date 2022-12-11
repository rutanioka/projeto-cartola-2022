package repository

import (
	"database/sql"
	"github.com/rutanioka/Projeto-Cartola-2022/Consolidador/internal/infra/db"
)


type Repository struct{
	dbConn *sql.DB
	*db.Queries
}