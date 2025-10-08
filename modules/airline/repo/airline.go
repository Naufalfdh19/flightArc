package repo

import (
	"context"
	"database/sql"
)

type AirlineRepo interface {
	IsAirlineExistsById(ctx context.Context, id int) bool 
}

type AirlineRepoImpl struct {
	db *sql.DB
}

func NewAirlineRepo(db *sql.DB) AirlineRepoImpl {
	return AirlineRepoImpl{
		db: db,
	}
}

func (r AirlineRepoImpl) IsAirlineExistsById(ctx context.Context, id int) bool {
	var exists bool
	query := `SELECT EXISTS(
		SELECT 1 
		FROM airlines 
		WHERE id = $1 AND deleted_at IS NULL)`
	_ = r.db.QueryRow(query, id).Scan(&exists)
	return exists
}
