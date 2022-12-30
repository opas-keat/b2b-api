package repo

import (
	"context"
	"gorm.io/gorm"
	"member/constant"
)

func GetTxFromContext(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(constant.TxContextKey).(*gorm.DB); ok {
		return tx
	}
	return nil
}
