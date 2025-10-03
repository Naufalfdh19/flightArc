package repo

import (
	"github.com/jackc/pgx/v5"
)

type AdminRepo interface {
}

type AdminRepoImpl struct {
	db *pgx.Conn
}

func NewAdminRepo(db *pgx.Conn) AdminRepoImpl {
	return AdminRepoImpl{
		db: db,
	}
}
