package database

import (
	"fmt"
	"go.uber.org/dig"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"systemlink/configs"
	"time"
)

type GatewayParams struct {
	dig.In

	Config *configs.AppConfig
}

func New(params GatewayParams) (*gorm.DB, error) {
	//dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
	//	params.Config.DBHost,
	//	params.Config.DBUser,
	//	params.Config.DBPassword,
	//	params.Config.DBName,
	//	params.Config.DBPort,
	//)

	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		params.Config.DBUser,
		params.Config.DBPassword,
		params.Config.DBHost,
		params.Config.DBPort,
		params.Config.DBName,
	)

	_logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		Logger: _logger,
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}
