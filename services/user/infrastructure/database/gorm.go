package database

import (
	"fmt"
	"go.uber.org/dig"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
	"user/configs"
)

type GatewayParams struct {
	dig.In

	Config *configs.AppConfig
	//Tracer *opentracing.Tracer `optional:"true"`
}

func New(params GatewayParams) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
		params.Config.DBHost,
		params.Config.DBUser,
		params.Config.DBPassword,
		params.Config.DBName,
		params.Config.DBPort,
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

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: _logger,
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}
