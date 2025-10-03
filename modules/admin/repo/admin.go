package repo

import (
	"database/sql"
)

type AdminRepo interface {
}

type AdminRepoImpl struct {
	db *sql.DB
}

func NewAdminRepo(db *sql.DB) AdminRepoImpl {
	return AdminRepoImpl{
		db: db,
	}
}
