package repo

import (
	"context"
	"logs/constant"

	"gorm.io/gorm"
)

func GetTxFromContext(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(constant.TxContextKey).(*gorm.DB); ok {
		return tx
	}
	return nil
}
