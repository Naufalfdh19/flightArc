package repo

import (
	"gorm.io/gorm"
)

type AdminRepo interface {
}

type AdminRepoImpl struct {
	db *gorm.DB
}

func NewAdminRepo(db *gorm.DB) AdminRepoImpl {
	return AdminRepoImpl{
		db: db,
	}
}
