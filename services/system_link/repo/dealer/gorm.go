package dealer

import (
	"gorm.io/gorm"
	"systemlink/repo"
)

type Gorm struct {
	repo.Gorm
}

func NewGormRepo(db *gorm.DB) Repo {
	return &Gorm{repo.NewGorm(db)}
}
