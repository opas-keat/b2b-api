package user

import (
	"gorm.io/gorm"
	"user/repo"
)

type Gorm struct {
	repo.Gorm
}

func NewGormRepo(db *gorm.DB) Repo {
	return &Gorm{repo.NewGorm(db)}
}
