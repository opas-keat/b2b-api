package repo

import (
	"b2b/constant"
	"context"
	"gorm.io/gorm"
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
