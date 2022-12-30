package user

import (
	"gorm.io/gorm"
	"member/repo"
)

type Gorm struct {
	repo.Gorm
}

func NewGormRepo(db *gorm.DB) Repo {
	return &Gorm{repo.NewGorm(db)}
}
