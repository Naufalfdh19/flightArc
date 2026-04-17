package pagination

type Pagination struct {
	Page int
	TotalPage    int
	TotalElement int
	Data         any
}

type PaginationDto struct {
	Page int `json:"page"`
	TotalPage int `json:"total_page"`
	TotalElement int `json:"total_element"`
	Data         any `json:"data"`
}

type Converter struct {}

func (c Converter) ToDto(pagination Pagination) PaginationDto {
	return PaginationDto{
		Page: pagination.Page,
		TotalPage: pagination.TotalPage,
		TotalElement: pagination.TotalElement,
	}
}

