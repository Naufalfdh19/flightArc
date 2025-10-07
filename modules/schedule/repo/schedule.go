package repo

import (
	"context"
	"database/sql"
	"flight/modules/schedule/entity"
	"flight/modules/schedule/queryparams"
	"flight/pkg/apperror"
	"flight/pkg/constant"
)

type ScheduleRepo interface {
	GetSchedules(ctx context.Context, queryParams queryparams.QueryParams) ([]entity.Schedule, error)
	GetTotalSchedule(ctx context.Context) (int, error)
}

type ScheduleRepoImpl struct {
	db *sql.DB
}

func NewScheduleRepo(db *sql.DB) ScheduleRepoImpl {
	return ScheduleRepoImpl{
		db: db,
	}
}

func (r ScheduleRepoImpl) GetSchedules(ctx context.Context, queryParams queryparams.QueryParams) ([]entity.Schedule, error) {
	var schedules []entity.Schedule

	query := `SELECT id, origin, destination, status, departure_date
				FROM flight.flight_schedules
				WHERE deleted_at IS NULL`
	query += queryparams.AddPagination(queryParams)

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, apperror.NewErrInternalServerError(constant.SERVER, apperror.ErrInternalServerError, err)
	}
	defer rows.Close()

	for rows.Next() {
		var Schedule entity.Schedule

		err := rows.Scan(
			&Schedule.Id,
			&Schedule.Origin,
			&Schedule.Destination,
			&Schedule.Status,
			&Schedule.DepartureDate)
		if err != nil {
			return nil, apperror.NewErrInternalServerError(constant.SERVER, apperror.ErrInternalServerError, err)
		}
		schedules = append(schedules, Schedule)
	}

	return schedules, nil
}

func (r ScheduleRepoImpl) GetTotalSchedule(ctx context.Context) (int, error) {
	var totalSchedule int
	query := `SELECT COUNT(*) 
				FROM flight.flight_schedules
				WHERE deleted_at IS NULL`

	err := r.db.QueryRow(query).Scan(&totalSchedule)
	if err != nil {
		return 0, apperror.NewErrInternalServerError(constant.SERVER, apperror.ErrInternalServerError, err)
	}

	return totalSchedule, nil
}

// {
// 	{
// 		seat_number: 1,
// 		class: business,
// 		price: 100000,
// 		isEmergency:
// 		isAvailable:
// 	},
// }


