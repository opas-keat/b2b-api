package repo

import (
	"context"
	"gorm.io/gorm"
	"users/constant"
)

func GetTxFromContext(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(constant.TxContextKey).(*gorm.DB); ok {
		return tx
	}
	return nil
}

func WherePointer[T any](format string, value *T) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if value != nil {
			return db.Where(format, *value)
		}
		return db
	}
}
