package queryparams

import (
	"fmt"
)

type QueryParams struct {
	OriginCode string
	OriginCity string
	DestinationCode string
	DestinationCity string
	Page  int
	Limit int
}

type QueryParamsDto struct {
	OriginCode string `form:"origin_code"`
	OriginCity string `form:"origin_city"`
	DestinationCode string `form:"destination_code"`
	DestinationCity string `form:"destination_city"`
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

type QueryParamsConverter struct{}

func (c QueryParamsConverter) ConvertDtoToEntity(queryparamsDto QueryParamsDto) QueryParams {
	return QueryParams{
		OriginCode: queryparamsDto.OriginCode,
		OriginCity: queryparamsDto.OriginCity,
		DestinationCode: queryparamsDto.DestinationCode,
		DestinationCity: queryparamsDto.DestinationCity,
		Page:  queryparamsDto.Page,
		Limit: queryparamsDto.Limit,
	}
}

func AddFilter(queryParams QueryParams) string {
	filter := ""

	if (queryParams.OriginCity != "") {
		filter += fmt.Sprintf(" AND s.origin_city = '%s'", queryParams.OriginCity)
	}
	if (queryParams.OriginCode != "") {
		filter += fmt.Sprintf(" AND s.origin_code = '%s'", queryParams.OriginCode)
	}
	if (queryParams.DestinationCity != "") {
		filter += fmt.Sprintf(" AND s.destination_city = '%s'", queryParams.DestinationCity)
	}
	if (queryParams.DestinationCode != "") {
		filter += fmt.Sprintf(" AND s.destination_code = '%s'", queryParams.DestinationCode)
	}

	return filter
}

func AddPagination(queryParams QueryParams) string {
	return fmt.Sprintf(` LIMIT %d OFFSET %d`, queryParams.Limit, (queryParams.Page-1)*queryParams.Limit)
}

func CheckPage(queryParams *QueryParams, totalPage int) {
	if queryParams.Page <= 0 {
		queryParams.Page = 1
	} else if queryParams.Page > totalPage {
		queryParams.Page = totalPage
	}
}

func CheckLimit(queryParams *QueryParams) {
	if queryParams.Limit < 1 {
		queryParams.Limit = 1
	}
}
