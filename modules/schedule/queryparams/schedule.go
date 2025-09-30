package queryparams

import (
	"fmt"
	"log"
)

type QueryParams struct {
	Page  int
	Limit int
}

type QueryParamsDto struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

type QueryParamsConverter struct{}

func (c QueryParamsConverter) ConvertDtoToEntity(queryparamsDto QueryParamsDto) QueryParams {
	return QueryParams{
		Page:  queryparamsDto.Page,
		Limit: queryparamsDto.Limit,
	}
}

func AddPagination(queryParams QueryParams) string {
	log.Println("testboy: ", queryParams.Page, queryParams.Limit)
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