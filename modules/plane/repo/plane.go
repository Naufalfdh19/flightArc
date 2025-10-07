package repo

import (
	"context"
	"database/sql"
	"flight/modules/plane/entity"
	"flight/pkg/apperror"
	"flight/pkg/constant"
)

type PlaneRepo interface {
	AddPlane(ctx context.Context, plane entity.Plane) error
	IsPlaneExistsByRegistrationCode(ctx context.Context, code string) bool 
}

type PlaneRepoImpl struct {
	db *sql.DB
}

func NewPlaneRepo(db *sql.DB) PlaneRepoImpl {
	return PlaneRepoImpl{
		db: db,
	}
}

func (r PlaneRepoImpl) AddPlane(ctx context.Context, plane entity.Plane) error {
	query := `INSERT INTO users (name, seats, capacity, registration_code, status, airline_id)
				VALUES ($1, $2, $3, $4, $5, $6)`

				
	_, err := r.db.Exec(query, 
		plane.Name, plane.Seats, plane.Capacity, plane.RegistrationCode, plane.Status, plane.AirlineId)

	if err != nil {
		return apperror.NewErrInternalServerError(constant.SERVER, apperror.ErrInternalServerError, err)
	}

	return nil
}

func (r PlaneRepoImpl) IsPlaneExistsByRegistrationCode(ctx context.Context, code string) bool {
	var exists bool
	query := `SELECT EXISTS(
		SELECT 1 
		FROM planes 
		WHERE registration_code = $1 AND deleted_at IS NULL)`
	_ = r.db.QueryRow(query, code).Scan(&exists)
	return exists
}

