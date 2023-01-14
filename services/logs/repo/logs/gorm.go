package logs

import (
	"logs/repo"

	"gorm.io/gorm"
)

type Gorm struct {
	repo.Gorm
}

func NewGormRepo(db *gorm.DB) Repo {
	return &Gorm{repo.NewGorm(db)}
}
