package common

import (
	"gorm.io/gorm"
)

func Paginate(page, limit, totalpage int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page = checkPage(page, totalpage)
		offset := (page - 1) * limit
		return db.Offset(offset).Limit(limit)
	}
}

func checkPage(page, totalpage int) int {
	if page > totalpage {
		return totalpage
	} else if page < 0 {
		return 1
	}
	return page
}
