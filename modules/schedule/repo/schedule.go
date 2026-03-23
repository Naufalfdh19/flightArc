package repo

import (
	"context"
	"database/sql"
	"flight/modules/schedule/entity"
	"flight/modules/schedule/queryparams"
	"flight/pkg/apperror"
	"flight/pkg/constant"
	"log"
)

type ScheduleRepo interface {
	GetFlights(ctx context.Context, queryParams queryparams.QueryParams) ([]entity.Flight, error)
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

func (r ScheduleRepoImpl) GetFlights(ctx context.Context, queryParams queryparams.QueryParams) ([]entity.Flight, error) {
	var flights []entity.Flight

	query := `SELECT 
					s.id,
					s.origin_code, 
					org.name AS origin_airport_name,
					org.city AS origin_city,
					org.country AS origin_country,
					s.destination_code, 
					dst.name AS destination_airport_name,
					dst.city AS destination_city,
					dst.country AS destination_country,
					s.status, 
					s.departure_time, 
					s.arrival_time
				FROM schedules s
				LEFT JOIN airports org ON s.origin_code = org.code
				LEFT JOIN airports dst ON s.destination_code = dst.code
				WHERE s.deleted_at IS NULL`
	
	log.Print(queryParams.Page, queryParams.Limit)
	query += queryparams.AddPagination(queryParams)

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, apperror.NewErrInternalServerError(constant.SERVER, apperror.ErrInternalServerError, err)
	}
	defer rows.Close()

	for rows.Next() {
		var flight entity.Flight

		err := rows.Scan(
			&flight.Id,
			&flight.Origin.Code,
			&flight.Origin.Name,
			&flight.Origin.City,
			&flight.Origin.Country,
			&flight.Destination.Code,
			&flight.Destination.Name,
			&flight.Destination.City,
			&flight.Destination.Country,
			&flight.Status,
			&flight.DepartureTime,
			&flight.ArrivalTime,
		)
		if err != nil {
			return nil, apperror.NewErrInternalServerError(constant.SERVER, apperror.ErrInternalServerError, err)
		}
		flights = append(flights, flight)
	}

	return flights, nil
}

func (r ScheduleRepoImpl) GetTotalSchedule(ctx context.Context) (int, error) {
	var totalSchedule int
	query := `SELECT COUNT(*) 
				FROM schedules
				WHERE deleted_at IS NULL`

	err := r.db.QueryRow(query).Scan(&totalSchedule)
	if err != nil {
		return 0, apperror.NewErrInternalServerError(constant.SERVER, apperror.ErrInternalServerError, err)
	}

	return totalSchedule, nil
}
