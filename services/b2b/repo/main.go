package repo

import (
	"context"
	"gorm.io/gorm"
)

type Gorm struct {
	db *gorm.DB
}

func NewGorm(db *gorm.DB) Gorm {
	return Gorm{db: db}
}

func (g *Gorm) GetDB(ctx context.Context) *gorm.DB {
	return g.db.WithContext(ctx)
}

func (g *Gorm) WithTxFromCtx(ctx context.Context) *gorm.DB {
	if tx := GetTxFromContext(ctx); tx != nil {
		return tx.WithContext(ctx)
	}
	return g.GetDB(ctx)
}
