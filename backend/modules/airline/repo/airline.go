package repo

import (
	"context"

	"gorm.io/gorm"
)

type AirlineRepo interface {
	IsAirlineExistsById(ctx context.Context, id int) bool
}

type AirlineRepoImpl struct {
	db *gorm.DB
}

func NewAirlineRepo(db *gorm.DB) AirlineRepoImpl {
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
	_ = r.db.Raw(query, id).Scan(&exists)
	return exists
}
